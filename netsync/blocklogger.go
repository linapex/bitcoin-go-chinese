
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:54</date>
//</624461737554350080>

//版权所有（c）2015-2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package netsync

import (
	"sync"
	"time"

	"github.com/btcsuite/btclog"
	"github.com/btcsuite/btcutil"
)

//BlockProgressLogger按顺序为其他服务提供定期日志记录
//向用户显示某些“操作”的进度，这些操作涉及某些或所有当前操作
//阻碍。例如：同步到最佳链，索引所有块等。
type blockProgressLogger struct {
	receivedLogBlocks int64
	receivedLogTx     int64
	lastBlockLogTime  time.Time

	subsystemLogger btclog.Logger
	progressAction  string
	sync.Mutex
}

//NewblockProgressLogger返回一个新的块进度记录器。
//进度消息模板如下：
//ProgressAction NumProcessed Blocks Block在最后一个时间段
//（numtxs，height lastblockheight，lastblocktimestamp）
func newBlockProgressLogger(progressMessage string, logger btclog.Logger) *blockProgressLogger {
	return &blockProgressLogger{
		lastBlockLogTime: time.Now(),
		progressAction:   progressMessage,
		subsystemLogger:  logger,
	}
}

//logblockheight将新的块高度记录为要显示的信息消息
//用户的进度。为了防止垃圾邮件，它将日志记录限制为一个
//每10秒发送一条消息，包括持续时间和总数。
func (b *blockProgressLogger) LogBlockHeight(block *btcutil.Block) {
	b.Lock()
	defer b.Unlock()

	b.receivedLogBlocks++
	b.receivedLogTx += int64(len(block.MsgBlock().Transactions))

	now := time.Now()
	duration := now.Sub(b.lastBlockLogTime)
	if duration < time.Second*10 {
		return
	}

//将持续时间截断为10毫秒。
	durationMillis := int64(duration / time.Millisecond)
	tDuration := 10 * time.Millisecond * time.Duration(durationMillis/10)

//有关新块高度的日志信息。
	blockStr := "blocks"
	if b.receivedLogBlocks == 1 {
		blockStr = "block"
	}
	txStr := "transactions"
	if b.receivedLogTx == 1 {
		txStr = "transaction"
	}
	b.subsystemLogger.Infof("%s %d %s in the last %s (%d %s, height %d, %s)",
		b.progressAction, b.receivedLogBlocks, blockStr, tDuration, b.receivedLogTx,
		txStr, block.Height(), block.MsgBlock().Header.Timestamp)

	b.receivedLogBlocks = 0
	b.receivedLogTx = 0
	b.lastBlockLogTime = now
}

func (b *blockProgressLogger) SetLastLogTime(time time.Time) {
	b.lastBlockLogTime = time
}

