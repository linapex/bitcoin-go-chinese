
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:52</date>
//</624461727647404032>

//版权所有（c）2015-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package main

import (
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/database"
)

//BlockRegionCmd定义FetchBlockRegion的配置选项
//命令。
type blockRegionCmd struct{}

var (
//blockregioncfg定义命令的配置选项。
	blockRegionCfg = blockRegionCmd{}
)

//执行是命令的主要入口点。它由解析器调用。
func (cmd *blockRegionCmd) Execute(args []string) error {
//设置全局配置选项并确保它们有效。
	if err := setupGlobalConfig(); err != nil {
		return err
	}

//确保预期参数。
	if len(args) < 1 {
		return errors.New("required block hash parameter not specified")
	}
	if len(args) < 2 {
		return errors.New("required start offset parameter not " +
			"specified")
	}
	if len(args) < 3 {
		return errors.New("required region length parameter not " +
			"specified")
	}

//分析参数。
	blockHash, err := chainhash.NewHashFromStr(args[0])
	if err != nil {
		return err
	}
	startOffset, err := strconv.ParseUint(args[1], 10, 32)
	if err != nil {
		return err
	}
	regionLen, err := strconv.ParseUint(args[2], 10, 32)
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
		log.Infof("Fetching block region %s<%d:%d>", blockHash,
			startOffset, startOffset+regionLen-1)
		region := database.BlockRegion{
			Hash:   blockHash,
			Offset: uint32(startOffset),
			Len:    uint32(regionLen),
		}
		startTime := time.Now()
		regionBytes, err := tx.FetchBlockRegion(&region)
		if err != nil {
			return err
		}
		log.Infof("Loaded block region in %v", time.Since(startTime))
		log.Infof("Double Hash: %s", chainhash.DoubleHashH(regionBytes))
		log.Infof("Region Hex: %s", hex.EncodeToString(regionBytes))
		return nil
	})
}

//用法覆盖命令的用法显示。
func (cmd *blockRegionCmd) Usage() string {
	return "<block-hash> <start-offset> <length-of-region>"
}

