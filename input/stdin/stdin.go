// Input-plug: stdininput
// The plug's function input the data from the standard-input.
package stdininput

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Sirupsen/logrus"

	"github.com/admpub/logcool/utils"
)

const (
	ModuleName = "stdin"
)

// Define stdininput' config.
type InputConfig struct {
	utils.InputConfig

	hostname string
}

func init() {
	utils.RegistInputHandler(ModuleName, InitHandler)
}

// Init stdininput Handler.
func InitHandler(confraw *utils.ConfigRaw) (retconf utils.TypeInputConfig, err error) {
	conf := InputConfig{
		InputConfig: utils.InputConfig{
			CommonConfig: utils.CommonConfig{
				Type: ModuleName,
			},
		},
	}
	if err = utils.ReflectConfig(confraw, &conf); err != nil {
		return
	}

	if conf.hostname, err = os.Hostname(); err != nil {
		return
	}

	retconf = &conf
	return
}

// Input's start,and this is the main function of input.
func (t *InputConfig) Start() {
	t.Invoke(t.echo)
}

func (t *InputConfig) echo(logger *logrus.Logger, ctx context.Context, inchan utils.InChan) (err error) {
	defer func() {
		if err != nil {
			logger.Errorln(err)
		}
	}()

	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		select {
		case <-ctx.Done():
			close(inchan)
			return
		default:
			// Sleep some Nanoseconds wait for event have been deal.
			time.Sleep(300000 * time.Nanosecond)
			fmt.Print("logcool#")
			data, _, _ := reader.ReadLine()
			command := string(data)
			event := utils.LogEvent{
				Timestamp: time.Now(),
				Message:   command,
				Extra: map[string]interface{}{
					"host": t.hostname,
				},
			}
			inchan <- event
			if command == "quit" {
				os.Exit(0)
			}
		}
	}

	return
}
