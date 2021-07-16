package config

/**
 *	日志配置
 */

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type RedefineLogFormatter struct {
	//	日志时间
	Time string
	//	项目/服务名称
	Name string
	//	日志级别
	Level string
	//	协程名称
	Thread string
	//	traceId
	TraceId string
	//	spanId
	SpanId string
}

const (
	InfoLevel  = "info"
	DebugLevel = "debug"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

func init() {

	dateTime := time.Now().Format("2006-01-02 15:04:05.000")
	//	pid := syscall.Getegid()
	rlf := &RedefineLogFormatter{
		Name: "TestDemo",
		//	设置日志级别信息
		Level: InfoLevel,
		Time:  dateTime,
	}
	//	设置日志级别
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(rlf)
	log.SetReportCaller(true)
}

func (rlf *RedefineLogFormatter) Format(entry *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// entry.Message 就是需要打印的日志
	b.WriteString(fmt.Sprintf("%s\t%s\t%s\t-->:\t%s\n", rlf.Time, rlf.Name, rlf.Level, entry.Message))
	return b.Bytes(), nil
}
