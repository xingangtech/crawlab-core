package utils

import (
	"fmt"
	"github.com/crawlab-team/crawlab-core/sys_exec"
	"github.com/crawlab-team/go-trace"
	"github.com/spf13/viper"
)

func GetApiAddress() (res string) {
	apiAddress := viper.GetString("api.address")
	if apiAddress == "" {
		return "http://localhost:8000"
	}
	return apiAddress
}

func ImportDemo() (err error) {
	cmdStr := fmt.Sprintf("python -m crawlab-demo import -a %s", GetApiAddress())
	cmd := sys_exec.BuildCmd(cmdStr)
	if err := cmd.Run(); err != nil {
		trace.PrintError(err)
	}
	return nil
}

func ReimportDemo() (err error) {
	cmdStr := fmt.Sprintf("python -m crawlab-demo reimport -a %s", GetApiAddress())
	cmd := sys_exec.BuildCmd(cmdStr)
	if err := cmd.Run(); err != nil {
		trace.PrintError(err)
	}
	return nil
}