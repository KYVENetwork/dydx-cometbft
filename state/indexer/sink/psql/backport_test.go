package psql

import (
	"github.com/KYVENetwork/dydx-cometbft/state/indexer"
	"github.com/KYVENetwork/dydx-cometbft/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
