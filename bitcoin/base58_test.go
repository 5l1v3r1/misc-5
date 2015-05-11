package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestDecodeBase58(t *testing.T) {
	address := " 1DTjvhLV6S72NQrSDrCX1GTCb9B3D5pmCB "
	expected := "0088b028348642ad1bbaa8fcc054273070eda045fe"

	result, match := DecodeBase58Check(address)
	if !match {
		t.Error("Invalid checksum matching")
	}

	hexEncodedResult := hex.EncodeToString(result)

	if expected != hexEncodedResult {
		t.Errorf("Decoded address is %s instead of %s", hexEncodedResult, expected)
	}

}

func TestEncodeBase58Check(t *testing.T) {

	data, _ := hex.DecodeString("0088b028348642ad1bbaa8fcc054273070eda045fe")
	address := "1DTjvhLV6S72NQrSDrCX1GTCb9B3D5pmCB"

	encoded := EncodeBase58Check(data)
	if encoded != address {
		t.Errorf("Encoded address is %s instead of %s", encoded, address)
	}
}
