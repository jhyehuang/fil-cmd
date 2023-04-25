package intern

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/filecoin-project/lotus/chain/types"
	chaintypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/jhyehuang/fil-cmd/intern/log"
	"github.com/rickiey/loggo"
	"golang.org/x/xerrors"
	"strconv"
)

// Have a type with some exported methods
type SimpleServerHandler struct {
	n int
}

func (h *SimpleServerHandler) GetPower(miner string, heigh string) ([]interface{}, error) {
	ctx := context.Background()
	fullNodeApi, err := NewFullNodeAPIV2("", "")
	if err != nil {
		loggo.Error(err)
		return nil, err
	}
	tStart := build.Clock.Now()
	// MiningBase is the tipset on top of which we plan to construct our next block.
	// Refer to godocs on GetBestMiningCandidate.
	type MiningBase struct {
		TipSet     *types.TipSet
		NullRounds abi.ChainEpoch
	}

	mineraddr, err := address.NewFromString(miner)
	fmt.Printf("mineraddr: %s\n", mineraddr)
	if err != nil {
		loggo.Error(err)
		return nil, err
	}
	epoch, err := strconv.ParseUint(heigh, 10, 64)
	if err != nil {
		loggo.Error(err)
		return nil, err
	}

	currHead, err := fullNodeApi.ChainHead(ctx)
	currEpoch := currHead.Height()

	fmt.Printf("currEpoch: %d\n", currEpoch)

	tipSet, err := fullNodeApi.ChainGetTipSetByHeight(ctx, abi.ChainEpoch(epoch-1), chaintypes.EmptyTSK)
	if err != nil {
		err = xerrors.Errorf("failed to get mining base info: %w", err)
		return nil, err
	}

	base := &MiningBase{TipSet: tipSet}

	var winner *types.ElectionProof
	var mbi *api.MiningBaseInfo
	var rbase types.BeaconEntry

	mbi, err = fullNodeApi.MinerGetBaseInfo(ctx, mineraddr, abi.ChainEpoch(epoch), base.TipSet.Key())
	if err != nil {
		loggo.Errorf("failed to get mining base info: %s", err)
		err = xerrors.Errorf("failed to get mining base info: %w", err)
		return nil, err
	}

	bvals := mbi.BeaconEntries
	rbase = mbi.PrevBeaconEntry
	if len(bvals) > 0 {
		rbase = bvals[len(bvals)-1]
	}
	//
	//winner, err = gen.IsRoundWinner(ctx, base.TipSet, abi.ChainEpoch(epoch), mineraddr, rbase, mbi, fullNodeApiAuth)
	//if err != nil {
	//	err = xerrors.Errorf("failed to check if we win next round: %w", err)
	//	return
	//}

	logStruct := []interface{}{
		"tookMilliseconds", (build.Clock.Now().UnixNano() - tStart.UnixNano()) / 1_000_000,
		"forRound", int64(epoch),
		"baseEpoch", int64(base.TipSet.Height()),
		"baseDeltaSeconds", uint64(tStart.Unix()) - base.TipSet.MinTimestamp(),
		"nullRounds", int64(base.NullRounds),
		"lateStart", false,
		"beaconEpoch", rbase.Round,
		"lookbackEpochs", int64(policy.ChainFinality), // hardcoded as it is unlikely to change again: https://github.com/filecoin-project/lotus/blob/v1.8.0/chain/actors/policy/policy.go#L180-L186
		"networkPowerAtLookback", mbi.NetworkPower.String(),
		"minerPowerAtLookback", mbi.MinerPower.String(),
		"isEligible", mbi.EligibleForMining,
		"isWinner", (winner != nil),
		"error", err,
	}

	log.Warnf("completed mineOne", logStruct)

	return logStruct, nil
}
