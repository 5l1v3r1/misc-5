package key

import (
	"crypto/rand"
	"math/big"

	"encoding/hex"

	"github.com/titanous/bitcoin-crypto/bitelliptic"

	"github.com/robvanmieghem/misc/bitcoin"
)

//WIFEncoding is the version code for WIF encoded private keys
const WIFEncoding byte = 128

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

//EncodeToHex serializes to hexadecimal representation
func (priv *PrivateKey) EncodeToHex() (encoded string) {
	encoded = hex.EncodeToString(priv.secret)
	return
}

//EncodeToWIF serializes to Wallet Import Format
func (priv *PrivateKey) EncodeToWIF() (encoded string) {
	return bitcoin.EncodeBase58Check(priv.secret, WIFEncoding)
}

//EncodeToWIFCompressed serializes to compressed Wallet Import Format
func (priv *PrivateKey) EncodeToWIFCompressed() (encoded string) {
	const compressionFlag byte = 1
	secretWithCompressionFlag := append(priv.secret, compressionFlag)
	return bitcoin.EncodeBase58Check(secretWithCompressionFlag, WIFEncoding)
}

//PublicKey of a secp256k1 keypair
type PublicKey struct {
	x *big.Int
	y *big.Int
}

//Encode creates a Base58Check representation of a public key
func (pub *PublicKey) Encode() string {
	const version byte = 0x0

	marshalled := secp256k1.Marshal(pub.x, pub.y)
	hashed := bitcoin.Hash160(marshalled)
	return bitcoin.EncodeBase58Check(hashed[:], version)
}

//New generates a new secp256k1 public/private key pair
func New() (priv PrivateKey, pub PublicKey, err error) {
	priv.secret, pub.x, pub.y, err = secp256k1.GenerateKey(rand.Reader)
	return
}
