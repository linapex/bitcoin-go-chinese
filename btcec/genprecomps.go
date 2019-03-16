
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:50</date>
//</624461717388136448>

//
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

//
//
//
//

package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	fi, err := os.Create("secp256k1.go")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

//
	serialized := btcec.S256().SerializedBytePoints()
	var compressed bytes.Buffer
	w := zlib.NewWriter(&compressed)
	if _, err := w.Write(serialized); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w.Close()

//
	encoded := make([]byte, base64.StdEncoding.EncodedLen(compressed.Len()))
	base64.StdEncoding.Encode(encoded, compressed.Bytes())

fmt.Fprintln(fi, "//
fmt.Fprintln(fi, "//
fmt.Fprintln(fi, "//可以在许可证文件中找到的许可证。“）
	fmt.Fprintln(fi)
	fmt.Fprintln(fi, "package btcec")
	fmt.Fprintln(fi)
fmt.Fprintln(fi, "//自动生成的文件（请参见genprecomps.go）“）
fmt.Fprintln(fi, "//不要编辑“”
	fmt.Fprintln(fi)
	fmt.Fprintf(fi, "var secp256k1BytePoints = %q\n", string(encoded))

	a1, b1, a2, b2 := btcec.S256().EndomorphismVectors()
	fmt.Println("The following values are the computed linearly " +
		"independent vectors needed to make use of the secp256k1 " +
		"endomorphism:")
	fmt.Printf("a1: %x\n", a1)
	fmt.Printf("b1: %x\n", b1)
	fmt.Printf("a2: %x\n", a2)
	fmt.Printf("b2: %x\n", b2)
}

