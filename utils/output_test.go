package utils_test

import (
	"context"
	"fmt"
	"testing"

	outputstdout "github.com/admpub/logcool/output/stdout"
	. "github.com/admpub/logcool/utils"
)

func Test_RegistOutputHandler(t *testing.T) {
	RegistOutputHandler("stdout", outputstdout.InitHandler)
}

func Test_RunOutputs(t *testing.T) {
	config, err := LoadFromString(context.Background(), `
	{
		"input": [{
			"type": "file",
			"path": "/tmp/log/log.log",
			"sincedb_path": "",
			"start_position": "beginning"
		}]
	}
	`)
	RegistOutputHandler("stdout", outputstdout.InitHandler)
	if err != nil {
		fmt.Println(err)
	}
	config.RunOutputs()
}
