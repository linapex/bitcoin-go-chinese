
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:59</date>
//</624461755325616128>

//版权所有（c）2013-2016 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package wire

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/davecgh/go-spew/spew"
)

//testx测试MSGTX API。
func TestTx(t *testing.T) {
	pver := ProtocolVersion

//阻止100000哈希。
	hashStr := "3ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506"
	hash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		t.Errorf("NewHashFromStr: %v", err)
	}

//确保命令为预期值。
	wantCmd := "tx"
	msg := NewMsgTx(1)
	if cmd := msg.Command(); cmd != wantCmd {
		t.Errorf("NewMsgAddr: wrong command - got %v want %v",
			cmd, wantCmd)
	}

//确保最大有效负载是最新协议版本的预期值。
	wantPayload := uint32(1000 * 4000)
	maxPayload := msg.MaxPayloadLength(pver)
	if maxPayload != wantPayload {
		t.Errorf("MaxPayloadLength: wrong max payload length for "+
			"protocol version %d - got %v, want %v", pver,
			maxPayload, wantPayload)
	}

//确保我们得到相同的事务输出点数据。
//注意：这是一个块散列和组成的索引，但我们只是
//测试包功能。
	prevOutIndex := uint32(1)
	prevOut := NewOutPoint(hash, prevOutIndex)
	if !prevOut.Hash.IsEqual(hash) {
		t.Errorf("NewOutPoint: wrong hash - got %v, want %v",
			spew.Sprint(&prevOut.Hash), spew.Sprint(hash))
	}
	if prevOut.Index != prevOutIndex {
		t.Errorf("NewOutPoint: wrong index - got %v, want %v",
			prevOut.Index, prevOutIndex)
	}
	prevOutStr := fmt.Sprintf("%s:%d", hash.String(), prevOutIndex)
	if s := prevOut.String(); s != prevOutStr {
		t.Errorf("OutPoint.String: unexpected result - got %v, "+
			"want %v", s, prevOutStr)
	}

//确保我们得到相同的事务输入。
	sigScript := []byte{0x04, 0x31, 0xdc, 0x00, 0x1b, 0x01, 0x62}
	witnessData := [][]byte{
		{0x04, 0x31},
		{0x01, 0x43},
	}
	txIn := NewTxIn(prevOut, sigScript, witnessData)
	if !reflect.DeepEqual(&txIn.PreviousOutPoint, prevOut) {
		t.Errorf("NewTxIn: wrong prev outpoint - got %v, want %v",
			spew.Sprint(&txIn.PreviousOutPoint),
			spew.Sprint(prevOut))
	}
	if !bytes.Equal(txIn.SignatureScript, sigScript) {
		t.Errorf("NewTxIn: wrong signature script - got %v, want %v",
			spew.Sdump(txIn.SignatureScript),
			spew.Sdump(sigScript))
	}
	if !reflect.DeepEqual(txIn.Witness, TxWitness(witnessData)) {
		t.Errorf("NewTxIn: wrong witness data - got %v, want %v",
			spew.Sdump(txIn.Witness),
			spew.Sdump(witnessData))
	}

//确保我们得到相同的事务输出。
	txValue := int64(5000000000)
	pkScript := []byte{
0x41, //OPDA DATA65
		0x04, 0xd6, 0x4b, 0xdf, 0xd0, 0x9e, 0xb1, 0xc5,
		0xfe, 0x29, 0x5a, 0xbd, 0xeb, 0x1d, 0xca, 0x42,
		0x81, 0xbe, 0x98, 0x8e, 0x2d, 0xa0, 0xb6, 0xc1,
		0xc6, 0xa5, 0x9d, 0xc2, 0x26, 0xc2, 0x86, 0x24,
		0xe1, 0x81, 0x75, 0xe8, 0x51, 0xc9, 0x6b, 0x97,
		0x3d, 0x81, 0xb0, 0x1c, 0xc3, 0x1f, 0x04, 0x78,
		0x34, 0xbc, 0x06, 0xd6, 0xd6, 0xed, 0xf6, 0x20,
		0xd1, 0x84, 0x24, 0x1a, 0x6a, 0xed, 0x8b, 0x63,
0xa6, //65字节签名
0xac, //奥普克西格
	}
	txOut := NewTxOut(txValue, pkScript)
	if txOut.Value != txValue {
		t.Errorf("NewTxOut: wrong pk script - got %v, want %v",
			txOut.Value, txValue)

	}
	if !bytes.Equal(txOut.PkScript, pkScript) {
		t.Errorf("NewTxOut: wrong pk script - got %v, want %v",
			spew.Sdump(txOut.PkScript),
			spew.Sdump(pkScript))
	}

//确保正确添加事务输入。
	msg.AddTxIn(txIn)
	if !reflect.DeepEqual(msg.TxIn[0], txIn) {
		t.Errorf("AddTxIn: wrong transaction input added - got %v, want %v",
			spew.Sprint(msg.TxIn[0]), spew.Sprint(txIn))
	}

//确保正确添加事务输出。
	msg.AddTxOut(txOut)
	if !reflect.DeepEqual(msg.TxOut[0], txOut) {
		t.Errorf("AddTxIn: wrong transaction output added - got %v, want %v",
			spew.Sprint(msg.TxOut[0]), spew.Sprint(txOut))
	}

