
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:49</date>
//</624461713646817280>

//版权所有（c）2013-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package blockchain

import (
	"github.com/btcsuite/btclog"
)

//日志是一个没有输出过滤器初始化的日志程序。这个
//意味着在调用方之前，包默认不会执行任何日志记录
//请求它。
var log btclog.Logger

//默认的日志记录量为“无”。
func init() {
	DisableLog()
}

//DisableLog禁用所有库日志输出。日志记录输出被禁用
//默认情况下，直到调用uselogger。
func DisableLog() {
	log = btclog.Disabled
}

//uselogger使用指定的记录器输出包日志信息。
func UseLogger(logger btclog.Logger) {
	log = logger
}

