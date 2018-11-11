package influxdb

import (
	"context"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/admpub/logcool/utils"
	client "github.com/influxdata/influxdb/client/v2"
	"github.com/webx-top/echo"
)

var (
	DefaultBulkSize      uint64 = 10000
	DefaultWriteInterval uint   = 300 //seconds
)

const (
	ModuleName = "influxdb"
)

// OutputConfig Define outputstdout' config.
type OutputConfig struct {
	utils.OutputConfig
	WriteInterval int64    `json:"writeInterval"` //seconds
	BulkSize      uint64   `json:"bulkSize"`
	DSN           string   `json:"dsn"` //influx data source
	Database      string   `json:"database"`
	Table         string   `json:"table"`
	Precision     string   `json:"precision"`
	Tags          []string `json:"tags"`
	// RetentionPolicy is the retention policy of the points.
	RetentionPolicy string `json:"retentionPolicy"`
	// Write consistency is the number of servers required to confirm write.
	WriteConsistency string `json:"writeConsistency"`

	client        client.Client
	batchPoints   client.BatchPoints
	writeInterval time.Duration
}

func init() {
	utils.RegistOutputHandler(ModuleName, InitHandler)
}

// InitHandler Init Handler.
func InitHandler(ctx context.Context, confraw *utils.ConfigRaw) (retconf utils.TypeOutputConfig, err error) {
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

	//influxdb dsn: http://localhost:8086@admpub@123456@log-analysis-db@s
	dsnSli := strings.Split(conf.DSN, "@")

	// Create a new HTTPClient
	conf.httpClient, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     dsnSli[0],
		Username: dsnSli[1],
		Password: dsnSli[2],
	})
	if err != nil {
		return
	}
	if len(conf.Database) == 0 && len(dsnSli) > 3 {
		conf.Database = dsnSli[3]
	}
	if len(conf.Precision) == 0 && len(dsnSli) > 4 {
		conf.Database = dsnSli[4]
	}
	if conf.BulkSize == 0 {
		conf.BulkSize = DefaultBulkSize
	}
	if conf.WriteInterval <= 0 {
		conf.WriteInterval = DefaultWriteInterval
	}
	conf.writeInterval = time.Duration(conf.WriteInterval) * time.Second
	if len(conf.Tags) == 0 {
		conf.Tags = []string{
			`Path`, `Method`, `Scheme`, `StatusCode`,
		}
	}
	// Create a new point batch
	conf.batchPoints, err = client.NewBatchPoints(client.BatchPointsConfig{
		Database:         conf.Database,
		Precision:        conf.Precision,
		RetentionPolicy:  conf.RetentionPolicy,
		WriteConsistency: conf.WriteConsistency,
	})
	if err != nil {
		return
	}
	go conf.flush(ctx)
	retconf = &conf
	return
}

// Event Input's event,and this is the main function of output.
func (oc *OutputConfig) Event(ctx context.Context, event utils.LogEvent) (err error) {
	err = oc.addPoint(event.Extra)
	return
}

func (oc *OutputConfig) addPoint(message echo.Store) (err error) {
	fields := make(echo.Store)
	for k, v := range message {
		fields[k] = v
	}

	tags := map[string]string{}
	for _, tag := range oc.Tags {
		tags[tag] = fields.Get(tag).String()
	}
	fields.Delete(oc.Tags...)
	table := event.Format(oc.Table)
	var pt *client.Point
	pt, err = client.NewPoint(table, tags, fields, v.TimeLocal)
	if err != nil {
		return
	}
	oc.batchPoints.AddPoint(pt)
	// if the number of batch points are less than the batch size then we don't need to write them yet
	if len(oc.batchPoints.Points()) < oc.BulkSize {
		return nil
	}
	return oc.writePoints()
}

func (oc *OutputConfig) writePoints() (err error) {
	// Write the batch
	err = conf.client.Write(conf.batchPoints)
	return
}

func (oc *OutputConfig) Close() {
	if conf.client == nil {
		return
	}
	// Close client resources
	err := conf.client.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func (oc *OutputConfig) flush(ctx context.Context) {
	lock := sync.Mutex{}
	write := func() error {
		lock.Lock()
		err := oc.writePoints()
		lock.Unlock()
		if err != nil {
			log.Errorln(err)
		}
	}
	for {
		select {
		case <-ctx.Done():
			write()
			oc.Close()
			return
		default:
			time.Sleep(oc.writeInterval)
			if write() != nil {
				return
			}
		}
	}
	return
}
