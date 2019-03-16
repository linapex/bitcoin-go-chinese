
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:50</date>
//</624461716922568704>

//版权所有（c）2013-2014 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

/*
包btcec实现了对比特币所需椭圆曲线的支持。

比特币使用椭圆曲线加密，使用Koblitz曲线
（特别是secp256k1）用于加密函数。见
http://www.secg.org/collateral/sec2_final.pdf了解
标准。

此包提供实现
加密/椭圆曲线接口，以允许使用这些曲线
使用Go提供的标准Crypto/ECDSA包。帮手
提供了分析签名和公钥的功能
标准格式。它设计用于BTCD，但应该


**/

package btcec

