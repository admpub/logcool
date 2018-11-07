//Filter-plug: grok
//grok is regexp plugin and for filter data you like.
package grok

import (
	"regexp"

	"github.com/admpub/logcool/utils"
)

const (
	ModuleName = "grok"
)

// Define zeus' config.
type FilterConfig struct {
	utils.FilterConfig
	Match string `json:"match"`
	Model string `json:"model"`
}

func init() {
	utils.RegistFilterHandler(ModuleName, InitHandler)
}

// Init grok Handler.
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

	tfc = &conf
	return
}

// Filter's event,and this is the main function of filter.
func (fc *FilterConfig) Event(event utils.LogEvent) utils.LogEvent {
	if event.Extra == nil {
		event.Extra = make(map[string]interface{})
	}
	re := regexp.MustCompile(fc.Match)
	value := re.FindString(event.Message)
	if len(value) > 0 {
		if fc.Model == "over" {
			event.Message = value
		} else {
			event.Extra["data"] = event.Format(value)
		}
	}
	return event
}
