package sr25519

import cmtjson "github.com/KYVENetwork/dydx-cometbft/libs/json"

const (
	PrivKeyName = "dydxcometbft/PrivKeySr25519"
	PubKeyName  = "dydxcometbft/PubKeySr25519"
)

func init() {
	cmtjson.RegisterType(PubKey{}, PubKeyName)
	cmtjson.RegisterType(PrivKey{}, PrivKeyName)
}
