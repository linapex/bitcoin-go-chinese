
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:58</date>
//</624461754474172416>

//版权所有（c）2013-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package wire

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/davecgh/go-spew/spew"
)

//testnotfound测试msgnotfound API。
func TestNotFound(t *testing.T) {
	pver := ProtocolVersion

//确保命令为预期值。
	wantCmd := "notfound"
	msg := NewMsgNotFound()
	if cmd := msg.Command(); cmd != wantCmd {
		t.Errorf("NewMsgNotFound: wrong command - got %v want %v",
			cmd, wantCmd)
	}

//确保最大有效负载是最新协议版本的预期值。
//num inventory vectors（varint）+允许的最大库存向量。
	wantPayload := uint32(1800009)
	maxPayload := msg.MaxPayloadLength(pver)
	if maxPayload != wantPayload {
		t.Errorf("MaxPayloadLength: wrong max payload length for "+
			"protocol version %d - got %v, want %v", pver,
			maxPayload, wantPayload)
	}

//确保正确添加库存向量。
	hash := chainhash.Hash{}
	iv := NewInvVect(InvTypeBlock, &hash)
	err := msg.AddInvVect(iv)
	if err != nil {
		t.Errorf("AddInvVect: %v", err)
	}
	if msg.InvList[0] != iv {
		t.Errorf("AddInvVect: wrong invvect added - got %v, want %v",
			spew.Sprint(msg.InvList[0]), spew.Sprint(iv))
	}

//确保在每个
//消息返回错误。
	for i := 0; i < MaxInvPerMsg; i++ {
		err = msg.AddInvVect(iv)
	}
	if err == nil {
		t.Errorf("AddInvVect: expected error on too many inventory " +
			"vectors not received")
	}
}

//testnotfoundwire测试msgnotfound wire编码和解码
//库存向量和协议版本的数量。
func TestNotFoundWire(t *testing.T) {
//块203707哈希。
	hashStr := "3264bc2ac36a60840790ba1d475d01367e7c723da941069e9dc"
	blockHash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		t.Errorf("NewHashFromStr: %v", err)
	}

//块203707哈希的事务1。
	hashStr = "d28a3dc7392bf00a9855ee93dd9a81eff82a2c4fe57fbd42cfe71b487accfaf0"
	txHash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		t.Errorf("NewHashFromStr: %v", err)
	}

	iv := NewInvVect(InvTypeBlock, blockHash)
	iv2 := NewInvVect(InvTypeTx, txHash)

//空的未找到消息。
	NoInv := NewMsgNotFound()
	NoInvEncoded := []byte{
0x00, //库存向量数量的变量
	}

