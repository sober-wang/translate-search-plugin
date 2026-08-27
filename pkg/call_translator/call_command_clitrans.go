/*
Package calltranslator
CallCliTransToChinese 使用 clitrans 进行翻译
依赖：https://github.com/wfxr/clitrans
*/
package calltranslator

import (
	"errors"
	"fmt"
	log "log/slog"
	"os/exec"
)

// CallCliTransToChinese 使用 clitrans 命令翻译，https://github.com/wfxr/clitrans
func CallCliTransToChinese(str string) (string, error) {
	tranLine := fmt.Sprintf("clitrans %s", str)
	cmd := exec.Command("sh", "-c", tranLine)
	output, err := cmd.Output()
	if err != nil {
		log.Error("TransToChinese", "err", err.Error())
		return "", errors.New("not find definition")
	}
	return string(output), nil

}
