package utils

import (
	"testing"
)

func TestKeysFormat(t *testing.T) {
	t.Log("### TestKeysFormat ###")

	GenerateKeys()

	for i := 0; i < 5; i++ {
		s := <-NewKey
		if len(s) == 32 {
			t.Logf("Generated key: %q\n", s)
		} else {
			t.Error("The generated key violated the keyformat")
		}
	}
}

// Generate API keys. Test uniqueness of keys
// Pre-condition: none
// Post-condition: n unique keys are generated
func TestGenerateKeys(t *testing.T) {
	t.Log("### TestGenerateKeys ###")

	keys := make(map[string]bool)
	sem := make(chan struct{}, 1)
	done := make(chan struct{})

	GenerateKeys()

	nRoutines := 100
	nKeys := 200
	for i := 0; i < nRoutines; i++ {
		go func() {
			for j := 0; j < nKeys; j++ {
				s := <-NewKey
				sem <- struct{}{}
				if keys[s] {
					t.Errorf("%q was not unique!", s)
				} else {
					keys[s] = false
				}
				<-sem
			}
			done <- struct{}{}
		}()
	}

	for i := 0; i < 100; i++ {
		<-done
	}

	if len(keys) < (nRoutines * nKeys) {
		t.Error("Not enough keys generated, which means some wasn't unique")
	}
}

func BenchmarkGenerateKey(b *testing.B) {
	GenerateKeys()
	for i := 0; i < b.N; i++ {
		<-NewKey
	}
}
