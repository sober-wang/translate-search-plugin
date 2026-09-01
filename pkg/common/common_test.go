package common

import (
	"testing"
)

//func TestReadConfig(t *testing.T) {
//	cfg, err := ReadConfig()
//	if err != nil {
//		t.Error(err)
//	}
//	t.Logf("\nserverName: %s\naddress: %s\ninterface: %s", cfg.ServiceName, cfg.ObjectPath, cfg.Interface)
//
//}

func TestIncludeMathOperator(t *testing.T) {
	strs := []string{
		"清醒的",
		"sober",
		"sober*wang",
		"ws1992jx@163.com",
		"$",
		"*++--*==",
		"sober(wang)",
		"100%",
	}
	for _, v := range strs {
		if IncludeMathOperator(v) {
			t.Logf("str: %s", v)
		}
	}
}
