package intern

import "github.com/filecoin-project/go-state-types/big"

type GetBlockRewardRequest struct {
	Height int64 `json:"height"`
}

type GetBlockRewardResponse struct {
	MinerID string  `json:"miner_id"`
	Height  int64   `json:"height"`
	Reward  big.Int `json:"reward"`
}

type GetMinerBlockCountResponse struct {
	MinerID    string  `json:"miner_id"`
	BlockCount int64   `json:"block_count"`
	Reward     big.Int `json:"reward"`
}
