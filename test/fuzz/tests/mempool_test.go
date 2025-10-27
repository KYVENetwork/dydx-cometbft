//go:build gofuzz || go1.21

package tests

import (
	"testing"

	abciclient "github.com/KYVENetwork/dydx-cometbft/abci/client"
	"github.com/KYVENetwork/dydx-cometbft/abci/example/kvstore"
	"github.com/KYVENetwork/dydx-cometbft/config"
	cmtsync "github.com/KYVENetwork/dydx-cometbft/libs/sync"
	mempool "github.com/KYVENetwork/dydx-cometbft/mempool"
)

func FuzzMempool(f *testing.F) {
	app := kvstore.NewInMemoryApplication()
	mtx := new(cmtsync.Mutex)
	conn := abciclient.NewLocalClient(mtx, app)
	err := conn.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp := mempool.NewCListMempool(cfg, conn, 0)

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.CheckTx(data, nil, mempool.TxInfo{})
	})
}
