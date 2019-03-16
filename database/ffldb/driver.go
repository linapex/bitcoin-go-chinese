
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:52</date>
//</624461729325125632>

//版权所有（c）2015-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package ffldb

import (
	"fmt"

	"github.com/btcsuite/btcd/database"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btclog"
)

var log = btclog.Disabled

const (
	dbType = "ffldb"
)

//parseargs解析数据库open/create方法中的参数。
func parseArgs(funcName string, args ...interface{}) (string, wire.BitcoinNet, error) {
	if len(args) != 2 {
		return "", 0, fmt.Errorf("invalid arguments to %s.%s -- "+
			"expected database path and block network", dbType,
			funcName)
	}

	dbPath, ok := args[0].(string)
	if !ok {
		return "", 0, fmt.Errorf("first argument to %s.%s is invalid -- "+
			"expected database path string", dbType, funcName)
	}

	network, ok := args[1].(wire.BitcoinNet)
	if !ok {
		return "", 0, fmt.Errorf("second argument to %s.%s is invalid -- "+
			"expected block network", dbType, funcName)
	}

	return dbPath, network, nil
}

//OpenDBDriver是在驱动程序注册期间提供的回调，它打开
//要使用的现有数据库。
func openDBDriver(args ...interface{}) (database.DB, error) {
	dbPath, network, err := parseArgs("Open", args...)
	if err != nil {
		return nil, err
	}

	return openDB(dbPath, network, false)
}

//createdbDriver是在驱动程序注册期间提供的回调，
//创建、初始化并打开数据库以供使用。
func createDBDriver(args ...interface{}) (database.DB, error) {
	dbPath, network, err := parseArgs("Create", args...)
	if err != nil {
		return nil, err
	}

	return openDB(dbPath, network, true)
}

//useLogger是在驱动程序注册期间提供的回调，用于设置
//当前记录器到提供的记录器。
func useLogger(logger btclog.Logger) {
	log = logger
}

func init() {
//登记驾驶员。
	driver := database.Driver{
		DbType:    dbType,
		Create:    createDBDriver,
		Open:      openDBDriver,
		UseLogger: useLogger,
	}
	if err := database.RegisterDriver(driver); err != nil {
		panic(fmt.Sprintf("Failed to regiser database driver '%s': %v",
			dbType, err))
	}
}

