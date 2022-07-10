# 1. 总结几种 socket 粘包的解包方式: fix length, delimiter based, length field based frame decoder。尝试举例其应用

发送数据过程：应用程序发送消息包，消息包以数据流的形式放入缓冲区，等缓冲区的数据流到达一定阈值后，再发送到网络上

接受数据过程：接受到网络过来的数据流，放入缓冲区，缓冲区的数据流到达一定阈值后，通知应用程序进行读取数据

在数据发送和接受的过程中，都是对数据流进行操作，而数据流本身没有任何开始和结束的边界。因此正确地解析出数据包，就要知道数据在流中的开始和结束位置。

## fix length frame decoder

数据发送方每次发送固定长度的数据，且不超出缓冲区，接收方获取同样长度的数据来解码拼成一个数据包。

## delimiter based frame decoder

数据发送方在数据中添加特殊的分隔符来标记边界，接收方读到分隔符时解码拼成一个数据包。

## length field based frame decoder

数据发送方在消息包头添加长度信息，接收方获取指定长度的数据来解码拼成一个数据包。

#  2.实现一个从 socket connection 中解码出 goim 协议的解码器。

[decoder.go](./decoder.go) 