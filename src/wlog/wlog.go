package wlog

import (
    "log"
    "os"
    "strings"
)

var (
    Logging wlog
)

const (
    DEBUG = 0
    INFO = 1
    WARN = 2
    ERROR = 3
)

type wlog struct {
    logger * log.Logger
    detail * log.Logger
    level int
}

func (clog *wlog)init(){
    clog.logger = log.New(os.Stdout,"",log.Ldate | log.Ltime)
    clog.detail = log.New(os.Stdout,"",log.Ldate | log.Ltime | log.Lshortfile)
    clog.level = DEBUG
}

func (clog *wlog)SetLevel(level string){
    level = strings.ToLower(level)

    switch {
        case level == "debug":
            clog.level = DEBUG
        case level == "info":
            clog.level = INFO
        case level == "warn":
            clog.level = WARN
        case level == "error":
            clog.level = ERROR
        default:
            clog.level = INFO
    }
}

func (clog *wlog)log(level int, format string, v ...interface{}){
    if level < clog.level {
        return
    }

    switch {
        case level == DEBUG:
             clog.logger.Printf("[DEBUG] " + format, v...)
        case level == INFO:
             clog.logger.Printf("[Info ] " + format, v...)
        case level == WARN:
             clog.logger.Printf("[Warn ] " + format, v...)
        case level == ERROR:
             clog.logger.Printf("[Error] " + format, v...)
        default:
             clog.logger.Printf("[Info ] " + format, v...)
    }
}

func (clog *wlog)Detail(format string, v ...interface{}){
    clog.detail.Printf("[Detail]" + format, v...)
}

func (clog *wlog)Debug(format string, v ...interface{}){
    clog.log(DEBUG, format, v...)
}

func (clog *wlog)Info(format string, v ...interface{}){
    clog.log(INFO, format, v...)
}

func (clog *wlog)Warn(format string, v ...interface{}){
    clog.log(WARN, format, v...)
}

func (clog *wlog)Error(format string, v ...interface{}){
    clog.log(ERROR, format, v...)
}

func init(){
    Logging.init()
}

