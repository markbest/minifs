## minifs
Golang实现的简单文件存储工具

## 使用方法
- 编译minifs并加入path中

```
go build -o minifs main.go
```
- 创建master服务器（host：主机地址，port：端口号）

```
minifs master -host=127.0.0.1 -port=1234
```
- 创建node服务器（host：主机地址，port：端口号）

```
minifs node -host=127.0.0.1 -port=1235
```
- 上传文件（dir：目录，include：上传文件类型）

```
minifs upload file1 file2 ...
minifs upload -dir=test -include=.pdf
```  

## 核心实现
- 使用grpc各个节点之间的通信
- 使用leveldb进行文件存储
- 使用zookeeper进行分布式服务注册和发现
- 使用随机分配请求到各个节点

## 可用命令
- minifs master -host=[host] -port=[port] 
- minifs node -host=[host] -port=[port] 
- minifs upload file1 file2 ...
- minifs upload -dir=[dir] -include=[include]

## To Do
- [x] 多个文件上传
- [x] 整个目录批量上传
- [x] 使用分布式进行功能扩展
- [ ] 已上传文件的列表展示
- [ ] Master服务器的负载、属性展示
- [ ] 分布式节点之间文件存储的同步