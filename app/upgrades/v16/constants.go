package v16

import (
	ratelimittypes "github.com/Stride-Labs/ibc-rate-limiting/ratelimit/types"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"

	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"

	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/gaia/v18/app/upgrades"
	"github.com/cosmos/gaia/v18/x/globalfee"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v16"

	RateLimitDenom         = "uatom"
	RateLimitDurationHours = 24
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			ratelimittypes.ModuleName,
			icacontrollertypes.SubModuleName,
			ibcfeetypes.ModuleName,
			feemarkettypes.ModuleName,
		},
		Deleted: []string{
			globalfee.ModuleName,
		},
	},
}