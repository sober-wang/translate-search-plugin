package common

import "testing"

func TestReadConfig(t *testing.T) {
	cfg, err := ReadConfig()
	if err != nil {
		t.Error(err)
	}
	t.Logf("\nserverName: %s\naddress: %s\ninterface: %s", cfg.ServerName, cfg.Address, cfg.Interface)

}
