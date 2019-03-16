
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:50</date>
//</624461718021476352>

//版权所有（c）2015 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package btcec

//
//不要编辑

var secp256k1BytePoints = "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

