package utils_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/admpub/logcool/filter/zeus"
	. "github.com/admpub/logcool/utils"
)

func Test_RegistFilterHandler(t *testing.T) {
	RegistFilterHandler("zeus", zeus.InitHandler)
}

func Test_RunFilters(t *testing.T) {
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
	RegistFilterHandler("zeus", zeus.InitHandler)
	if err != nil {
		fmt.Println(err)
	}
	config.RunFilters()
}
