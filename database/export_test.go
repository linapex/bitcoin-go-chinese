
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:52</date>
//</624461728729534464>

//版权所有（c）2015-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

/*
此测试文件是数据库包的一部分，而不是
数据库测试包，以便它能够桥接对内部的访问，以便正确测试
不可能或不能通过公众可靠测试的案例
接口。函数、常量和变量仅在
正在运行测试。
**/


package database

//tstnumerrorcodes使内部numerorcodes参数可用于
//测试包。
const TstNumErrorCodes = numErrorCodes

