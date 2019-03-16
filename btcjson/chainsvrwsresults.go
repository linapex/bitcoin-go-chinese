
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:50</date>
//</624461719833415680>

//版权所有（c）2015-2017 BTCSuite开发者
//版权所有（c）2015-2017法令开发商
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package btcjson

//sessionresult对session命令中的数据进行建模。
type SessionResult struct {
	SessionID uint64 `json:"sessionid"`
}

//RescannedBlock包含单个的哈希和所有发现的事务
//重新扫描的块。
//
//注意：这是从中导入的BTCSuite扩展
//github.com/decred/dcrd/dcrjson。
type RescannedBlock struct {
	Hash         string   `json:"hash"`
	Transactions []string `json:"transactions"`
}

