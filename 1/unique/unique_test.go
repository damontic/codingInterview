package unique

import (
	"testing"
)

func TestUnique(t *testing.T) {
	if HasAllUniqueCharacters("david"){
		t.FailNow()
	}
        if !HasAllUniqueCharacters("golan"){
                t.FailNow()
        }
}

func TestUniqueNoDS(t *testing.T) {
        if HasAllUniqueCharactersNoDS("david"){
                t.FailNow()
        }
        if !HasAllUniqueCharactersNoDS("golan"){
                t.FailNow()
        }
}
