package bitcoin

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

//Hash256 return Bitcoin's 256-bit hash (double SHA-256).
func Hash256(b []byte) [32]byte {
	hash1 := sha256.Sum256(b)
	hash2 := sha256.Sum256(hash1[:])
	return hash2
}

//Hash160 returns Bitcoin's 160-bit hash(SHA-256 followed by RIPEMD-160)
func Hash160(b []byte) [20]byte {

	sha256hash := sha256.Sum256(b)
	ripemd160hasher := ripemd160.New()

	var ripemd160hash [20]byte
	ripemd160hasher.Write(sha256hash[:])
	copy(ripemd160hash[:], ripemd160hasher.Sum(nil))

	return ripemd160hash
}
