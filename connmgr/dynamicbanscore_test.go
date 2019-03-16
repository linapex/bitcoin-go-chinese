
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:52</date>
//</624461727190224896>

//版权所有（c）2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package connmgr

import (
	"math"
	"testing"
	"time"
)

//testdynamicbanscoredecay测试在
//动态核心。
func TestDynamicBanScoreDecay(t *testing.T) {
	var bs DynamicBanScore
	base := time.Now()

	r := bs.increase(100, 50, base)
	if r != 150 {
		t.Errorf("Unexpected result %d after ban score increase.", r)
	}

	r = bs.int(base.Add(time.Minute))
	if r != 125 {
		t.Errorf("Halflife check failed - %d instead of 125", r)
	}

	r = bs.int(base.Add(7 * time.Minute))
	if r != 100 {
		t.Errorf("Decay after 7m - %d instead of 100", r)
	}
}

//testdynamicbanscoreLifetime测试dynamicbanscore是否正确生成零
//一旦达到最大年龄。
func TestDynamicBanScoreLifetime(t *testing.T) {
	var bs DynamicBanScore
	base := time.Now()

	r := bs.increase(0, math.MaxUint32, base)
	r = bs.int(base.Add(Lifetime * time.Second))
if r != 3 { //3，而不是4由于精度损失和截断3.999…
		t.Errorf("Pre max age check with MaxUint32 failed - %d", r)
	}
	r = bs.int(base.Add((Lifetime + 1) * time.Second))
	if r != 0 {
		t.Errorf("Zero after max age check failed - %d instead of 0", r)
	}
}

//testdynamicbanscore测试导出dynamicbanscore的函数。指数的
//衰减或其他基于时间的行为由其他函数进行测试。
func TestDynamicBanScoreReset(t *testing.T) {
	var bs DynamicBanScore
	if bs.Int() != 0 {
		t.Errorf("Initial state is not zero.")
	}
	bs.Increase(100, 0)
	r := bs.Int()
	if r != 100 {
		t.Errorf("Unexpected result %d after ban score increase.", r)
	}
	bs.Reset()
	if bs.Int() != 0 {
		t.Errorf("Failed to reset ban score.")
	}
}

