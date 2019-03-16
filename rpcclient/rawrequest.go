
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:56</date>
//</624461742142918656>

//版权所有（c）2014-2017 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

package rpcclient

import (
	"encoding/json"
	"errors"

	"github.com/btcsuite/btcd/btcjson"
)

//futurerawresult是未来交付rawrequest rpc结果的承诺。
//调用（或适用的错误）。
type FutureRawResult chan *response

//接收等待未来承诺的响应并返回原始
//响应，或者请求失败时出错。
func (r FutureRawResult) Receive() (json.RawMessage, error) {
	return receiveFuture(r)
}

//rawrequestasync返回可用于获取
//通过调用接收在将来某个时间的自定义RPC请求的结果
//函数。
//
//有关阻止版本和更多详细信息，请参阅rawrequest。
func (c *Client) RawRequestAsync(method string, params []json.RawMessage) FutureRawResult {
//方法不能为空。
	if method == "" {
		return newFutureError(errors.New("no method"))
	}

//当没有参数时，将参数封送为“[]”而不是“null”
//通过。
	if params == nil {
		params = []json.RawMessage{}
	}

//使用提供的方法和参数创建原始JSON-RPC请求
//把它整理好。这样做而不是使用sendcmd函数
//因为这依赖于编组注册的btcjson命令，而不是
//而不是自定义命令。
	id := c.NextID()
	rawRequest := &btcjson.Request{
		Jsonrpc: "1.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}
	marshalledJSON, err := json.Marshal(rawRequest)
	if err != nil {
		return newFutureError(err)
	}

//生成请求并将其与要响应的通道一起发送。
	responseChan := make(chan *response, 1)
	jReq := &jsonRequest{
		id:             id,
		method:         method,
		cmd:            nil,
		marshalledJSON: marshalledJSON,
		responseChan:   responseChan,
	}
	c.sendRequest(jReq)

	return responseChan
}

//RawRequest allows the caller to send a raw or custom request to the server.
//此方法可用于发送和接收
//此客户端包未处理或部分代理的请求
//如果一个请求不能
//直接处理。
func (c *Client) RawRequest(method string, params []json.RawMessage) (json.RawMessage, error) {
	return c.RawRequestAsync(method, params).Receive()
}

