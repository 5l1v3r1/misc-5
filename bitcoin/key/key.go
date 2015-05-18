package key

import (
	"crypto/rand"
	"math/big"

	"github.com/titanous/bitcoin-crypto/bitelliptic"

	"github.com/robvanmieghem/misc/bitcoin"
)

var secp256k1 = bitelliptic.S256()

//PrivateKey is the 256bit secret of a secp256k1 elliptic curve
type PrivateKey struct {
	secret []byte
}

//ToPublicKey creates the public key from a private key
func (priv *PrivateKey) ToPublicKey() (pub PublicKey) {
	pub.x, pub.y = secp256k1.ScalarBaseMult(priv.secret)
	return
}

//PublicKey of a secp256k1 keypair
type PublicKey struct {
	x *big.Int
	y *big.Int
}

//Marshal converts a public key into the form specified in section 4.3.6 of ANSI X9.62.
func (pub *PublicKey) marshal() []byte {
	return secp256k1.Marshal(pub.x, pub.y)
}

//Encode creates a Base58Check representation of a public key
func (pub *PublicKey) Encode() string {
	const version byte = 0x0

	marshalled := pub.marshal()
	hashed := bitcoin.Hash160(marshalled)
	return bitcoin.EncodeBase58Check(hashed[:], version)
}

//New generates a new secp256k1 public/private key pair
func New() (priv PrivateKey, pub PublicKey, err error) {
	priv.secret, pub.x, pub.y, err = secp256k1.GenerateKey(rand.Reader)
	return
}