//确保副本生成了相同的事务消息。
	newMsg := msg.Copy()
	if !reflect.DeepEqual(newMsg, msg) {
		t.Errorf("Copy: mismatched tx messages - got %v, want %v",
			spew.Sdump(newMsg), spew.Sdump(msg))
	}
}

//testxthash测试准确生成事务散列的能力。
func TestTxHash(t *testing.T) {
//来自块113875的第一个事务的哈希。
	hashStr := "f051e59b5e2503ac626d03aaeac8ab7be2d72ba4b7e97119c5852d70d52dcb86"
	wantHash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		t.Errorf("NewHashFromStr: %v", err)
		return
	}

//来自区块113875的第一笔交易。
	msgTx := NewMsgTx(1)
	txIn := TxIn{
		PreviousOutPoint: OutPoint{
			Hash:  chainhash.Hash{},
			Index: 0xffffffff,
		},
		SignatureScript: []byte{0x04, 0x31, 0xdc, 0x00, 0x1b, 0x01, 0x62},
		Sequence:        0xffffffff,
	}
	txOut := TxOut{
		Value: 5000000000,
		PkScript: []byte{
0x41, //OPDA DATA65
			0x04, 0xd6, 0x4b, 0xdf, 0xd0, 0x9e, 0xb1, 0xc5,
			0xfe, 0x29, 0x5a, 0xbd, 0xeb, 0x1d, 0xca, 0x42,
			0x81, 0xbe, 0x98, 0x8e, 0x2d, 0xa0, 0xb6, 0xc1,
			0xc6, 0xa5, 0x9d, 0xc2, 0x26, 0xc2, 0x86, 0x24,
			0xe1, 0x81, 0x75, 0xe8, 0x51, 0xc9, 0x6b, 0x97,
			0x3d, 0x81, 0xb0, 0x1c, 0xc3, 0x1f, 0x04, 0x78,
			0x34, 0xbc, 0x06, 0xd6, 0xd6, 0xed, 0xf6, 0x20,
			0xd1, 0x84, 0x24, 0x1a, 0x6a, 0xed, 0x8b, 0x63,
0xa6, //65字节签名
0xac, //奥普克西格
		},
	}
	msgTx.AddTxIn(&txIn)
	msgTx.AddTxOut(&txOut)
	msgTx.LockTime = 0

//确保所生成的哈希是预期的。
	txHash := msgTx.TxHash()
	if !txHash.IsEqual(wantHash) {
		t.Errorf("TxHash: wrong hash - got %v, want %v",
			spew.Sprint(txHash), spew.Sprint(wantHash))
	}
}

//testxtsha测试生成事务的wtxid和txid的能力
//有证人的准确输入。
func TestWTxSha(t *testing.T) {
	hashStrTxid := "0f167d1385a84d1518cfee208b653fc9163b605ccf1b75347e2850b3e2eb19f3"
	wantHashTxid, err := chainhash.NewHashFromStr(hashStrTxid)
	if err != nil {
		t.Errorf("NewShaHashFromStr: %v", err)
		return
	}
	hashStrWTxid := "0858eab78e77b6b033da30f46699996396cf48fcf625a783c85a51403e175e74"
	wantHashWTxid, err := chainhash.NewHashFromStr(hashStrWTxid)
	if err != nil {
		t.Errorf("NewShaHashFromStr: %v", err)
		return
	}

//来自SEGNET以前版本的23157块。
	msgTx := NewMsgTx(1)
	txIn := TxIn{
		PreviousOutPoint: OutPoint{
			Hash: chainhash.Hash{
				0xa5, 0x33, 0x52, 0xd5, 0x13, 0x57, 0x66, 0xf0,
				0x30, 0x76, 0x59, 0x74, 0x18, 0x26, 0x3d, 0xa2,
				0xd9, 0xc9, 0x58, 0x31, 0x59, 0x68, 0xfe, 0xa8,
				0x23, 0x52, 0x94, 0x67, 0x48, 0x1f, 0xf9, 0xcd,
			},
			Index: 19,
		},
		Witness: [][]byte{
{ //70字节签名
				0x30, 0x43, 0x02, 0x1f, 0x4d, 0x23, 0x81, 0xdc,
				0x97, 0xf1, 0x82, 0xab, 0xd8, 0x18, 0x5f, 0x51,
				0x75, 0x30, 0x18, 0x52, 0x32, 0x12, 0xf5, 0xdd,
				0xc0, 0x7c, 0xc4, 0xe6, 0x3a, 0x8d, 0xc0, 0x36,
				0x58, 0xda, 0x19, 0x02, 0x20, 0x60, 0x8b, 0x5c,
				0x4d, 0x92, 0xb8, 0x6b, 0x6d, 0xe7, 0xd7, 0x8e,
				0xf2, 0x3a, 0x2f, 0xa7, 0x35, 0xbc, 0xb5, 0x9b,
				0x91, 0x4a, 0x48, 0xb0, 0xe1, 0x87, 0xc5, 0xe7,
				0x56, 0x9a, 0x18, 0x19, 0x70, 0x01,
			},
{ //33字节序列化发布密钥
				0x03, 0x07, 0xea, 0xd0, 0x84, 0x80, 0x7e, 0xb7,
				0x63, 0x46, 0xdf, 0x69, 0x77, 0x00, 0x0c, 0x89,
				0x39, 0x2f, 0x45, 0xc7, 0x64, 0x25, 0xb2, 0x61,
				0x81, 0xf5, 0x21, 0xd7, 0xf3, 0x70, 0x06, 0x6a,
				0x8f,
			},
		},
		Sequence: 0xffffffff,
	}
	txOut := TxOut{
		Value: 395019,
		PkScript: []byte{
0x00, //版本0见证程序
0x14, //OPDA DATAY20
			0x9d, 0xda, 0xc6, 0xf3, 0x9d, 0x51, 0xe0, 0x39,
			0x8e, 0x53, 0x2a, 0x22, 0xc4, 0x1b, 0xa1, 0x89,
0x40, 0x6a, 0x85, 0x23, //20字节pub密钥哈希
		},
	}
	msgTx.AddTxIn(&txIn)
	msgTx.AddTxOut(&txOut)
	msgTx.LockTime = 0

//确保正确的txid和wtxid按预期生产。
	txid := msgTx.TxHash()
	if !txid.IsEqual(wantHashTxid) {
		t.Errorf("TxSha: wrong hash - got %v, want %v",
			spew.Sprint(txid), spew.Sprint(wantHashTxid))
	}
	wtxid := msgTx.WitnessHash()
	if !wtxid.IsEqual(wantHashWTxid) {
		t.Errorf("WTxSha: wrong hash - got %v, want %v",
			spew.Sprint(wtxid), spew.Sprint(wantHashWTxid))
	}
}

