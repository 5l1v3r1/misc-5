package key

import (
	"encoding/hex"
	"testing"
)

func TestPublicKeyEncode(t *testing.T) {
	priv := new(PrivateKey)
	priv.secret, _ = hex.DecodeString("18E14A7B6A307F426A94F8114701E7C8E774E7F9A47E2C2035DB29A206321725")
	pub := priv.ToPublicKey()
	address := pub.Encode()
	expectedAddress := "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM"
	if address != expectedAddress {
		t.Errorf("Incorrect address generated from private key: %s instead of %s", address, expectedAddress)
	}
}

func TestNew(t *testing.T) {
	_, _, err := New()
	if err != nil {
		t.Error("Error while creating keypair")
	}

}

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		New()
	}
}
