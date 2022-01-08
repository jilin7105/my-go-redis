## 简单使用说明

- 服务端文件位置  main.go    可以使用debug 或者 go build 随意
- 客户端文件位置  TcpClient/client.go   可以使用debug 或者 go build 随意


## 客户端支持命令
- get {key}
- set {key} {value}
- info 用于查看当前内存中key的数量
- rdb  进行rdb操作

尚不完善，不支持key value中含有符号，嘿嘿

## 文件夹说明
- AOF 用户aof相关数据恢复操作
- Cmd 命令集，tcp 请求体解析
- commonSTR 通用报错文案
- DataStructure 数据类型结构体，暂时只有字符串类型
- FileAction 读写文件相关操作，维护文件io池
- MemoryManagement redis数据，维护，
- RDB 用于RDB相关数据恢复操作
- TcpClient tcp 客户端demo
- TcpHelp   tcp 请求返回处理，防止粘包-自定义协议
- TcpServer tcp 服务端监听


