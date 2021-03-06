
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 20:02:50</date>
//</624461720596779008>

//版权所有（c）2015 BTCSuite开发者
//此源代码的使用由ISC控制
//可以在许可文件中找到的许可证。

/*
包btcjson提供了使用比特币json-rpc api的原语。

概述

当通过JSON-RPC协议进行通信时，所有命令都需要
以适当的格式排列到线和线之间。这个包裹
提供数据结构和原语以简化此过程。

此外，它还提供一些附加功能，如自定义命令
注册、命令分类和基于反射的帮助生成。

JSON-RPC协议概述

使用此包不需要此信息，但它确实需要
提供一些关于编组和解组的直觉
下面讨论的是在引擎盖下进行。

正如JSON-RPC规范所定义的，在
电线：

  -请求对象
    “jsonrpc”：“1.0”，“id”：“someid”，“method”：“somemethod”，“params”：[someparams]
    注意：通知的格式相同，但id字段为空。

  -响应对象
    “result”：something，“error”：null，“id”：“someid”
    “result”：空，“error”：“code”：somint，“message”：somestring，“id”：“someid”

对于请求，根据
正在发送的方法（a.k.a.命令）。每个参数可以像int一样简单
或者包含许多嵌套字段的复杂结构。ID字段用于
确定一个请求，并将包含在相关响应中。

使用异步传输（如WebSockets）时，自发
通知也是可能的。如图所示，它们与请求相同
对象，但它们的ID字段设置为空。因此，服务器将
忽略ID字段设置为空的请求，而客户端可以选择
消费或忽略它们。

不幸的是，原来的比特币JSON-RPC API（因此任何兼容的
使用它）并不总是遵循规范，有时会返回错误
结果字段中的字符串，某些命令有空错误。然而，
在大多数情况下，错误字段将按失败时所述进行设置。

编组和解编

根据上面的讨论，应该很容易看到这类的类型
包映射到协议的必需部分

  -请求对象（类型请求）
    -命令（type<foo>cmd）
    -通知（类型<foo>ntfn）
  -响应对象（类型响应）
    -结果（类型<foo>result）

为了简化请求和响应的编组，marshalCmd和
提供了MarshalResponse函数。它们返回准备就绪的原始字节
穿过电线。

解组接收到的请求对象是一个两步过程：
  1）通过json将原始字节解封为请求结构实例。解封
  2）在未合并请求的结果字段上使用UnmarshalCmd创建
     设置了所有结构字段的具体命令或通知实例
     因此

使用这种方法是因为它为调用者提供了访问
请求中不属于命令的字段，如ID。

解组接收到的响应对象也是一个两步过程：
  1）通过json.unmashal将原始字节解组为响应结构实例。
  2）根据ID，取消标记未标记的结果字段
     创建具体类型实例的响应

如上所述，使用这种方法是因为它为调用者提供了对
响应中的字段，如ID和错误。

命令创建

这个包提供了两种创建新命令的方法。首先，
首选的方法是使用一个新的<foo>cmd函数。这允许
静态编译时检查有助于确保参数与
结构定义。

第二种方法是使用方法（命令）名称的newcmd函数
和变量参数。功能包括全面检查以确保
根据提供的方法，参数是准确的，但是这些检查是，
显然，运行时意味着在代码
实际执行。但是，它对于用户提供的命令非常有用
这是有意动态的。

自定义命令注册

这个包的命令处理是围绕注册的概念构建的。
命令。这对于已经由
包，但这也意味着调用者可以轻松地提供自定义命令
具有与内置命令相同的功能。使用registerCmd
为此目的发挥作用。

所有注册方法的列表都可以通过registeredCmdMethods获得。
功能。

司令部检查

所有已注册的命令都使用标识信息的标志进行注册，例如
无论该命令是应用于链服务器、钱包服务器还是
通知以及要使用的方法名。可以获得这些标志
使用methodUsageFlags标志，可以使用
CmdMethod函数。

帮助生成

为了方便为RPC服务器的用户提供一致的帮助，此包
显示GenerateHelp和函数，该函数使用对已注册的
命令或通知以及提供的预期结果类型
生成最终帮助文本。

此外，还提供了methodusageText函数以生成一致的
使用反射的已注册命令和通知的单行用法。

错误

此包支持两种不同类型的错误：

  -有关编组或解编或不当使用的一般错误
    包（类型错误）
  -rpc错误，作为
    json-rpc响应（rpcerror类型）

第一类错误（类型错误）通常表示程序员错误
并且可以通过正确使用API来避免。此类型的错误将是
从该包中可用的各种函数返回。他们识别
不支持的字段类型、尝试注册格式错误的命令，
并尝试创建参数数目不正确的新命令。
通过将错误断言为
*btcjson.error并访问errorcode字段。

另一方面，第二类错误（rpcerror类型）对于
将错误返回到RPC客户端。因此，它们在以前的
描述的响应类型。
**/

package btcjson

