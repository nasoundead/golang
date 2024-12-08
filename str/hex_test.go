package str

import (
	"testing"
)

func Test_BytesToHex(t *testing.T) {
	resStr := BytesToHexString([]byte("abc"))
	t.Log(resStr)
	expected := "616263"
	if resStr != expected {
		t.Errorf("[Input]abc[Output]Expedted %s, but got %s", expected, resStr)
	}

}
