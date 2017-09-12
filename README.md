## minifs
mini版的简单文件存储工具

## 使用方法
- 编译minifs并加入path中
```
go build -o minifs main.go
```
- 创建master服务器
```
minifs server -port=1234 -dir=/data
```
- 上传文件
```
minifs upload -host=127.0.0.1 -port=1234 file
```