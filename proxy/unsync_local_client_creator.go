package proxy

import (
	abcicli "github.com/KYVENetwork/dydx-cometbft/abci/client"
	"github.com/KYVENetwork/dydx-cometbft/abci/types"
)

var (
	_ ClientCreator = (*unsynchronizedClientCreator)(nil)
)

// NewUnsynchronizedLocalClientCreator creates a local client that is unsynchronized. It is expected that the
// provided application perform all synchronization necessary to prevent unexpected results.
func NewUnsynchronizedLocalClientCreator(app types.Application) ClientCreator {
	return &unsynchronizedClientCreator{app: app}
}

type unsynchronizedClientCreator struct {
	app types.Application
}

func (u unsynchronizedClientCreator) NewABCIClient() (abcicli.Client, error) {
	return abcicli.NewUnsyncLocalClient(u.app), nil
}
