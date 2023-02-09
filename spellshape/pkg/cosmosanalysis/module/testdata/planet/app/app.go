package app

import (
	"github.com/cosmos/cosmos-sdk/api/tendermint/abci"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/tendermint/planet/x/mars"
)

type Foo struct {
	FooKeeper foo.keeper
}

var ModuleBasics = module.NewBasicManager(mars.AppModuleBasic{})

func (f Foo) Name() string { return app.BaseApp.Name() }
func (f Foo) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (f Foo) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}
func (f Foo) RegisterAPIRoutes()         {}
func (f Foo) RegisterTxService()         {}
func (f Foo) RegisterTendermintService() {}
