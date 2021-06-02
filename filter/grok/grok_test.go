package grok

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/admpub/logcool/utils"
)

func init() {
	utils.RegistFilterHandler(ModuleName, InitHandler)
}

func Test_InitHandler(t *testing.T) {
	config := utils.ConfigRaw{}
	co, err := InitHandler(&config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(co)
}

func Test_Event(t *testing.T) {
	conf, err := utils.LoadFromString(context.Background(), `{
		"filter": [{
			"type": "grok",
            "macth":"2016-09",
            "model":"over"
		}]
	}`)
	if err != nil {
		fmt.Println(err)
	}

	timestamp := time.Now()

	inchan := conf.Get(reflect.TypeOf(make(utils.InChan))).
		Interface().(utils.InChan)

	outchan := conf.Get(reflect.TypeOf(make(utils.OutChan))).
		Interface().(utils.OutChan)

	err = conf.RunFilters()

	inchan <- utils.LogEvent{
		Timestamp: timestamp,
		Message:   "filter test message",
	}

	event := <-outchan
	fmt.Println(event)
}
