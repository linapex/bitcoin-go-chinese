
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:53</date>
//</624461730608582656>

//版权所有（c）2015-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

/*
包treap实现用于保持有序的treap数据结构
使用二进制搜索树和堆语义组合的键/值对。
它是一种自组织随机数据结构，不需要
维持平衡的复杂操作。搜索、插入和删除
操作都是O（log n）。同时提供可变和不可变的变体。

可变变量通常更快，因为它能够简单地更新
t进行修改时重新启动。然而，一个可变的叛国者是不安全的
在不小心使用锁的情况下，调用方必须谨慎地进行并发访问。
在迭代时获取，因为它可以从迭代器下更改。

不变的变体通过为所有人创建一个新版本的treap来工作。
通过用具有更新值的新节点替换修改后的节点进行突变
与以前的版本共享所有未修改的节点时。这是非常
在并发应用程序中很有用，因为调用者只需要原子地
在执行任何
突变。所有读卡器都可以简单地将其现有指针用作快照
因为它所指的叛国罪是不变的。这有效地提供了O（1）
snapshot capability with efficient memory usage characteristics since the old
只有在不再有任何对节点的引用之前，节点才会保持分配状态。
**/

package treap

