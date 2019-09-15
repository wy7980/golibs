package wlog

func SetLevel(level string){
    Logging.SetLevel(level)
}

func Detail(format string, v ...interface{}){
    Logging.Detail(format, v...)
}

func Debug(format string, v ...interface{}){
    Logging.Debug(format, v...)
}

func Info(format string, v ...interface{}){
    Logging.Info(format, v...)
}

func Warn(format string, v ...interface{}){
    Logging.Warn(format, v...)
}

func Error(format string, v ...interface{}){
    Logging.Error(format, v...)
}
