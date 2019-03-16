
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:50</date>
//</624461720789716992>

//版权所有（c）2014 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package btcjson_test

import (
	"testing"

	"github.com/btcsuite/btcd/btcjson"
)

//TesterRorCodeStringer测试错误代码类型的字符串化输出。
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   btcjson.ErrorCode
		want string
	}{
		{btcjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{btcjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{btcjson.ErrInvalidType, "ErrInvalidType"},
		{btcjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{btcjson.ErrUnexportedField, "ErrUnexportedField"},
		{btcjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{btcjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{btcjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{btcjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{btcjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{btcjson.ErrNumParams, "ErrNumParams"},
		{btcjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

//检测没有添加字符串的其他错误代码。
	if len(tests)-1 != int(btcjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
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

//TesterRor测试错误类型的错误输出。
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   btcjson.Error
		want string
	}{
		{
			btcjson.Error{Description: "some error"},
			"some error",
		},
		{
			btcjson.Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

