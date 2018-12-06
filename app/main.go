package app

import (
	"github.com/kooksee/hashnet/cmn"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

type KApp struct {
	valUpdates []types.ValidatorUpdate
	types.BaseApplication
	logger log.Logger
}

func New(logger log.Logger) *KApp {
	// 初始化mint模块
	return &KApp{logger: logger.With("module", "kapp")}
}

// 实现abci的Info协议
func (app *KApp) Info(req types.RequestInfo) (res types.ResponseInfo) {
	d, _ := cmn.JsonMarshalToString(req)
	app.logger.Info(d, "abci", "Info")

	res.Version = req.Version

	return
}

// 实现abci的SetOption协议
func (app *KApp) SetOption(req types.RequestSetOption) types.ResponseSetOption {
	return types.ResponseSetOption{Code: types.CodeTypeOK}
}

// 实现abci的CheckTx协议
func (app *KApp) CheckTx(txHash []byte) (res types.ResponseCheckTx) {
	return
}

// 实现abci的DeliverTx协议
func (app *KApp) DeliverTx(txBytes []byte) (res types.ResponseDeliverTx) {
	return
}

// Commit will panic if InitChain was not called
func (app *KApp) Commit() types.ResponseCommit {
	return types.ResponseCommit{}
}

func (app *KApp) Query(reqQuery types.RequestQuery) (res types.ResponseQuery) {
	return
}

// Save the validators in the merkle tree
func (app *KApp) InitChain(req types.RequestInitChain) types.ResponseInitChain {
	return types.ResponseInitChain{}
}

func (app *KApp) BeginBlock(req types.RequestBeginBlock) types.ResponseBeginBlock {
	d, _ := cmn.JsonMarshalToString(req)
	app.logger.Info(d, "abci", "BeginBlock")

	app.valUpdates = make([]types.ValidatorUpdate, 0)
	return types.ResponseBeginBlock{}
}

func (app *KApp) EndBlock(req types.RequestEndBlock) types.ResponseEndBlock {
	d, _ := cmn.JsonMarshalToString(req)
	app.logger.Info(d, "abci", "EndBlock")

	return types.ResponseEndBlock{ValidatorUpdates: app.valUpdates}
}
