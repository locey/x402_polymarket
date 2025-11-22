package svc

import (
	"github.com/locey/x402_polymarket/PolymarketBase/evm/erc"
	//"github.com/ProjectsTask/PolymarketBase/image"
	"github.com/locey/x402_polymarket/PolymarketBase/stores/xkv"
	"gorm.io/gorm"

	"github.com/locey/x402_polymarket/PolymarketBackend/src/dao"
)

type CtxConfig struct {
	db *gorm.DB
	//imageMgr image.ImageManager
	dao     *dao.Dao
	KvStore *xkv.Store
	Evm     erc.Erc
}

type CtxOption func(conf *CtxConfig)

func NewServerCtx(options ...CtxOption) *ServerCtx {
	c := &CtxConfig{}
	for _, opt := range options {
		opt(c)
	}
	return &ServerCtx{
		DB: c.db,
		//ImageMgr: c.imageMgr,
		KvStore: c.KvStore,
		Dao:     c.dao,
	}
}

func WithKv(kv *xkv.Store) CtxOption {
	return func(conf *CtxConfig) {
		conf.KvStore = kv
	}
}

func WithDB(db *gorm.DB) CtxOption {
	return func(conf *CtxConfig) {
		conf.db = db
	}
}

func WithDao(dao *dao.Dao) CtxOption {
	return func(conf *CtxConfig) {
		conf.dao = dao
	}
}