//TESTTXWIRE测试MSGTX线对各种数字进行编码和解码
//事务输入和输出以及协议版本。
func TestTxWire(t *testing.T) {
//清空Tx消息。
	noTx := NewMsgTx(1)
	noTx.Version = 1
	noTxEncoded := []byte{
0x01, 0x00, 0x00, 0x00, //版本
0x00,                   //输入事务数的变量
0x00,                   //输出事务数的变量
0x00, 0x00, 0x00, 0x00, //锁定时间
	}

	tests := []struct {
in   *MsgTx          //要编码的邮件
out  *MsgTx          //预期的解码消息
buf  []byte          //有线编码
pver uint32          //有线编码协议版本
enc  MessageEncoding //消息编码格式
	}{
//无事务的最新协议版本。
		{
			noTx,
			noTx, noTxEncoded,
			ProtocolVersion,
			BaseEncoding,
		},

//具有多个事务的最新协议版本。
		{
			multiTx,
			multiTx,
			multiTxEncoded,
			ProtocolVersion,
			BaseEncoding,
		},

//协议版本bip0035，无事务。
		{
			noTx,
			noTx,
			noTxEncoded,
			BIP0035Version,
			BaseEncoding,
		},

//协议版本bip0035，具有多个事务。
		{
			multiTx,
			multiTx,
			multiTxEncoded,
			BIP0035Version,
			BaseEncoding,
		},

//协议版本bip0031，无事务。
		{
			noTx,
			noTx,
			noTxEncoded,
			BIP0031Version,
			BaseEncoding,
		},

//协议版本bip0031具有多个事务的版本。
		{
			multiTx,
			multiTx,
			multiTxEncoded,
			BIP0031Version,
			BaseEncoding,
		},

//协议版本NetAddressTimeVersion，无事务。
		{
			noTx,
			noTx,
			noTxEncoded,
			NetAddressTimeVersion,
			BaseEncoding,
		},

//协议版本NetAddressTimeVersion，具有多个事务。
		{
			multiTx,
			multiTx,
			multiTxEncoded,
			NetAddressTimeVersion,
			BaseEncoding,
		},

//协议版本multipleaddressversion，无事务。
		{
			noTx,
			noTx,
			noTxEncoded,
			MultipleAddressVersion,
			BaseEncoding,
		},

//具有多个事务的协议版本multipleaddressversion。
		{
			multiTx,
			multiTx,
			multiTxEncoded,
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
		var msg MsgTx
		rbuf := bytes.NewReader(test.buf)
		err = msg.BtcDecode(rbuf, test.pver, test.enc)
		if err != nil {
			t.Errorf("BtcDecode #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&msg, test.out) {
			t.Errorf("BtcDecode #%d\n got: %s want: %s", i,
				spew.Sdump(&msg), spew.Sdump(test.out))
			continue
		}
	}
}

//TestTxWireErrors对线编码和解码执行负测试
//以确认错误路径正常工作。
func TestTxWireErrors(t *testing.T) {
//在这里特别使用协议版本60002，而不是最新版本
//因为测试数据使用的是用该协议编码的字节
//版本。
	pver := uint32(60002)

	tests := []struct {
in       *MsgTx          //编码值
buf      []byte          //有线编码
pver     uint32          //有线编码协议版本
enc      MessageEncoding //消息编码格式
max      int             //引发错误的固定缓冲区的最大大小
writeErr error           //预期的写入错误
readErr  error           //预期的读取错误
	}{
//强制版本错误。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 0, io.ErrShortWrite, io.EOF},
//强制事务输入数出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 4, io.ErrShortWrite, io.EOF},
//在事务输入前一个块哈希中强制出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 5, io.ErrShortWrite, io.EOF},
//在事务输入前一个块输出索引中强制出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 37, io.ErrShortWrite, io.EOF},
//强制事务输入签名脚本长度出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 41, io.ErrShortWrite, io.EOF},
//在事务输入签名脚本中强制出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 42, io.ErrShortWrite, io.EOF},
//强制事务输入序列出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 49, io.ErrShortWrite, io.EOF},
//强制事务输出数出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 53, io.ErrShortWrite, io.EOF},
//强制事务输出值出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 54, io.ErrShortWrite, io.EOF},
//强制事务输出pk脚本长度出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 62, io.ErrShortWrite, io.EOF},
//事务输出pk脚本中的强制错误。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 63, io.ErrShortWrite, io.EOF},
//强制事务输出锁定时间出错。
		{multiTx, multiTxEncoded, pver, BaseEncoding, 206, io.ErrShortWrite, io.EOF},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
//编码为有线格式。
		w := newFixedWriter(test.max)
		err := test.in.BtcEncode(w, test.pver, test.enc)
		if err != test.writeErr {
			t.Errorf("BtcEncode #%d wrong error got: %v, want: %v",
				i, err, test.writeErr)
			continue
		}

//从有线格式解码。
		var msg MsgTx
		r := newFixedReader(test.max, test.buf)
		err = msg.BtcDecode(r, test.pver, test.enc)
		if err != test.readErr {
			t.Errorf("BtcDecode #%d wrong error got: %v, want: %v",
				i, err, test.readErr)
			continue
		}
	}
}

//testxtserialize测试MSGTX序列化和反序列化。
func TestTxSerialize(t *testing.T) {
	noTx := NewMsgTx(1)
	noTx.Version = 1
	noTxEncoded := []byte{
0x01, 0x00, 0x00, 0x00, //版本
0x00,                   //输入事务数的变量
0x00,                   //输出事务数的变量
0x00, 0x00, 0x00, 0x00, //锁定时间
	}

	tests := []struct {
in           *MsgTx //要编码的邮件
out          *MsgTx //预期的解码消息
buf          []byte //序列化数据
pkScriptLocs []int  //预期的输出脚本位置
witness      bool   //使用见证编码序列化
	}{
//没有交易。
		{
			noTx,
			noTx,
			noTxEncoded,
			nil,
			false,
		},

//多个交易。
		{
			multiTx,
			multiTx,
			multiTxEncoded,
			multiTxPkScriptLocs,
			false,
		},
//多个输出见证事务。
		{
			multiWitnessTx,
			multiWitnessTx,
			multiWitnessTxEncoded,
			multiWitnessTxPkScriptLocs,
			true,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
//序列化事务。
		var buf bytes.Buffer
		err := test.in.Serialize(&buf)
		if err != nil {
			t.Errorf("Serialize #%d error %v", i, err)
			continue
		}
		if !bytes.Equal(buf.Bytes(), test.buf) {
			t.Errorf("Serialize #%d\n got: %s want: %s", i,
				spew.Sdump(buf.Bytes()), spew.Sdump(test.buf))
			continue
		}

//反序列化事务。
		var tx MsgTx
		rbuf := bytes.NewReader(test.buf)
		if test.witness {
			err = tx.Deserialize(rbuf)
		} else {
			err = tx.DeserializeNoWitness(rbuf)
		}
		if err != nil {
			t.Errorf("Deserialize #%d error %v", i, err)
			continue
		}
		if !reflect.DeepEqual(&tx, test.out) {
			t.Errorf("Deserialize #%d\n got: %s want: %s", i,
				spew.Sdump(&tx), spew.Sdump(test.out))
			continue
		}

//确保公钥脚本位置准确。
		pkScriptLocs := test.in.PkScriptLocs()
		if !reflect.DeepEqual(pkScriptLocs, test.pkScriptLocs) {
			t.Errorf("PkScriptLocs #%d\n got: %s want: %s", i,
				spew.Sdump(pkScriptLocs),
				spew.Sdump(test.pkScriptLocs))
			continue
		}
		for j, loc := range pkScriptLocs {
			wantPkScript := test.in.TxOut[j].PkScript
			gotPkScript := test.buf[loc : loc+len(wantPkScript)]
			if !bytes.Equal(gotPkScript, wantPkScript) {
				t.Errorf("PkScriptLocs #%d:%d\n unexpected "+
					"script got: %s want: %s", i, j,
					spew.Sdump(gotPkScript),
					spew.Sdump(wantPkScript))
			}
		}
	}
}

//TestTxSerializeErrors对有线编码和解码执行负测试
//以确认错误路径正常工作。
func TestTxSerializeErrors(t *testing.T) {
	tests := []struct {
in       *MsgTx //编码值
buf      []byte //序列化数据
max      int    //引发错误的固定缓冲区的最大大小
writeErr error  //预期的写入错误
readErr  error  //预期的读取错误
	}{
//强制版本错误。
		{multiTx, multiTxEncoded, 0, io.ErrShortWrite, io.EOF},
//强制事务输入数出错。
		{multiTx, multiTxEncoded, 4, io.ErrShortWrite, io.EOF},
//在事务输入前一个块哈希中强制出错。
		{multiTx, multiTxEncoded, 5, io.ErrShortWrite, io.EOF},
//在事务输入前一个块输出索引中强制出错。
		{multiTx, multiTxEncoded, 37, io.ErrShortWrite, io.EOF},
//强制事务输入签名脚本长度出错。
		{multiTx, multiTxEncoded, 41, io.ErrShortWrite, io.EOF},
//在事务输入签名脚本中强制出错。
		{multiTx, multiTxEncoded, 42, io.ErrShortWrite, io.EOF},
//强制事务输入序列出错。
		{multiTx, multiTxEncoded, 49, io.ErrShortWrite, io.EOF},
//强制事务输出数出错。
		{multiTx, multiTxEncoded, 53, io.ErrShortWrite, io.EOF},
//强制事务输出值出错。
		{multiTx, multiTxEncoded, 54, io.ErrShortWrite, io.EOF},
//强制事务输出pk脚本长度出错。
		{multiTx, multiTxEncoded, 62, io.ErrShortWrite, io.EOF},
//事务输出pk脚本中的强制错误。
		{multiTx, multiTxEncoded, 63, io.ErrShortWrite, io.EOF},
//强制事务输出锁定时间出错。
		{multiTx, multiTxEncoded, 206, io.ErrShortWrite, io.EOF},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
//序列化事务。
		w := newFixedWriter(test.max)
		err := test.in.Serialize(w)
		if err != test.writeErr {
			t.Errorf("Serialize #%d wrong error got: %v, want: %v",
				i, err, test.writeErr)
			continue
		}

//反序列化事务。
		var tx MsgTx
		r := newFixedReader(test.max, test.buf)
		err = tx.Deserialize(r)
		if err != test.readErr {
			t.Errorf("Deserialize #%d wrong error got: %v, want: %v",
				i, err, test.readErr)
			continue
		}
	}
}

//TestTxOverflowErrors执行测试以确保对事务进行反序列化
//有意为变量使用大值
//正确处理输入和输出。否则，这可能是潜在的
//用作攻击向量。
func TestTxOverflowErrors(t *testing.T) {
//具体使用协议版本70001和事务版本1
//这里不是最新的值，因为测试数据使用
//用这些版本编码的字节。
	pver := uint32(70001)
	txVer := uint32(1)

	tests := []struct {
buf     []byte          //有线编码
pver    uint32          //有线编码协议版本
enc     MessageEncoding //消息编码格式
version uint32          //事务处理版本
err     error           //期望误差
	}{
//声称有~uint64（0）输入的事务。
		{
			[]byte{
0x00, 0x00, 0x00, 0x01, //版本
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
0xff, //输入事务数的变量
			}, pver, BaseEncoding, txVer, &MessageError{},
		},

//声称有~uint64（0）输出的事务。
		{
			[]byte{
0x00, 0x00, 0x00, 0x01, //版本
0x00, //输入事务数的变量
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
0xff, //输出事务数的变量
			}, pver, BaseEncoding, txVer, &MessageError{},
		},

//带有签名脚本的输入的事务
//声称有~uint64（0）长度。
		{
			[]byte{
0x00, 0x00, 0x00, 0x01, //版本
0x01, //输入事务数的变量
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //上一个输出哈希
0xff, 0xff, 0xff, 0xff, //前期产出指数
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
0xff, //签名脚本长度的变量
			}, pver, BaseEncoding, txVer, &MessageError{},
		},

//具有带有公钥脚本的输出的事务
//声称有~uint64（0）长度。
		{
			[]byte{
0x00, 0x00, 0x00, 0x01, //版本
0x01, //输入事务数的变量
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //上一个输出哈希
0xff, 0xff, 0xff, 0xff, //前期产出指数
0x00,                   //签名脚本长度的变量
0xff, 0xff, 0xff, 0xff, //序列
0x01,                                           //输出事务数的变量
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //交易金额
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
0xff, //公钥脚本长度的变量
			}, pver, BaseEncoding, txVer, &MessageError{},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
//从有线格式解码。
		var msg MsgTx
		r := bytes.NewReader(test.buf)
		err := msg.BtcDecode(r, test.pver, test.enc)
		if reflect.TypeOf(err) != reflect.TypeOf(test.err) {
			t.Errorf("BtcDecode #%d wrong error got: %v, want: %v",
				i, err, reflect.TypeOf(test.err))
			continue
		}

//从有线格式解码。
		r = bytes.NewReader(test.buf)
		err = msg.Deserialize(r)
		if reflect.TypeOf(err) != reflect.TypeOf(test.err) {
			t.Errorf("Deserialize #%d wrong error got: %v, want: %v",
				i, err, reflect.TypeOf(test.err))
			continue
		}
	}
}

//testxtSerializeSizeStripped执行测试以确保
//各种交易都是准确的。
func TestTxSerializeSizeStripped(t *testing.T) {
//清空Tx消息。
	noTx := NewMsgTx(1)
	noTx.Version = 1

	tests := []struct {
in   *MsgTx //TX编码
size int    //应为序列化大小
	}{
//没有输入或输出。
		{noTx, 10},

//带输入和输出的转换。
		{multiTx, 210},

//具有包括见证数据的输入的事务，以及
//一个输出。请注意，它使用的是SerializeSizeStripped，
//排除由于见证数据编码而产生的附加字节。
		{multiWitnessTx, 82},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		serializedSize := test.in.SerializeSizeStripped()
		if serializedSize != test.size {
			t.Errorf("MsgTx.SerializeSizeStripped: #%d got: %d, want: %d", i,
				serializedSize, test.size)
			continue
		}
	}
}

//testxtWitnessSize执行测试以确保
//包括见证数据的各种类型的事务都是准确的。
func TestTxWitnessSize(t *testing.T) {
	tests := []struct {
in   *MsgTx //TX编码
size int    //带见证人的预期序列化大小
	}{
//具有包括见证数据的输入的事务，以及
//一个输出。
		{multiWitnessTx, 190},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		serializedSize := test.in.SerializeSize()
		if serializedSize != test.size {
			t.Errorf("MsgTx.SerializeSize: #%d got: %d, want: %d", i,
				serializedSize, test.size)
			continue
		}
	}
}

//multix是一个具有输入和输出的MSGTX，用于各种测试。
var multiTx = &MsgTx{
	Version: 1,
	TxIn: []*TxIn{
		{
			PreviousOutPoint: OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0x31, 0xdc, 0x00, 0x1b, 0x01, 0x62,
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
0x41, //OPDA DATA65
				0x04, 0xd6, 0x4b, 0xdf, 0xd0, 0x9e, 0xb1, 0xc5,
				0xfe, 0x29, 0x5a, 0xbd, 0xeb, 0x1d, 0xca, 0x42,
				0x81, 0xbe, 0x98, 0x8e, 0x2d, 0xa0, 0xb6, 0xc1,
				0xc6, 0xa5, 0x9d, 0xc2, 0x26, 0xc2, 0x86, 0x24,
				0xe1, 0x81, 0x75, 0xe8, 0x51, 0xc9, 0x6b, 0x97,
				0x3d, 0x81, 0xb0, 0x1c, 0xc3, 0x1f, 0x04, 0x78,
				0x34, 0xbc, 0x06, 0xd6, 0xd6, 0xed, 0xf6, 0x20,
				0xd1, 0x84, 0x24, 0x1a, 0x6a, 0xed, 0x8b, 0x63,
0xa6, //65字节签名
0xac, //奥普克西格
			},
		},
		{
			Value: 0x5f5e100,
			PkScript: []byte{
0x41, //OPDA DATA65
				0x04, 0xd6, 0x4b, 0xdf, 0xd0, 0x9e, 0xb1, 0xc5,
				0xfe, 0x29, 0x5a, 0xbd, 0xeb, 0x1d, 0xca, 0x42,
				0x81, 0xbe, 0x98, 0x8e, 0x2d, 0xa0, 0xb6, 0xc1,
				0xc6, 0xa5, 0x9d, 0xc2, 0x26, 0xc2, 0x86, 0x24,
				0xe1, 0x81, 0x75, 0xe8, 0x51, 0xc9, 0x6b, 0x97,
				0x3d, 0x81, 0xb0, 0x1c, 0xc3, 0x1f, 0x04, 0x78,
				0x34, 0xbc, 0x06, 0xd6, 0xd6, 0xed, 0xf6, 0x20,
				0xd1, 0x84, 0x24, 0x1a, 0x6a, 0xed, 0x8b, 0x63,
0xa6, //65字节签名
0xac, //奥普克西格
			},
		},
	},
	LockTime: 0,
}

//multitx encoded是使用协议版本的multitx的线编码字节。
//60002，用于各种测试。
var multiTxEncoded = []byte{
0x01, 0x00, 0x00, 0x00, //版本
0x01, //输入事务数的变量
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //上一个输出哈希
0xff, 0xff, 0xff, 0xff, //前期产出指数
0x07,                                     //签名脚本长度的变量
0x04, 0x31, 0xdc, 0x00, 0x1b, 0x01, 0x62, //签名脚本
0xff, 0xff, 0xff, 0xff, //序列
0x02,                                           //输出事务数的变量
0x00, 0xf2, 0x05, 0x2a, 0x01, 0x00, 0x00, 0x00, //交易金额
0x43, //pk脚本长度的变量
0x41, //OPDA DATA65
	0x04, 0xd6, 0x4b, 0xdf, 0xd0, 0x9e, 0xb1, 0xc5,
	0xfe, 0x29, 0x5a, 0xbd, 0xeb, 0x1d, 0xca, 0x42,
	0x81, 0xbe, 0x98, 0x8e, 0x2d, 0xa0, 0xb6, 0xc1,
	0xc6, 0xa5, 0x9d, 0xc2, 0x26, 0xc2, 0x86, 0x24,
	0xe1, 0x81, 0x75, 0xe8, 0x51, 0xc9, 0x6b, 0x97,
	0x3d, 0x81, 0xb0, 0x1c, 0xc3, 0x1f, 0x04, 0x78,
	0x34, 0xbc, 0x06, 0xd6, 0xd6, 0xed, 0xf6, 0x20,
	0xd1, 0x84, 0x24, 0x1a, 0x6a, 0xed, 0x8b, 0x63,
0xa6,                                           //65字节签名
0xac,                                           //奥普克西格
0x00, 0xe1, 0xf5, 0x05, 0x00, 0x00, 0x00, 0x00, //交易金额
0x43, //pk脚本长度的变量
0x41, //OPDA DATA65
	0x04, 0xd6, 0x4b, 0xdf, 0xd0, 0x9e, 0xb1, 0xc5,
	0xfe, 0x29, 0x5a, 0xbd, 0xeb, 0x1d, 0xca, 0x42,
	0x81, 0xbe, 0x98, 0x8e, 0x2d, 0xa0, 0xb6, 0xc1,
	0xc6, 0xa5, 0x9d, 0xc2, 0x26, 0xc2, 0x86, 0x24,
	0xe1, 0x81, 0x75, 0xe8, 0x51, 0xc9, 0x6b, 0x97,
	0x3d, 0x81, 0xb0, 0x1c, 0xc3, 0x1f, 0x04, 0x78,
	0x34, 0xbc, 0x06, 0xd6, 0xd6, 0xed, 0xf6, 0x20,
	0xd1, 0x84, 0x24, 0x1a, 0x6a, 0xed, 0x8b, 0x63,
0xa6,                   //65字节签名
0xac,                   //奥普克西格
0x00, 0x00, 0x00, 0x00, //锁定时间
}

//multitxpkscriptlocs是公钥脚本的位置信息
//位于multix。
var multiTxPkScriptLocs = []int{63, 139}

//multiWitnessTx是一个带有见证数据输入的MSGTX，以及
//各种测试中使用的输出。
var multiWitnessTx = &MsgTx{
	Version: 1,
	TxIn: []*TxIn{
		{
			PreviousOutPoint: OutPoint{
				Hash: chainhash.Hash{
					0xa5, 0x33, 0x52, 0xd5, 0x13, 0x57, 0x66, 0xf0,
					0x30, 0x76, 0x59, 0x74, 0x18, 0x26, 0x3d, 0xa2,
					0xd9, 0xc9, 0x58, 0x31, 0x59, 0x68, 0xfe, 0xa8,
					0x23, 0x52, 0x94, 0x67, 0x48, 0x1f, 0xf9, 0xcd,
				},
				Index: 19,
			},
			SignatureScript: []byte{},
			Witness: [][]byte{
{ //70字节签名
					0x30, 0x43, 0x02, 0x1f, 0x4d, 0x23, 0x81, 0xdc,
					0x97, 0xf1, 0x82, 0xab, 0xd8, 0x18, 0x5f, 0x51,
					0x75, 0x30, 0x18, 0x52, 0x32, 0x12, 0xf5, 0xdd,
					0xc0, 0x7c, 0xc4, 0xe6, 0x3a, 0x8d, 0xc0, 0x36,
					0x58, 0xda, 0x19, 0x02, 0x20, 0x60, 0x8b, 0x5c,
					0x4d, 0x92, 0xb8, 0x6b, 0x6d, 0xe7, 0xd7, 0x8e,
					0xf2, 0x3a, 0x2f, 0xa7, 0x35, 0xbc, 0xb5, 0x9b,
					0x91, 0x4a, 0x48, 0xb0, 0xe1, 0x87, 0xc5, 0xe7,
					0x56, 0x9a, 0x18, 0x19, 0x70, 0x01,
				},
{ //33字节序列化发布密钥
					0x03, 0x07, 0xea, 0xd0, 0x84, 0x80, 0x7e, 0xb7,
					0x63, 0x46, 0xdf, 0x69, 0x77, 0x00, 0x0c, 0x89,
					0x39, 0x2f, 0x45, 0xc7, 0x64, 0x25, 0xb2, 0x61,
					0x81, 0xf5, 0x21, 0xd7, 0xf3, 0x70, 0x06, 0x6a,
					0x8f,
				},
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*TxOut{
		{
			Value: 395019,
PkScript: []byte{ //2WKH输出
0x00, //版本0见证程序
0x14, //OPDA DATAY20
				0x9d, 0xda, 0xc6, 0xf3, 0x9d, 0x51, 0xe0, 0x39,
				0x8e, 0x53, 0x2a, 0x22, 0xc4, 0x1b, 0xa1, 0x89,
0x40, 0x6a, 0x85, 0x23, //20字节pub密钥哈希
			},
		},
	},
}

//multiWitnessTxEncoded是包含输入的multiWitnessTx的线编码字节。
//使用协议版本70012的见证数据，并用于
//测验。
var multiWitnessTxEncoded = []byte{
0x1, 0x0, 0x0, 0x0, //版本
0x0, //表示0个输入的标记字节，或SEGMIT编码的Tx
0x1, //标志字节
0x1, //输入数量变量
	0xa5, 0x33, 0x52, 0xd5, 0x13, 0x57, 0x66, 0xf0,
	0x30, 0x76, 0x59, 0x74, 0x18, 0x26, 0x3d, 0xa2,
	0xd9, 0xc9, 0x58, 0x31, 0x59, 0x68, 0xfe, 0xa8,
0x23, 0x52, 0x94, 0x67, 0x48, 0x1f, 0xf9, 0xcd, //上一个输出哈希
0x13, 0x0, 0x0, 0x0, //小endian上一个输出索引
0x0,                    //无SIG脚本（这是见证输入）
0xff, 0xff, 0xff, 0xff, //序列
0x1,                                    //输出数量变量
0xb, 0x7, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, //输出量
0x16, //pk脚本长度的变量
0x0,  //版本0见证程序
0x14, //OPDA DATAY20
	0x9d, 0xda, 0xc6, 0xf3, 0x9d, 0x51, 0xe0, 0x39,
	0x8e, 0x53, 0x2a, 0x22, 0xc4, 0x1b, 0xa1, 0x89,
0x40, 0x6a, 0x85, 0x23, //20字节pub密钥哈希
0x2,  //见证堆栈上的两项
0x46, //70字节堆栈项
	0x30, 0x43, 0x2, 0x1f, 0x4d, 0x23, 0x81, 0xdc,
	0x97, 0xf1, 0x82, 0xab, 0xd8, 0x18, 0x5f, 0x51,
	0x75, 0x30, 0x18, 0x52, 0x32, 0x12, 0xf5, 0xdd,
	0xc0, 0x7c, 0xc4, 0xe6, 0x3a, 0x8d, 0xc0, 0x36,
	0x58, 0xda, 0x19, 0x2, 0x20, 0x60, 0x8b, 0x5c,
	0x4d, 0x92, 0xb8, 0x6b, 0x6d, 0xe7, 0xd7, 0x8e,
	0xf2, 0x3a, 0x2f, 0xa7, 0x35, 0xbc, 0xb5, 0x9b,
	0x91, 0x4a, 0x48, 0xb0, 0xe1, 0x87, 0xc5, 0xe7,
	0x56, 0x9a, 0x18, 0x19, 0x70, 0x1,
0x21, //33字节堆栈项
	0x3, 0x7, 0xea, 0xd0, 0x84, 0x80, 0x7e, 0xb7,
	0x63, 0x46, 0xdf, 0x69, 0x77, 0x0, 0xc, 0x89,
	0x39, 0x2f, 0x45, 0xc7, 0x64, 0x25, 0xb2, 0x61,
	0x81, 0xf5, 0x21, 0xd7, 0xf3, 0x70, 0x6, 0x6a,
	0x8f,
0x0, 0x0, 0x0, 0x0, //锁定时间
}

//multiWitnessTxEncodedOnZeroFlag是错误的有线编码字节，用于
//多见证Tx，包括有见证数据的输入。而不是标志字节
//如果设置为0x01，则标志为0x00，这将触发解码错误。
var multiWitnessTxEncodedNonZeroFlag = []byte{
0x1, 0x0, 0x0, 0x0, //版本
0x0, //表示0个输入的标记字节，或SEGMIT编码的Tx
0x0, //标志字节不正确（应为0x01）
0x1, //输入数量变量
	0xa5, 0x33, 0x52, 0xd5, 0x13, 0x57, 0x66, 0xf0,
	0x30, 0x76, 0x59, 0x74, 0x18, 0x26, 0x3d, 0xa2,
	0xd9, 0xc9, 0x58, 0x31, 0x59, 0x68, 0xfe, 0xa8,
0x23, 0x52, 0x94, 0x67, 0x48, 0x1f, 0xf9, 0xcd, //上一个输出哈希
0x13, 0x0, 0x0, 0x0, //小endian上一个输出索引
0x0,                    //无SIG脚本（这是见证输入）
0xff, 0xff, 0xff, 0xff, //序列
0x1,                                    //输出数量变量
0xb, 0x7, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, //输出量
0x16, //pk脚本长度的变量
0x0,  //版本0见证程序
0x14, //OPDA DATAY20
	0x9d, 0xda, 0xc6, 0xf3, 0x9d, 0x51, 0xe0, 0x39,
	0x8e, 0x53, 0x2a, 0x22, 0xc4, 0x1b, 0xa1, 0x89,
0x40, 0x6a, 0x85, 0x23, //20字节pub密钥哈希
0x2,  //见证堆栈上的两项
0x46, //70字节堆栈项
	0x30, 0x43, 0x2, 0x1f, 0x4d, 0x23, 0x81, 0xdc,
	0x97, 0xf1, 0x82, 0xab, 0xd8, 0x18, 0x5f, 0x51,
	0x75, 0x30, 0x18, 0x52, 0x32, 0x12, 0xf5, 0xdd,
	0xc0, 0x7c, 0xc4, 0xe6, 0x3a, 0x8d, 0xc0, 0x36,
	0x58, 0xda, 0x19, 0x2, 0x20, 0x60, 0x8b, 0x5c,
	0x4d, 0x92, 0xb8, 0x6b, 0x6d, 0xe7, 0xd7, 0x8e,
	0xf2, 0x3a, 0x2f, 0xa7, 0x35, 0xbc, 0xb5, 0x9b,
	0x91, 0x4a, 0x48, 0xb0, 0xe1, 0x87, 0xc5, 0xe7,
	0x56, 0x9a, 0x18, 0x19, 0x70, 0x1,
0x21, //33字节堆栈项
	0x3, 0x7, 0xea, 0xd0, 0x84, 0x80, 0x7e, 0xb7,
	0x63, 0x46, 0xdf, 0x69, 0x77, 0x0, 0xc, 0x89,
	0x39, 0x2f, 0x45, 0xc7, 0x64, 0x25, 0xb2, 0x61,
	0x81, 0xf5, 0x21, 0xd7, 0xf3, 0x70, 0x6, 0x6a,
	0x8f,
0x0, 0x0, 0x0, 0x0, //锁定时间
}

//multitxpkscriptlocs是公钥脚本的位置信息
//位于多见证Tx。
var multiWitnessTxPkScriptLocs = []int{58}

