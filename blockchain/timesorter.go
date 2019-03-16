
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:49</date>
//</624461715123212288>

//版权所有（c）2013-2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package blockchain

//TimeSorter实现Sort.Interface以允许时间戳切片
//分类。
type timeSorter []int64

//len返回切片中的时间戳数。它是
//
func (s timeSorter) Len() int {
	return len(s)
}

//
//Sort.Interface实现。
func (s timeSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//less返回索引为i的timstamp是否应在
//带有索引j的时间戳。它是sort.interface实现的一部分。
func (s timeSorter) Less(i, j int) bool {
	return s[i] < s[j]
}

