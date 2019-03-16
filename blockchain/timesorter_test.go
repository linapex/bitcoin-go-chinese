
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:49</date>
//</624461715194515456>

//版权所有（c）2013-2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package blockchain

import (
	"reflect"
	"sort"
	"testing"
)

//TestTimesorter测试Timesorter实现。
func TestTimeSorter(t *testing.T) {
	tests := []struct {
		in   []int64
		want []int64
	}{
		{
			in: []int64{
1351228575, //10月26日星期五05:16:15 UTC 2012（布洛克205000）
1348310759, //2012年9月22日星期六10:45:59 UTC（20万块）
1305758502, //5月18日星期三22:41:42 UTC 2011（12.5万区）
1347777156, //太阳9月16日06:32:36 UTC 2012（布洛克199000）
1349492104, //2012年10月6日星期六02:55:04 UTC（布洛克202000）
			},
			want: []int64{
1305758502, //5月18日星期三22:41:42 UTC 2011（12.5万区）
1347777156, //太阳9月16日06:32:36 UTC 2012（布洛克199000）
1348310759, //2012年9月22日星期六10:45:59 UTC（20万块）
1349492104, //2012年10月6日星期六02:55:04 UTC（布洛克202000）
1351228575, //10月26日星期五05:16:15 UTC 2012（布洛克205000）
			},
		},
	}

	for i, test := range tests {
		result := make([]int64, len(test.in))
		copy(result, test.in)
		sort.Sort(timeSorter(result))
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("timeSorter #%d got %v want %v", i, result,
				test.want)
			continue
		}
	}
}

