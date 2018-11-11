package stdout

import (
	"context"
	"fmt"

	"github.com/admpub/logcool/utils"
)

const (
	ModuleName = "stdout"
)

// OutputConfig Define outputstdout' config.
type OutputConfig struct {
	utils.OutputConfig
}

func init() {
	utils.RegistOutputHandler(ModuleName, InitHandler)
}

// InitHandler Init outputstdout Handler.
func InitHandler(confraw *utils.ConfigRaw) (retconf utils.TypeOutputConfig, err error) {
	conf := OutputConfig{
		OutputConfig: utils.OutputConfig{
			CommonConfig: utils.CommonConfig{
				Type: ModuleName,
			},
		},
	}
	if err = utils.ReflectConfig(confraw, &conf); err != nil {
		return
	}

	retconf = &conf
	return
}

// Event Input's event,and this is the main function of output.
func (oc *OutputConfig) Event(ctx context.Context, event utils.LogEvent) (err error) {
	raw, err := event.MarshalIndent()
	if err != nil {
		return
	}

	fmt.Println(string(raw))
	return
}
