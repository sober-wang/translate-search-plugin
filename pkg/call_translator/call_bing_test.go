package calltranslator

import "testing"

func TestQueryBingDict(t *testing.T) {
	str := "How are you?"
	zhStr, err := QueryBingDict(str)
	if err != nil {
		t.Error(err)
	}
	t.Log(zhStr)
}
