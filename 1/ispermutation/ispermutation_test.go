package ispermutation

import "testing"

func TestIsPermutation(t *testing.T) {
	if ! IsPermutation("david", "divad"){
		t.Fail()
	}
}

func TestIsPermutation2(t *testing.T) {
        if IsPermutation("david", "colombia"){
                t.Fail()
        }
}

func TestIsPermutation3(t *testing.T) {
        if IsPermutation("david", "dibad"){
                t.Fail()
        }
}

func TestIsPermutation4(t *testing.T) {
        if ! IsPermutation("david", "ddiav"){
                t.Fail()
        }
}
