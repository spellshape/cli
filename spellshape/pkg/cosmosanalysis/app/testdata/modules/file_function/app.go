package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	foomodule "github.com/username/test/x/foo"
)

var ModuleBasics = module.NewBasicManager(basicModules()...)

func basicModules() {
	return []module.AppModuleBasic{
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		gov.NewAppModuleBasic([]govclient.ProposalHandler{
			paramsclient.ProposalHandler,
		}),
		foomodule.AppModuleBasic{},
	}
}

type Foo struct{}

func (Foo) Name() string {
	return "foo"
}

func (Foo) BeginBlocker(sdk.Context, abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return abci.ResponseBeginBlock{}
}

func (Foo) EndBlocker(sdk.Context, abci.RequestEndBlock) abci.ResponseEndBlock {
	return abci.ResponseEndBlock{}
}
