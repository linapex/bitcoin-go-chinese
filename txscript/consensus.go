
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:56</date>
//</624461743866777600>

//版权所有（c）2015-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package txscript

const (
//LockTimeThreshold是锁定时间低于的数字。
//解释为块编号。因为平均一个街区
//每10分钟生成一次，这允许大约9512个块
//年。
LockTimeThreshold = 5e8 //1985年11月5日星期二00:53:20 UTC
)

