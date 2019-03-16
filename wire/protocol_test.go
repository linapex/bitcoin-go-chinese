
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:59</date>
//</624461756055425024>

//版权所有（c）2013-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package wire

import "testing"

//TestServiceFlagstringer测试服务标志类型的字符串化输出。
func TestServiceFlagStringer(t *testing.T) {
	tests := []struct {
		in   ServiceFlag
		want string
	}{
		{0, "0x0"},
		{SFNodeNetwork, "SFNodeNetwork"},
		{SFNodeGetUTXO, "SFNodeGetUTXO"},
		{SFNodeBloom, "SFNodeBloom"},
		{SFNodeWitness, "SFNodeWitness"},
		{SFNodeXthin, "SFNodeXthin"},
		{SFNodeBit5, "SFNodeBit5"},
		{SFNodeCF, "SFNodeCF"},
		{SFNode2X, "SFNode2X"},
		{0xffffffff, "SFNodeNetwork|SFNodeGetUTXO|SFNodeBloom|SFNodeWitness|SFNodeXthin|SFNodeBit5|SFNodeCF|SFNode2X|0xffffff00"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

//TestBitcoinneStringer测试比特币网络类型的串化输出。
func TestBitcoinNetStringer(t *testing.T) {
	tests := []struct {
		in   BitcoinNet
		want string
	}{
		{MainNet, "MainNet"},
		{TestNet, "TestNet"},
		{TestNet3, "TestNet3"},
		{SimNet, "SimNet"},
		{0xffffffff, "Unknown BitcoinNet (4294967295)"},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

