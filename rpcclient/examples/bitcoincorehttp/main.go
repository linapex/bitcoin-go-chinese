
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:55</date>
//</624461741182423040>

//版权所有（c）2014-2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package main

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
//使用HTTP Post模式连接到本地比特币核心RPC服务器。
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:8332",
		User:         "yourrpcuser",
		Pass:         "yourrpcpass",
HTTPPostMode: true, //比特币核心仅支持HTTP Post模式
DisableTLS:   true, //比特币核心默认不提供TLS
	}
//注意，通知参数为nil，因为通知是
//HTTP POST模式不支持。
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

//获取当前块计数。
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}

