package key

import "testing"

func TestNewt(t *testing.T) {
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
