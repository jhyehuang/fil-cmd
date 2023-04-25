package cmd

import (
	"context"
	"github.com/filecoin-project/lotus/api"
	"github.com/jhyehuang/fil-cmd/config"
	"github.com/jhyehuang/fil-cmd/intern"
	"github.com/jhyehuang/fil-cmd/intern/log"
	"github.com/rickiey/loggo"
	"time"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/actors"

	miner3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/types"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	"golang.org/x/xerrors"

	logging "github.com/ipfs/go-log/v2"

	"github.com/urfave/cli/v2"
)

var disputeLog = logging.Logger("disputer")

const Confidence = 9

type minerDeadline struct {
	miner address.Address
	index uint64
}

var ChainDisputeSetCmd = cli.Command{
	Name:  "disputer",
	Usage: "interact with the window post disputer",
	Flags: []cli.Flag{
		//&cli.StringFlag{
		//	Name:  "max-fee",
		//	Usage: "Spend up to X FIL per DisputeWindowedPoSt message",
		//},
		&cli.StringFlag{
			Name:  "from",
			Usage: "optionally specify the account to send messages from",
		},
	},
	Subcommands: []*cli.Command{
		disputerStartCmd,
	},
}

var disputerStartCmd = &cli.Command{
	Name:      "start",
	Usage:     "Start the window post disputer",
	ArgsUsage: "[minerAddress]",
	//Flags: []cli.Flag{
	//	&cli.Uint64Flag{
	//		Name:  "start-epoch",
	//		Usage: "only start disputing PoSts after this epoch ",
	//	},
	//},
	Action: func(cctx *cli.Context) error {

		url := config.LotusNodeAddr
		token := config.LotusToken
		Lv1api, err := intern.NewFullNodeAPIV2(url, token)
		if err != nil {
			loggo.Panicf("connecting with lotus failed: %s", err)
			return err
		}

		ctx := cctx.Context

		fromAddr, err := address.NewFromString(cctx.String("from"))
		if err != nil {
			loggo.Error(err)
			return err
		}

		mss, err := getMaxFee("")
		if err != nil {
			loggo.Error(err)
			return err
		}

		startEpoch := abi.ChainEpoch(0)

		disputeLog.Info("checking sync status")

		head, err := Lv1api.ChainHead(ctx)
		if err != nil {
			loggo.Error(err)
			return err
		}

		lastEpoch := head.Height()
		lastStatusCheckEpoch := lastEpoch
		log.Errorf("last epoch: %d", lastEpoch)

		// build initial deadlineMap

		minerList, err := Lv1api.StateListMiners(ctx, types.EmptyTSK)
		if err != nil {
			loggo.Error(err)
			return err
		}
		log.Infof("miner list: %d", len(minerList))

		//minerList = minerList[:1000]

		knownMiners := make(map[address.Address]struct{})
		deadlineMap := make(map[abi.ChainEpoch][]minerDeadline)
		for i, miner := range minerList {
			dClose, dl, err := makeMinerDeadline(ctx, Lv1api, miner)
			if err != nil {
				loggo.Error(err)
				return err
			}

			deadlineMap[dClose+Confidence] = append(deadlineMap[dClose+Confidence], *dl)

			knownMiners[miner] = struct{}{}
			log.Infof("miner: %d, %s", i, miner.String())
		}

		// when this fires, check for newly created miners, and purge any "missed" epochs from deadlineMap
		statusCheckTicker := time.NewTicker(time.Hour)
		defer statusCheckTicker.Stop()

		disputeLog.Info("starting up window post disputer")

		applyTsk := func(tsk types.TipSetKey) error {
			disputeLog.Infow("last checked epoch", "epoch", lastEpoch)
			dls, ok := deadlineMap[lastEpoch]
			log.Infof("deadlineMap: %d", len(deadlineMap))
			log.Infof("dls: %d", len(dls))

			delete(deadlineMap, lastEpoch)
			if !ok || startEpoch >= lastEpoch {

				// no deadlines closed at this epoch - Confidence, or we haven't reached the start cutoff yet
				log.Errorf("no deadlines closed at this epoch - Confidence, or we haven't reached the start cutoff yet")
				return nil
			}

			dpmsgs := make([]*types.Message, 0)

			startTime := time.Now()
			proofsChecked := uint64(0)

			// TODO: Parallelizeable
			for _, dl := range dls {
				fullDeadlines, err := Lv1api.StateMinerDeadlines(ctx, dl.miner, tsk)
				if err != nil {
					return xerrors.Errorf("failed to load deadlines: %w", err)
				}

				if int(dl.index) >= len(fullDeadlines) {
					return xerrors.Errorf("deadline index %d not found in deadlines", dl.index)
				}

				disputableProofs := fullDeadlines[dl.index].DisputableProofCount
				proofsChecked += disputableProofs

				ms, err := makeDisputeWindowedPosts(ctx, Lv1api, dl, disputableProofs, fromAddr)
				if err != nil {
					return xerrors.Errorf("failed to check for disputes: %w", err)
				}

				dpmsgs = append(dpmsgs, ms...)

				dClose, dl, err := makeMinerDeadline(ctx, Lv1api, dl.miner)
				if err != nil {
					return xerrors.Errorf("making deadline: %w", err)
				}

				deadlineMap[dClose+Confidence] = append(deadlineMap[dClose+Confidence], *dl)
			}

			disputeLog.Infow("checked proofs", "count", proofsChecked, "duration", time.Since(startTime))

			// TODO: Parallelizeable / can be integrated into the previous deadline-iterating for loop
			for _, dpmsg := range dpmsgs {
				disputeLog.Infow("disputing a PoSt", "miner", dpmsg.To)
				m, err := Lv1api.MpoolPushMessage(ctx, dpmsg, mss)
				if err != nil {
					disputeLog.Errorw("failed to dispute post message", "err", err.Error(), "miner", dpmsg.To)
				} else {
					disputeLog.Infow("submited dispute", "mcid", m.Cid(), "miner", dpmsg.To)
				}
			}

			return nil
		}

		disputeLoop := func() error {
			head, err := Lv1api.ChainHead(ctx)
			if err != nil {
				loggo.Error(err)
				return err
			}

			lastEpoch := head.Height()
			log.Infof("last epoch: %d", lastEpoch)
			if lastEpoch <= lastStatusCheckEpoch {
				return nil
			} else {
				for m, _ := range knownMiners {
					dClose, dl, err := makeMinerDeadline(ctx, Lv1api, m)
					if err != nil {
						continue
					}
					deadlineMap[dClose+Confidence] = append(deadlineMap[dClose+Confidence], *dl)

				}

				err = applyTsk(head.Key())
				if err != nil {
					log.Error(err)
					return err
				}
			}

			return nil
		}
		timer := time.NewTimer(time.Second * 30)

		for {
			select {
			case <-timer.C:

				err := disputeLoop()
				if err == context.Canceled {
					disputeLog.Info("disputer shutting down")
					break
				}
				if err != nil {
					disputeLog.Errorw("disputer shutting down", "err", err)
					return err
				}
				timer.Reset(time.Second * 30)
			case <-statusCheckTicker.C:
				disputeLog.Infof("running status check")

				minerList, err = Lv1api.StateListMiners(ctx, types.EmptyTSK)
				if err != nil {
					return xerrors.Errorf("failed to list miners: %w", err)
				}
				//minerList = minerList[:1000]

				for _, m := range minerList {
					_, ok := knownMiners[m]
					if !ok {
						dClose, dl, err := makeMinerDeadline(ctx, Lv1api, m)
						if err != nil {
							return xerrors.Errorf("making deadline: %w", err)
						}

						deadlineMap[dClose+Confidence] = append(deadlineMap[dClose+Confidence], *dl)

						knownMiners[m] = struct{}{}
					}
				}

				for ; lastStatusCheckEpoch < lastEpoch; lastStatusCheckEpoch++ {
					// if an epoch got "skipped" from the deadlineMap somehow, just fry it now instead of letting it sit around forever
					_, ok := deadlineMap[lastStatusCheckEpoch]
					if ok {
						disputeLog.Infow("epoch skipped during execution, deleting it from deadlineMap", "epoch", lastStatusCheckEpoch)
						delete(deadlineMap, lastStatusCheckEpoch)
					}
				}

				log.Infof("status check complete")
			case <-ctx.Done():
				return nil
			}
		}

		return nil
	},
}

