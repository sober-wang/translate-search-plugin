/*
Packge common
公共处理包，读取文件，判断一些参数等
*/
package common

import (
	log "log/slog"

	dbusapi "github.com/sober-wang/translate-search-plugin/pkg/dbus_api"

	"gopkg.in/ini.v1"
)

// IsOnlyEnglishLetters 是否只包含英文单词
func IsOnlyEnglishLetters(str string) bool {
	if len(str) == 0 {
		return false // 根据业务需求，空字符串可改为 true
	}
	for i := 0; i < len(str); i++ {
		c := str[i]
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			return false
		}
	}
	return true
}

// ReadConfig 获取配置文件信息
func ReadConfig() (dbusapi.GrandSearchPlugin, error) {
	var cfg dbusapi.GrandSearchPlugin
	cfgPath := "/usr/lib/x86_64-linux-gnu/dde-grand-search-daemon/plugins/searcher/translate-search-plugin.conf"
	log.Info("config path", "path", cfgPath)
	f, err := ini.Load(cfgPath)
	if err != nil {
		return cfg, err
	}
	cfg.ServiceName = f.Section("Grand Search").Key("DBusService").String()
	cfg.ObjectPath = f.Section("Grand Search").Key("DBusAddress").String()
	cfg.Interface = f.Section("Grand Search").Key("DBusInterface").String()
	return cfg, nil
}
