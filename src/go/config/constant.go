package config

import (
	"fmt"
	"time"
)

/*
author Lei Wang
可在此文件定义整个项目会用到的常量；当然，并不强制这么做
*/

// TimeLayoutStr 时间初始化常量，Go语言中必须使用此常量初始化时间，否则可能会出错
const TimeLayoutStr = "2006-01-02 15:04:05"

// formatExample 此函数仅展示如何使用时间初始化常量进行时间初始化
func formatExample() {
	fmt.Println(time.Now().Format(TimeLayoutStr))
}
