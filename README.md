## minifs
Golang实现的简单文件存储工具

## 使用方法
- 编译minifs并加入path中
```
go build -o minifs main.go
```
- 创建master服务器（port：端口号，dir：存储路径）
```
minifs server -port=1234 -dir=/data
```
- 上传文件（host：master IP，port：master端口号，dir：目录，include：上传文件类型）
```
minifs upload -host=127.0.0.1 -port=1234 file1 file2...
minifs upload -host=127.0.0.1 -port=1234 -dir=test -include=.pdf
```