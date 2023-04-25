package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/lotus/api"
	chaintypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/storage/sealer/storiface"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/jhyehuang/fil-cmd/intern/log"
	"net/http"
	"strings"
)

type FilRPC struct {
	httprpc *HTTPRpc
}

func New(url string, options ...func(rpc *FilRPC)) *FilRPC {
	token := ""
	minerAddr := strings.Split(url, ":")
	if len(minerAddr) > 1 {
		token = minerAddr[0]
		urlIp := strings.Split(minerAddr[1], "/")
		//fmt.Println("urlIp", urlIp)
		url = fmt.Sprintf("http://%s:%s/rpc/v0", urlIp[2], urlIp[4])
	}
	rpc := &FilRPC{
		httprpc: &HTTPRpc{
			url,
			token,
			http.DefaultClient,
			log.Logger,
		},
	}

	for _, option := range options {
		option(rpc)
	}

	return rpc
}

func (e *FilRPC) getTipset(method string, params ...interface{}) (*chaintypes.ExpTipSet, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset chaintypes.ExpTipSet
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) getOriTipset(method string, params ...interface{}) (*chaintypes.TipSet, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset chaintypes.TipSet
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) FilGetOriTipsetByHeight(number int) (*chaintypes.TipSet, error) {
	return e.getOriTipset("Filecoin.ChainGetTipSetByHeight", number, nil)
}

func (e *FilRPC) getMessages(method string, params ...interface{}) (Messages, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return Messages{}, err
	}
	if bytes.Equal(result, []byte("null")) {
		return Messages{}, nil
	}

	var msgs Messages
	err = json.Unmarshal(result, &msgs)
	if err != nil {
		return Messages{}, err
	}

	return msgs, nil
}

func (e *FilRPC) FilGetTipsetByHeight(number int) (*chaintypes.ExpTipSet, error) {
	return e.getTipset("Filecoin.ChainGetTipSetByHeight", number, nil)
}

func (e *FilRPC) FilGetMessagesByCID(blockcid string) (Messages, error) {

	cid := make(map[string]string)
	cid["/"] = blockcid

	msgs, err := e.getMessages("Filecoin.ChainGetBlockMessages", cid)
	msgs.Blockcid = blockcid
	return msgs, err
}

func (e *FilRPC) getActor(method string, params ...interface{}) (*chaintypes.Actor, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset chaintypes.Actor
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) FilGetActor(to address.Address, ts chaintypes.TipSetKey) (*chaintypes.Actor, error) {

	msgs, err := e.getActor("Filecoin.StateGetActor", to, chaintypes.EmptyTSK)

	return msgs, err
}

func (e *FilRPC) getStateSearchMsg(method string, params ...interface{}) (*api.MsgLookup, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset api.MsgLookup
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) FilGetStateSearchMsg(number cid.Cid, ts chaintypes.TipSetKey) (*api.MsgLookup, error) {
	//return e.getStateSearchMsg("Filecoin.StateSearchMsg", ts, number, 1000, false)
	return e.getStateSearchMsg("Filecoin.StateSearchMsg", number)
}

func (e *FilRPC) getOriMessages(method string, params ...interface{}) (*chaintypes.Message, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return &chaintypes.Message{}, err
	}
	if bytes.Equal(result, []byte("null")) {
		return &chaintypes.Message{}, nil
	}

	var msgs chaintypes.Message
	err = json.Unmarshal(result, &msgs)
	if err != nil {
		return &chaintypes.Message{}, err
	}

	return &msgs, nil
}

func (e *FilRPC) getNetworkVersion(method string, params ...interface{}) (uint, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return 0, err
	}
	if bytes.Equal(result, []byte("null")) {
		return 0, nil
	}

	var msgs network.Version
	err = json.Unmarshal(result, &msgs)
	if err != nil {
		return 0, err
	}

	return uint(msgs), nil
}

func (e *FilRPC) FilGetMessages(blockcid cid.Cid) (*chaintypes.Message, error) {

	msgs, err := e.getOriMessages("Filecoin.ChainGetMessage", blockcid)
	return msgs, err
}

func (e *FilRPC) FilGetStateNetworkVersion() (uint, error) {

	msgs, err := e.getNetworkVersion("Filecoin.StateNetworkVersion", chaintypes.EmptyTSK)
	return msgs, err
}

func (e *FilRPC) FilGetMinerVersion(m address.Address) (string, error) {
	return e.getMinerVersion("Filecoin.GetMinerVersion", m, nil)
}

