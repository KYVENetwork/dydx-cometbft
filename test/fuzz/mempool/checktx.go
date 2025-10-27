package reactor

import (
	"github.com/KYVENetwork/dydx-cometbft/abci/example/kvstore"
	"github.com/KYVENetwork/dydx-cometbft/config"
	mempl "github.com/KYVENetwork/dydx-cometbft/mempool"
	"github.com/KYVENetwork/dydx-cometbft/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewInMemoryApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false
	mempool = mempl.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mempool.CheckTx(data, nil, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
