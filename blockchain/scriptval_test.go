
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:49</date>
//</624461714477289472>

//版权所有（c）2013-2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package blockchain

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/btcsuite/btcd/txscript"
)

//testcheckblockscripts确保验证
//已知良好的块不会返回错误。
func TestCheckBlockScripts(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	testBlockNum := 277647
	blockDataFile := fmt.Sprintf("%d.dat.bz2", testBlockNum)
	blocks, err := loadBlocks(blockDataFile)
	if err != nil {
		t.Errorf("Error loading file: %v\n", err)
		return
	}
	if len(blocks) > 1 {
		t.Errorf("The test block file must only have one block in it")
		return
	}
	if len(blocks) == 0 {
		t.Errorf("The test block file may not be empty")
		return
	}

	storeDataFile := fmt.Sprintf("%d.utxostore.bz2", testBlockNum)
	view, err := loadUtxoView(storeDataFile)
	if err != nil {
		t.Errorf("Error loading txstore: %v\n", err)
		return
	}

	scriptFlags := txscript.ScriptBip16
	err = checkBlockScripts(blocks[0], view, scriptFlags, nil, nil)
	if err != nil {
		t.Errorf("Transaction script validation failed: %v\n", err)
		return
	}
}

