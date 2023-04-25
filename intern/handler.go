package intern

import (
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/jhyehuang/fil-cmd/intern/models/model"
	"github.com/jhyehuang/fil-cmd/pkg/log"
	"github.com/labstack/echo/v4"
	"net/http"
)

func serveProfile(c echo.Context) error {
	return nil
}

func (sm *ServiceManage) GetBlockReward(c echo.Context) error {

	miner := c.QueryParam("miner")
	heigh := c.QueryParam("epoch")

	var blockResp []GetBlockRewardResponse

	log.Errorf("search miner %+v", miner)
	log.Errorf("search heigh %+v", heigh)
	var minerHeader []model.BlockHeader
	if err := sm.DB.Debug().Where("miner = ? and height > ? ORDER BY height asc limit 100", miner, heigh).Find(&minerHeader).Error; err != nil {
		log.Errorf("search BlockHeader %+v", err)
		return err
	}
	var chainReward []model.ChainReward
	for _, v := range minerHeader {
		one_block := GetBlockRewardResponse{
			Height:  v.Height,
			MinerID: v.Miner,
		}

		var reward model.ChainReward
		if err := sm.DB.Where("height = ?", v.Height).First(&reward).Error; err != nil {
			log.Errorf("search ChainReward %+v", err)
			return err
		}
		chainReward = append(chainReward, reward)
		NewReward, err := big.FromString(reward.NewReward)
		if err != nil {
			log.Errorf("search NewReward %+v", err)
			return err
		}
		one_block.Reward = big.Mul(big.Div(NewReward, big.NewInt(builtin.ExpectedLeadersPerEpoch)), big.NewInt(*v.WinCount))

		var block_msg []model.BlockMessage
		if err := sm.DB.Debug().Where("block = ?", v.Cid).Find(&block_msg).Error; err != nil {
			log.Errorf("search Message %+v", err)
			return err
		}
		if len(block_msg) == 0 {
			blockResp = append(blockResp, one_block)
			continue
		}
		var cid []string
		for _, v := range block_msg {
			cid = append(cid, v.Message)
		}
		log.Infof("cid %+v", cid)

		var total_gas []string
		sqlstr := `select COALESCE(SUM(miner_tip),0)  as total_gas from derived_gas_outputs where cid in  ?  `
		if err := sm.DB.Debug().Raw(sqlstr, cid).Pluck("total_gas", &total_gas).Error; err != nil {
			log.Errorf("search DerivedGasOutput %+v", err)
			return err
		}

		TotalGas, err := big.FromString(total_gas[0])
		if err != nil {
			log.Errorf("search TotalGas %+v", err)
			return err
		}
		one_block.Reward = big.Add(one_block.Reward, TotalGas)
		blockResp = append(blockResp, one_block)
	}

	return c.JSON(http.StatusOK, blockResp)
}

func (sm *ServiceManage) GetMinerBlockCount(c echo.Context) error {

	miner := c.QueryParam("miner")

	var blockResp []GetMinerBlockCountResponse

	log.Errorf("search miner %+v", miner)

	return c.JSON(http.StatusOK, blockResp)
}
