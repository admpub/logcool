package utils_test

import (
	"context"
	"fmt"
	"testing"

	stdininput "github.com/admpub/logcool/input/stdin"
	. "github.com/admpub/logcool/utils"
)

func Test_RegistInputHandler(t *testing.T) {
	RegistInputHandler("stdin", stdininput.InitHandler)
}

func Test_RunInputs(t *testing.T) {
	config, err := LoadFromString(context.Background(), `
	{
		"input": [{
			"type": "file",
			"path": "./tmp/log/log.log",
			"sincedb_path": "",
			"start_position": "beginning"
		}]
	}
	`)
	RegistInputHandler("stdin", stdininput.InitHandler)
	if err != nil {
		fmt.Println(err)
	}
	config.RunInputs()
}
