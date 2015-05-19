package key

import (
	"encoding/hex"
	"testing"
)

func TestPrivateKeyEncodeToHex(t *testing.T) {
	key := "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"

	priv := new(PrivateKey)
	priv.secret, _ = hex.DecodeString(key)
	hexencoded := priv.EncodeToHex()
	if hexencoded != key {
		t.Errorf("Wrongly hex encoded private key: %s instead of %s", hexencoded, key)
	}
}

func TestPrivateKeyEncodeToWIF(t *testing.T) {
	key := "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"
	wif := "5J3mBbAH58CpQ3Y5RNJpUKPE62SQ5tfcvU2JpbnkeyhfsYB1Jcn"

	priv := new(PrivateKey)
	priv.secret, _ = hex.DecodeString(key)
	encoded := priv.EncodeToWIF()
	if encoded != wif {
		t.Errorf("Wrongly WIF encoded private key: %s instead of %s", encoded, wif)
	}
}

func TestPrivateKeyEncodeToWIFCompressed(t *testing.T) {
	key := "1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"
	wifcompressed := "KxFC1jmwwCoACiCAWZ3eXa96mBM6tb3TYzGmf6YwgdGWZgawvrtJ"

	priv := new(PrivateKey)
	priv.secret, _ = hex.DecodeString(key)
	encoded := priv.EncodeToWIFCompressed()
	if encoded != wifcompressed {
		t.Errorf("Wrongly WIFCompressed encoded private key: %s instead of %s", encoded, wifcompressed)
	}
}

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
