
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:53</date>
//</624461733498458112>

//版权所有（c）2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package rpctest

import (
	"reflect"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
)

//joinType是表示特定类型“节点联接”的枚举。节点
//join是一个同步工具，用于等待节点的子集
//关于属性的一致状态。
type JoinType uint8

const (
//Blocks is a JoinType which waits until all nodes share the same
//块高度。
	Blocks JoinType = iota

//mempools是一个joinType，它阻塞所有节点，直到所有节点都相同为止。
//内存池。
	Mempools
)

//JoinNodes is a synchronization tool used to block until all passed nodes are
//与属性完全同步。此函数将阻止
//一段时间，当所有节点按照
//传递了JoinType。此功能用于确保所有激活的测试
//在进行断言或
//在RPC测试中进行检查。
func JoinNodes(nodes []*Harness, joinType JoinType) error {
	switch joinType {
	case Blocks:
		return syncBlocks(nodes)
	case Mempools:
		return syncMempools(nodes)
	}
	return nil
}

//SyncMemPools将阻止，直到所有节点都具有相同的MemPools。
func syncMempools(nodes []*Harness) error {
	poolsMatch := false

retry:
	for !poolsMatch {
		firstPool, err := nodes[0].Node.GetRawMempool()
		if err != nil {
			return err
		}

//如果所有节点的mempool与
//第一个节点，然后我们就完成了。否则，返回顶部
//循环并在短的等待期后重试。
		for _, node := range nodes[1:] {
			nodePool, err := node.Node.GetRawMempool()
			if err != nil {
				return err
			}

			if !reflect.DeepEqual(firstPool, nodePool) {
				time.Sleep(time.Millisecond * 100)
				continue retry
			}
		}

		poolsMatch = true
	}

	return nil
}

//同步块将一直阻止，直到所有节点报告相同的最佳链。
func syncBlocks(nodes []*Harness) error {
	blocksMatch := false

retry:
	for !blocksMatch {
		var prevHash *chainhash.Hash
		var prevHeight int32
		for _, node := range nodes {
			blockHash, blockHeight, err := node.Node.GetBestBlock()
			if err != nil {
				return err
			}
			if prevHash != nil && (*blockHash != *prevHash ||
				blockHeight != prevHeight) {

				time.Sleep(time.Millisecond * 100)
				continue retry
			}
			prevHash, prevHeight = blockHash, blockHeight
		}

		blocksMatch = true
	}

	return nil
}

//ConnectNode establishes a new peer-to-peer connection between the "from"
//线束和“到”线束。所建立的连接被标记为持久连接，
//因此，在断开的情况下，“从”将尝试重新建立
//连接至“至”线束。
func ConnectNode(from *Harness, to *Harness) error {
	peerInfo, err := from.Node.GetPeerInfo()
	if err != nil {
		return err
	}
	numPeers := len(peerInfo)

	targetAddr := to.node.config.listen
	if err := from.Node.AddNode(targetAddr, rpcclient.ANAdd); err != nil {
		return err
	}

//阻止，直到建立新连接。
	peerInfo, err = from.Node.GetPeerInfo()
	if err != nil {
		return err
	}
	for len(peerInfo) <= numPeers {
		peerInfo, err = from.Node.GetPeerInfo()
		if err != nil {
			return err
		}
	}

	return nil
}

//拆下所有激活的测试线束。
func TearDownAll() error {
	harnessStateMtx.Lock()
	defer harnessStateMtx.Unlock()

	for _, harness := range testInstances {
		if err := harness.tearDown(); err != nil {
			return err
		}
	}

	return nil
}

//ActiveHarness返回当前所有活动测试线束的一部分。一
//如果线束已创建但尚未撕裂，则将其视为“激活”测试线束
//下来。
func ActiveHarnesses() []*Harness {
	harnessStateMtx.RLock()
	defer harnessStateMtx.RUnlock()

	activeNodes := make([]*Harness, 0, len(testInstances))
	for _, harness := range testInstances {
		activeNodes = append(activeNodes, harness)
	}

	return activeNodes
}

