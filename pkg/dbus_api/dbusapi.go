/*
Package dbusapi D-Bus 库调用，定义API接口方法。
*/
package dbusapi

import (
	"encoding/json"
	log "log/slog"
	"syscall"

	calltranslator "github.com/sober-wang/translate-search-plugin/pkg/call_translator"
	"github.com/sober-wang/translate-search-plugin/pkg/types"

	"github.com/godbus/dbus/v5"
)

// buildDBusErrorMsg 构建 dbus 错误信息
func buildDBusErrorMsg(err error) *dbus.Error {
	return dbus.NewError("org.freedesktop.DBus.Error.InvalidArgs: %s", []any{err})
}

// GrandSearchPlugin 定义插件需要实现的结构体
// 它的方法将作为 D-Bus 接口的一部分被导出
type GrandSearchPlugin struct {
	ServiceName string
	ObjectPath  string
	Interface   string
	SelfPid     int
}

// Search 方法接收搜索请求
func (gsp *GrandSearchPlugin) Search(jsonData string) (string, *dbus.Error) {
	log.Info("收到搜索请求", "data", jsonData)

	// 1. 解析输入 JSON (包含 "ver", "mID", "cont" 等字段)[reference:10]
	var rb types.ReqBody
	//var request map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &rb); err != nil {
		return "", buildDBusErrorMsg(err)
	}

	var rpb types.RespBody
	rpb.MID = rb.MID
	rpb.Ver = rb.Ver
	res, bingerr := calltranslator.QueryBingDict(rb.Count)
	var results string
	if bingerr != nil {
		str, err := calltranslator.CallCliTransToChinese(rb.Count)
		if err != nil {
			results = rpb.BuildEmptyResp()
			log.Error("Translate", "remote_err", bingerr.Error(), "local_err", err.Error())
			return results, nil
		}
		results = rpb.BuildResp(str) // 这是你需要实现的搜索函数
		log.Info("答复请求", "data", results)
		return results, nil
	}

	// 2. 执行你的搜索逻辑...
	results = rpb.BuildResp(res) // 这是你需要实现的搜索函数

	// 3. 将结果序列化为 JSON 字符串并返回
	log.Info("答复请求", "data", results)

	return results, nil
}

// Stop 方法用于停止搜索
func (gsp *GrandSearchPlugin) Stop(jsonData string) (bool, *dbus.Error) {
	log.Info("收到停止请求", "data", jsonData)
	syscall.Kill(gsp.SelfPid, syscall.SIGKILL)
	// 实现停止搜索的逻辑...
	return true, nil
}

// Action 方法处理用户点击结果的操作
func (gsp *GrandSearchPlugin) Action(jsonData string) (bool, *dbus.Error) {
	log.Info("收到操作请求", "data", jsonData)
	// 实现打开文件/应用等的逻辑...
	return true, nil
}

// DBusRun 启动服务并注册
func DBusRun(gsp GrandSearchPlugin) {
	log.Info("grandSearcPlugin", "serviceName", gsp.ServiceName, "objectPath", gsp.ObjectPath, "interface", gsp.Interface)
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Error("Failed to connect to session bus:", "err", err)
		return
	}
	defer conn.Close()

	// 2. 注册服务名 (需要与 .conf 文件中的 DBusService 一致)
	reply, err := conn.RequestName(gsp.ServiceName, dbus.NameFlagDoNotQueue)
	if err != nil || reply != dbus.RequestNameReplyPrimaryOwner {
		log.Error("Failed to register service name:", "err", err)
		return
	}
	log.Info("register", "reply", reply)

	// 3. 导出对象 (对象路径需要与 .conf 文件中的 DBusAddress 一致)
	//plugin := &GrandSearchPlugin{}
	err = conn.Export(&gsp, dbus.ObjectPath(gsp.ObjectPath), gsp.Interface) // 接口名需要与 .conf 文件中的 DBusInterface 一致
	if err != nil {
		log.Error("Failed to export object:", "err", err)
		return
	}
	// 保持服务运行，必须在 dbus 注册的函数保持运行，否则 dbus 服务退出将无法调用
	select {}
}
