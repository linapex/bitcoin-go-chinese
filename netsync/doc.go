
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:55</date>
//</624461737957003264>

//版权所有（c）2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

/*
包netsync实现并发安全块同步协议。这个
SyncManager与连接的对等机通信以执行初始块
下载、保持链和未确认的事务池同步，并宣布
连接到链条上的新块。当前同步管理器选择一个
同步它下载所有块的对等机，直到它与
同步对等机知道的最长链。
**/

package netsync

