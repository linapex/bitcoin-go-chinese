
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:52</date>
//</624461727534157824>

//版权所有（c）2015-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package main

import (
	"encoding/hex"
	"errors"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/database"
)

//fetchblockCmd定义fetchblock命令的配置选项。
type fetchBlockCmd struct{}

var (
//fetchblockcfg定义命令的配置选项。
	fetchBlockCfg = fetchBlockCmd{}
)

//执行是命令的主要入口点。它由解析器调用。
func (cmd *fetchBlockCmd) Execute(args []string) error {
//设置全局配置选项并确保它们有效。
	if err := setupGlobalConfig(); err != nil {
		return err
	}

	if len(args) < 1 {
		return errors.New("required block hash parameter not specified")
	}
	blockHash, err := chainhash.NewHashFromStr(args[0])
	if err != nil {
		return err
	}

//加载块数据库。
	db, err := loadBlockDB()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.View(func(tx database.Tx) error {
		log.Infof("Fetching block %s", blockHash)
		startTime := time.Now()
		blockBytes, err := tx.FetchBlock(blockHash)
		if err != nil {
			return err
		}
		log.Infof("Loaded block in %v", time.Since(startTime))
		log.Infof("Block Hex: %s", hex.EncodeToString(blockBytes))
		return nil
	})
}

//用法覆盖命令的用法显示。
func (cmd *fetchBlockCmd) Usage() string {
	return "<block-hash>"
}