//未找到包含多个库存向量的消息。
	MultiInv := NewMsgNotFound()
	MultiInv.AddInvVect(iv)
	MultiInv.AddInvVect(iv2)
	MultiInvEncoded := []byte{
0x02,                   //
0x02, 0x00, 0x00, 0x00, //输入块
		0xdc, 0xe9, 0x69, 0x10, 0x94, 0xda, 0x23, 0xc7,
		0xe7, 0x67, 0x13, 0xd0, 0x75, 0xd4, 0xa1, 0x0b,
		0x79, 0x40, 0x08, 0xa6, 0x36, 0xac, 0xc2, 0x4b,
0x26, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //块203707哈希
0x01, 0x00, 0x00, 0x00, //输入字体
		0xf0, 0xfa, 0xcc, 0x7a, 0x48, 0x1b, 0xe7, 0xcf,
		0x42, 0xbd, 0x7f, 0xe5, 0x4f, 0x2c, 0x2a, 0xf8,
		0xef, 0x81, 0x9a, 0xdd, 0x93, 0xee, 0x55, 0x98,
0x0a, 0xf0, 0x2b, 0x39, 0xc7, 0x3d, 0x8a, 0xd2, //203707号区块哈希的Tx 1
	}

	tests := []struct {
in   *MsgNotFound    //要编码的邮件
out  *MsgNotFound    //预期的解码消息
buf  []byte          //有线编码
pver uint32          //有线编码协议版本
enc  MessageEncoding //消息编码格式
	}{
//最新协议版本，无inv向量。
		{
			NoInv,
			NoInv,
			NoInvEncoded,
			ProtocolVersion,
			BaseEncoding,
		},

//具有多个inv向量的最新协议版本。
		{
			MultiInv,
			MultiInv,
			MultiInvEncoded,
			ProtocolVersion,
			BaseEncoding,
		},

//协议版本BIP0035版本无INV矢量。
		{
			NoInv,
			NoInv,
			NoInvEncoded,
			BIP0035Version,
			BaseEncoding,
		},

//协议版本bip0035，带有多个inv向量。
		{
			MultiInv,
			MultiInv,
			MultiInvEncoded,
			BIP0035Version,
			BaseEncoding,
		},

//协议版本bip0031版本无inv向量。
		{
			NoInv,
			NoInv,
			NoInvEncoded,
			BIP0031Version,
			BaseEncoding,
		},

//协议版本bip0031，带有多个inv向量。
		{
			MultiInv,
			MultiInv,
			MultiInvEncoded,
			BIP0031Version,
			BaseEncoding,
		},

//协议版本netaddresstimeversion没有inv向量。
		{
			NoInv,
			NoInv,
			NoInvEncoded,
			NetAddressTimeVersion,
			BaseEncoding,
		},

//协议版本netaddresstimeversion，带有多个inv向量。
		{
			MultiInv,
			MultiInv,
			MultiInvEncoded,
			NetAddressTimeVersion,
			BaseEncoding,
		},

//协议版本multipleaddressversion无inv向量。
		{
			NoInv,
			NoInv,
			NoInvEncoded,
			MultipleAddressVersion,
			BaseEncoding,
		},

//具有多个inv向量的协议版本multipleaddressversion。
		{
			MultiInv,
			MultiInv,
			MultiInvEncoded,
			MultipleAddressVersion,
			BaseEncoding,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
//将邮件编码为有线格式。
		var buf bytes.Buffer
		err := test.in.BtcEncode(&buf, test.pver, test.enc)
		if err != nil {
			t.Errorf("BtcEncode #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("BtcEncode #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

//从有线格式解码消息。
		var msg MsgNotFound
		rbuf := bytes.NewReader(test.buf)
		err = msg.BtcDecode(rbuf, test.pver, test.enc)
		if err != nil {
			t.Errorf("BtcDecode #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&msg, test.out) {
			t.Errorf("BtcDecode #%d\n got: %s want: %s", i,
				spew.Sdump(msg), spew.Sdump(test.out))
			continue
		}
	}
}

//TestNotFoundWireErrors对线编码和解码执行负测试
//以确认错误路径是否正常工作。
func TestNotFoundWireErrors(t *testing.T) {
	pver := ProtocolVersion
	wireErr := &MessageError{}

//块203707哈希。
	hashStr := "3264bc2ac36a60840790ba1d475d01367e7c723da941069e9dc"
	blockHash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		t.Errorf("NewHashFromStr: %v", err)
	}

	iv := NewInvVect(InvTypeBlock, blockHash)

//用于引发错误的基本消息。
	baseNotFound := NewMsgNotFound()
	baseNotFound.AddInvVect(iv)
	baseNotFoundEncoded := []byte{
0x02,                   //矢量数变量
0x02, 0x00, 0x00, 0x00, //输入块
		0xdc, 0xe9, 0x69, 0x10, 0x94, 0xda, 0x23, 0xc7,
		0xe7, 0x67, 0x13, 0xd0, 0x75, 0xd4, 0xa1, 0x0b,
		0x79, 0x40, 0x08, 0xa6, 0x36, 0xac, 0xc2, 0x4b,
0x26, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //块203707哈希
	}

//通过超过允许的最大inv来强制出错的消息
//向量。
	maxNotFound := NewMsgNotFound()
	for i := 0; i < MaxInvPerMsg; i++ {
		maxNotFound.AddInvVect(iv)
	}
	maxNotFound.InvList = append(maxNotFound.InvList, iv)
	maxNotFoundEncoded := []byte{
0xfd, 0x51, 0xc3, //INV矢量数变量（50001）
	}

	tests := []struct {
in       *MsgNotFound    //编码值
buf      []byte          //有线编码
pver     uint32          //有线编码协议版本
enc      MessageEncoding //消息编码格式
max      int             //引发错误的固定缓冲区的最大大小
writeErr error           //预期的写入错误
readErr  error           //预期的读取错误
	}{
//库存向量计数中的强制错误
		{baseNotFound, baseNotFoundEncoded, pver, BaseEncoding, 0, io.ErrShortWrite, io.EOF},
//库存列表中的强制错误。
		{baseNotFound, baseNotFoundEncoded, pver, BaseEncoding, 1, io.ErrShortWrite, io.EOF},
//强制错误大于最大库存向量。
		{maxNotFound, maxNotFoundEncoded, pver, BaseEncoding, 3, wireErr, wireErr},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
//编码为有线格式。
		w := newFixedWriter(test.max)
		err := test.in.BtcEncode(w, test.pver, test.enc)
		if reflect.TypeOf(err) != reflect.TypeOf(test.writeErr) {
			t.Errorf("BtcEncode #%d wrong error got: %v, want: %v",
				i, err, test.writeErr)
			continue
		}

//对于不属于messageerror类型的错误，请检查它们
//平等。
		if _, ok := err.(*MessageError); !ok {
			if err != test.writeErr {
				t.Errorf("BtcEncode #%d wrong error got: %v, "+
					"want: %v", i, err, test.writeErr)
				continue
			}
		}

//从有线格式解码。
		var msg MsgNotFound
		r := newFixedReader(test.max, test.buf)
		err = msg.BtcDecode(r, test.pver, test.enc)
		if reflect.TypeOf(err) != reflect.TypeOf(test.readErr) {
			t.Errorf("BtcDecode #%d wrong error got: %v, want: %v",
				i, err, test.readErr)
			continue
		}

//对于不属于messageerror类型的错误，请检查它们
//平等。
		if _, ok := err.(*MessageError); !ok {
			if err != test.readErr {
				t.Errorf("BtcDecode #%d wrong error got: %v, "+
					"want: %v", i, err, test.readErr)
				continue
			}
		}
	}
}

