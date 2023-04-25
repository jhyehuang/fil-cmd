package rpc

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/jhyehuang/fil-cmd/intern/log"
	"testing"
)

func Test_FilGetTipsetByHeight(t *testing.T) {
	nodeAddress := "http://"

	crawler := New(nodeAddress)
	result, err := crawler.FilGetTipsetByHeight(1374715)
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", result.Cids)
	log.Infof("-------Height: %d", result.Height)
	for _, l := range result.Blocks {
		log.Infof("------- %+v", l.ForkSignaling)
		log.Infof("------- %+v", l.BLSAggregate.Type)
		log.Infof("------- %+v", l.BLSAggregate.Data)
		log.Infof("------- %s", string(l.BLSAggregate.Data))
		log.Infof("------- %s", l.Miner.String())
		for _, b := range l.WinPoStProof {
			log.Infof("------- %x", b.PoStProof) // f033716 2329038 f036004 2329039 2329040
			str2 := string(b.ProofBytes[:])
			log.Infof("------- %s", str2)
		}

		//log.Infof("------- %X", l.ElectionProof.VRFProof)
		//log.Infof("------- Messages %+v", l.Messages)

	}
}

func Test_FilGetMarketListIncompleteDeals(t *testing.T) {
	nodeAddress := ""
	crawler := New(nodeAddress)
	result, err := crawler.FilGetMarketListIncompleteDeals()
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", result)

}

func Test_FilGetWorkerStats(t *testing.T) {
	nodeAddress := ""

	crawler := New(nodeAddress)
	result, err := crawler.FilGetWorkerStats()
	if err != nil {
		log.Errorf(err.Error())

	}
	log.Infof("------- %+v", result)

}

func Test_FilPostStartApi(t *testing.T) {
	nodeAddress := ""
	bodyStr := `{
		"method": "Filecoin.SectorsUpdateOfSxx",
		"id": 1,
		"params": [3, "Committing", "test-vm-12.test.office.sxxfuture.net"]
	}`
	var err error

	body := []byte(bodyStr)
	if err != nil {
		log.Errorf(err.Error())

	}

	crawler := New(nodeAddress)
	err = crawler.FilPostStartApi(body)
	if err != nil {
		log.Errorf(err.Error())

	}

}

func Test_FilPostStartP1(t *testing.T) {
	nodeAddress := "xxxxxxxxxxxx:/ip4/172.16.0.112/tcp/2345/http"
	bodyStr := `{
		"method": "Filecoin.SectorsUpdateOfSxx",
		"id": 1,
		"params": [3, "PreCommit1", "test-vm-12.test.office.sxxfuture.net"]
	}`
	var err error

	body := []byte(bodyStr)
	if err != nil {
		log.Errorf(err.Error())

	}

	crawler := New(nodeAddress)
	err = crawler.FilPostStartP1(body)
	if err != nil {
		log.Errorf(err.Error())

	}

}

func Test_FilPostStartC1(t *testing.T) {
	nodeAddress := ""
	bodyStr := `{
		"method": "Filecoin.SectorsUpdateOfSxx",
		"id": 1,
		"params": [3, "Committing", "test-vm-12.test.office.sxxfuture.net"]
	}`
	var err error

	body := []byte(bodyStr)
	if err != nil {
		log.Errorf(err.Error())

	}

	crawler := New(nodeAddress)
	err = crawler.FilPostStartC1(body)
	if err != nil {
		log.Errorf(err.Error())

	}

}

func Test_FilMinerGetBaseInfo(t *testing.T) {
	nodeAddress := ""
	crawler := New(nodeAddress)

	result, err := crawler.FilGetOriTipsetByHeight(2642038) //1380526
	if err != nil {
		log.Errorf(err.Error())

	}
	fmt.Printf("------- %+v\n", result.Key())
	addr, err := address.NewFromString("f01771695")
	if err != nil {
		log.Errorf(err.Error())
	}
	err = crawler.FilMinerGetBaseInfo(addr, 2642039, result.Key())
	if err != nil {
		log.Errorf(err.Error())

	}
}
