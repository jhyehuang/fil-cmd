// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model


const TableNameChainReward = "chain_rewards"

// ChainReward mapped from table <chain_rewards>
type ChainReward struct {
	StateRoot                         string   `gorm:"column:state_root;type:text;primaryKey" json:"state_root"`                                                        // CID of the parent state root.
	CumSumBaseline                    string  `gorm:"column:cum_sum_baseline;type:numeric;not null" json:"cum_sum_baseline"`                                           // Target that CumsumRealized needs to reach for EffectiveNetworkTime to increase. It is measured in byte-epochs (space * time) representing power committed to the network for some duration.
	CumSumRealized                    string  `gorm:"column:cum_sum_realized;type:numeric;not null" json:"cum_sum_realized"`                                           // Cumulative sum of network power capped by BaselinePower(epoch). It is measured in byte-epochs (space * time) representing power committed to the network for some duration.
	EffectiveBaselinePower            string  `gorm:"column:effective_baseline_power;type:numeric;not null" json:"effective_baseline_power"`                           // The baseline power (in bytes) at the EffectiveNetworkTime epoch.
	NewBaselinePower                  string  `gorm:"column:new_baseline_power;type:numeric;not null" json:"new_baseline_power"`                                       // The baseline power (in bytes) the network is targeting.
	NewRewardSmoothedPositionEstimate string  `gorm:"column:new_reward_smoothed_position_estimate;type:numeric;not null" json:"new_reward_smoothed_position_estimate"` // Smoothed reward position estimate - Alpha Beta Filter "position" (value) estimate in Q.128 format.
	NewRewardSmoothedVelocityEstimate string  `gorm:"column:new_reward_smoothed_velocity_estimate;type:numeric;not null" json:"new_reward_smoothed_velocity_estimate"` // Smoothed reward velocity estimate - Alpha Beta Filter "velocity" (rate of change of value) estimate in Q.128 format.
	TotalMinedReward                  string  `gorm:"column:total_mined_reward;type:numeric;not null" json:"total_mined_reward"`                                       // The total FIL (attoFIL) awarded to block miners.
	NewReward                         string `gorm:"column:new_reward;type:numeric" json:"new_reward"`                                                                // The reward to be paid in per WinCount to block producers. The actual reward total paid out depends on the number of winners in any round. This value is recomputed every non-null epoch and used in the next non-null epoch.
	EffectiveNetworkTime              int64   `gorm:"column:effective_network_time;type:bigint" json:"effective_network_time"`                                         // Ceiling of real effective network time "theta" based on CumsumBaselinePower(theta) == CumsumRealizedPower. Theta captures the notion of how much the network has progressed in its baseline and in advancing network time.
	Height                            int64    `gorm:"column:height;type:bigint;primaryKey" json:"height"`                                                              // Epoch this rewards summary applies to.
}

// TableName ChainReward's table name
func (*ChainReward) TableName() string {
	return TableNameChainReward
}
