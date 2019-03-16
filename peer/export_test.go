
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:55</date>
//</624461739408232448>

//版权所有（c）2015 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

/*
This test file is part of the peer package rather than than the peer_test
包装，使其能够连接内部构件，以正确测试
不可能或不能通过公共接口可靠地进行测试。
仅在运行测试时导出函数。
**/


package peer

//tstallowselfconns允许测试包通过
//禁用检测逻辑。
func TstAllowSelfConns() {
	allowSelfConns = true
}

