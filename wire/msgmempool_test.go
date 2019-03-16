
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:58</date>
//</624461754142822400>

//版权所有（c）2013-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package wire

import (
	"bytes"
	"testing"
)

func TestMemPool(t *testing.T) {
	pver := ProtocolVersion
	enc := BaseEncoding

//确保命令为预期值。
	wantCmd := "mempool"
	msg := NewMsgMemPool()
	if cmd := msg.Command(); cmd != wantCmd {
		t.Errorf("NewMsgMemPool: wrong command - got %v want %v",
			cmd, wantCmd)
	}

//确保最大有效负载为预期值。
	wantPayload := uint32(0)
	maxPayload := msg.MaxPayloadLength(pver)
	if maxPayload != wantPayload {
		t.Errorf("MaxPayloadLength: wrong max payload length for "+
			"protocol version %d - got %v, want %v", pver,
			maxPayload, wantPayload)
	}

//使用最新的协议版本进行测试编码。
	var buf bytes.Buffer
	err := msg.BtcEncode(&buf, pver, enc)
	if err != nil {
		t.Errorf("encode of MsgMemPool failed %v err <%v>", msg, err)
	}

//旧的协议版本应该无法编码，因为消息没有
//还存在。
	oldPver := BIP0035Version - 1
	err = msg.BtcEncode(&buf, oldPver, enc)
	if err == nil {
		s := "encode of MsgMemPool passed for old protocol version %v err <%v>"
		t.Errorf(s, msg, err)
	}

//使用最新的协议版本测试解码。
	readmsg := NewMsgMemPool()
	err = readmsg.BtcDecode(&buf, pver, enc)
	if err != nil {
		t.Errorf("decode of MsgMemPool failed [%v] err <%v>", buf, err)
	}

//旧的协议版本应该无法解码，因为消息没有
//还存在。
	err = readmsg.BtcDecode(&buf, oldPver, enc)
	if err == nil {
		s := "decode of MsgMemPool passed for old protocol version %v err <%v>"
		t.Errorf(s, msg, err)
	}
}

