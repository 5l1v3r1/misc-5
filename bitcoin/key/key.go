package key

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

//PrivateKey is the 256bit secret of a secp256k1 elliptic curve
type PrivateKey struct {
	secret []byte
}

//PublicKey of a secp256k1 keypair
type PublicKey struct {
	x *big.Int
	y *big.Int
}

//New generates a new secp256k1 public/private key pair
func New() (priv PrivateKey, pub PublicKey, err error) {
	priv.secret, pub.x, pub.y, err = elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	return
}