func (e *FilRPC) getMinerVersion(method string, params ...interface{}) (string, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return "", err
	}
	if bytes.Equal(result, []byte("null")) {
		return "", nil
	}

	var tipset string
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return "", err
	}

	return tipset, nil
}

func (e *FilRPC) FilGetDealStatus(miner address.Address, propCid cid.Cid, dealUUID *uuid.UUID) (*storagemarket.ProviderDealState, error) {
	return e.getDealStatus("Filecoin.ClientGetDealStatus", miner, propCid, dealUUID, nil)
}

func (e *FilRPC) getDealStatus(method string, params ...interface{}) (*storagemarket.ProviderDealState, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var tipset storagemarket.ProviderDealState
	err = json.Unmarshal(result, &tipset)
	if err != nil {
		return nil, err
	}

	return &tipset, nil
}

func (e *FilRPC) getMarketListIncompleteDeals(method string, params ...interface{}) (*[]storagemarket.MinerDeal, error) {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}
	var minerDeal []storagemarket.MinerDeal
	err = json.Unmarshal(result, &minerDeal)
	if err != nil {
		return nil, err
	}

	return &minerDeal, nil
}

func (e *FilRPC) getWorkerStats(method string, params ...interface{}) (*map[uuid.UUID]storiface.WorkerStats, error) {
	result, err := e.httprpc.CallDo(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}
	var workerStats map[uuid.UUID]storiface.WorkerStats
	err = json.Unmarshal(result, &workerStats)
	if err != nil {
		return nil, err
	}

	return &workerStats, nil
}

func (e *FilRPC) FilGetWorkerStats() (*map[uuid.UUID]storiface.WorkerStats, error) {
	return e.getWorkerStats("Filecoin.WorkerStats")
}

func (e *FilRPC) FilGetMarketListIncompleteDeals() (*[]storagemarket.MinerDeal, error) {
	return e.getMarketListIncompleteDeals("Filecoin.MarketListIncompleteDeals")
}

func (e *FilRPC) postStartApi(method string, params []byte) error {
	result, err := e.httprpc.CallDoBody(method, params)
	if err != nil {
		log.Errorf("postStartApi err:%v", err)
		return err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil
	}

	return nil
}
func (e *FilRPC) FilPostStartApi(payload []byte) error {
	return e.postStartApi("Filecoin.DealsImportDataOfSxx", payload)
}

func (e *FilRPC) postStartP1(method string, params []byte) error {
	result, err := e.httprpc.CallDoBody(method, params)
	if err != nil {
		log.Errorf("postStartP1 err:%v", err)
		return err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil
	}

	return nil
}
func (e *FilRPC) FilPostStartP1(payload []byte) error {
	return e.postStartP1("Filecoin.SectorsUpdateOfSxx", payload)
}

func (e *FilRPC) postStartC1(method string, params []byte) error {
	result, err := e.httprpc.CallDoBody(method, params)
	if err != nil {
		log.Errorf("postStartP2 err:%v", err)
		return err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil
	}

	return nil
}
func (e *FilRPC) FilPostStartC1(payload []byte) error {
	return e.postStartC1("Filecoin.SectorsUpdateOfSxx", payload)
}

func (e *FilRPC) postWriteBack(method string, params []byte) error {
	result, err := e.httprpc.CallDoBody(method, params)
	if err != nil {
		log.Errorf("postStartP2 err:%v", err)
		return err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil
	}

	return nil
}
func (e *FilRPC) FilPostWriteBack(payload []byte) error {
	return e.postWriteBack("Filecoin.SectorsUpdateOfSxx", payload)
}

func (e *FilRPC) getMinerGetBaseInfo(method string, params ...interface{}) error {
	result, err := e.httprpc.Call(method, params...)
	if err != nil {
		log.Errorf("getMinerGetBaseInfo err:%v", err)
		return err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil
	}
	log.Infof("getMinerGetBaseInfo result:%v", string(result))

	return nil
}
func (e *FilRPC) FilMinerGetBaseInfo(params ...interface{}) error {
	return e.getMinerGetBaseInfo("Filecoin.MinerGetBaseInfo", params...)
}

func (e *FilRPC) postWalletSign(method string, params ...interface{}) error {
	result, err := e.httprpc.CallDo(method, params...)
	if err != nil {
		log.Errorf("WalletSign err:%v", err)
		return err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil
	}

	return nil
}
func (e *FilRPC) FilWalletSign(params ...interface{}) error {
	return e.postWalletSign("Filecoin.WalletSign", params...)
}
