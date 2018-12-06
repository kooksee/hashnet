package usdb

import (
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/kooksee/hashnet/cmn"
	"github.com/pingcap/tidb/store/tikv"
	"github.com/tendermint/tendermint/libs/log"
	"os"
	"time"
)

var Name = "txs"

var tdb *TikvStore

func Init(logger log.Logger) {
	tikv.MaxConnectionCount = 256

	url := os.Getenv("TIKV")
	if url == "" {
		panic("please init tikv url")
	}

	store, err := tikv.Driver{}.Open(fmt.Sprintf("tikv://%s/pd", url))
	cmn.MustNotErr("TikvStore Init Error", err)

	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Sprintf("init cache error: %s ", err.Error()))
	}

	tdb = &TikvStore{
		name:  []byte(Name),
		c:     store,
		cache: cache,
	}
}

func GetDb() *TikvStore {
	if tdb == nil {
		panic("please init usdb")
	}
	return tdb
}
