package log

import (
	gocontext "context"
	"sync"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/leantech/school-system-api/config"
	"github.com/leantech/school-system-api/context"
	"github.com/sirupsen/logrus"
)

type key string

func (k key) String() string {
	return "key: " + string(k)
}

var (
	log  *logrus.Logger
	once sync.Once
)

const (
	serviceName string = "school-system-api"
)

func Init() {
	once.Do(func() {
		logConfig := config.GetEnv().Log
		log = logrus.New()
		log.Formatter = stackdriver.NewFormatter(
			stackdriver.WithService(serviceName),
		)
		if level, err := logrus.ParseLevel(logConfig.Level); err == nil {
			log.SetLevel(level)
		}
	})
}

func InitParams(ctx gocontext.Context) gocontext.Context {

	httpLog := new(HTTP)
	httpLog.Request = new(Request)
	httpLog.Response = new(Response)

	ctx = context.Set(ctx, HTTPKey, httpLog)

	return ctx
}

func NewEntry() *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"mutex": &sync.Mutex{},
		"type":  "json",
	})
}
