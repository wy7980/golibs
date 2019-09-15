package main

import "wlog"

func main() {
    wlog.SetLevel("debug")
    //wlog.Detail("测试日志detail: %s","http://www.golang.org")
    wlog.Debug("测试日志debug: %s","http://www.golang.org")
    wlog.Info("测试日志info: %s","http://www.golang.org")
    wlog.Warn("测试日志warn: %s","http://www.golang.org")
    wlog.Error("测试日志error: %s","http://www.golang.org")
}

