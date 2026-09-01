package main

import (
	"fmt"
	"io"
	log "log/slog"
	"os"
	"time"

	dbusapi "github.com/sober-wang/translate-search-plugin/pkg/dbus_api"
)

func init() {
	logPathName := fmt.Sprintf("/tmp/com.ddeGrandSearch.Translate_%s.log", time.Now().Format("20060102"))
	callLogFile, err := os.Create(logPathName)
	if err != nil {
		log.Error("open_logfile", "err", err.Error())
		return
	}
	mulWrt := io.MultiWriter(os.Stdout, callLogFile)
	handler := log.NewJSONHandler(mulWrt, &log.HandlerOptions{
		Level:     log.LevelDebug,
		AddSource: false,
	})
	log.SetDefault(log.New(handler))
	log.Default()
}

func main() {
	cfg, err := dbusapi.NewDBus()
	if err != nil {
		log.Error("read config", "err", err)
		return
	}

	dbusapi.DBusRun(cfg)

}
