//Filter-plug: metrics
//metrics is just for metrics and count.
package metrics

import (
	"encoding/json"
	"regexp"

	"github.com/admpub/logcool/utils"
)

const (
	ModuleName = "metrics"
)

// FilterConfig Define metrics' config.
type FilterConfig struct {
	utils.FilterConfig
	Tag   []string `json:"tag"`
	Alarm []int64  `json:"alarm"`

	metrics map[string]int64
}

func init() {
	utils.RegistFilterHandler(ModuleName, InitHandler)
}

// Init metrics Handler.
func InitHandler(confraw *utils.ConfigRaw) (tfc utils.TypeFilterConfig, err error) {
	conf := FilterConfig{
		FilterConfig: utils.FilterConfig{
			CommonConfig: utils.CommonConfig{
				Type: ModuleName,
			},
		},
	}
	// Reflect config from configraw.
	if err = utils.ReflectConfig(confraw, &conf); err != nil {
		return
	}
	conf.metrics = make(map[string]int64)

	tfc = &conf

	return
}

// Filter's event,and this is the main function of filter.
func (fc *FilterConfig) Event(event utils.LogEvent) utils.LogEvent {
	if event.Extra == nil {
		event.Extra = make(map[string]interface{})
	}

	for _, value := range fc.Tag {
		re := regexp.MustCompile(value)
		isv := re.FindString(event.Message)
		if isv != "" {
			fc.metrics[isv]++
		}
	}

	metrics, _ := json.Marshal(fc.metrics)
	event.Extra["metrics"] = event.Format(string(metrics))

	return event
}