// for a given miner, index, and maxPostIndex, tries to dispute posts from 0...postsSnapshotted-1
// returns a list of DisputeWindowedPoSt msgs that are expected to succeed if sent
func makeDisputeWindowedPosts(ctx context.Context, api api.FullNode, dl minerDeadline, postsSnapshotted uint64, sender address.Address) ([]*types.Message, error) {
	disputes := make([]*types.Message, 0)

	for i := uint64(0); i < postsSnapshotted; i++ {

		dpp, aerr := actors.SerializeParams(&miner3.DisputeWindowedPoStParams{
			Deadline:  dl.index,
			PoStIndex: i,
		})

		if aerr != nil {
			return nil, xerrors.Errorf("failed to serailize params: %w", aerr)
		}

		dispute := &types.Message{
			To:     dl.miner,
			From:   sender,
			Value:  big.Zero(),
			Method: builtin3.MethodsMiner.DisputeWindowedPoSt,
			Params: dpp,
		}

		rslt, err := api.StateCall(ctx, dispute, types.EmptyTSK)
		if err == nil && rslt.MsgRct.ExitCode == 0 {
			disputes = append(disputes, dispute)
		}

	}

	return disputes, nil
}

func makeMinerDeadline(ctx context.Context, api api.FullNode, mAddr address.Address) (abi.ChainEpoch, *minerDeadline, error) {
	dl, err := api.StateMinerProvingDeadline(ctx, mAddr, types.EmptyTSK)
	if err != nil {
		return -1, nil, xerrors.Errorf("getting proving index list: %w", err)
	}

	return dl.Close, &minerDeadline{
		miner: mAddr,
		index: dl.Index,
	}, nil
}

func getSender(ctx context.Context, api api.FullNode, fromStr string) (address.Address, error) {
	if fromStr == "" {
		return api.WalletDefaultAddress(ctx)
	}

	addr, err := address.NewFromString(fromStr)
	if err != nil {
		return address.Undef, err
	}

	has, err := api.WalletHas(ctx, addr)
	if err != nil {
		return address.Undef, err
	}

	if !has {
		return address.Undef, xerrors.Errorf("wallet doesn't contain: %s ", addr)
	}

	return addr, nil
}

func getMaxFee(maxStr string) (*api.MessageSendSpec, error) {
	if maxStr != "" {
		maxFee, err := types.ParseFIL(maxStr)
		if err != nil {
			return nil, xerrors.Errorf("parsing max-fee: %w", err)
		}
		return &api.MessageSendSpec{
			MaxFee: types.BigInt(maxFee),
		}, nil
	}

	return nil, nil
}
