# whicinth-Steganography-bd

风信——用图片承载无法言说的秘密

* 这里是whicinth后端服务源码
* 基于HTTP框架Hertz
* 配套[前端](https://github.com/xhdd123321/whicinth-Steganography-fd)戳这里

> 开发文档：[Re：从零开始的golang-hertz项目实战](https://zhu-an.cn/todo/Re：从零开始的golang-hertz项目实战/)

## 服务部署

### 部署运行
```shell
1. 下载golang安装包并上传至Linux: https://studygolang.com/dl
2. rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
3. go env -w GO111MODULE=on
4. go env -w GOPROXY=https://goproxy.cn,direct
5. go mod tidy
6. make build
7. make start
```

### 终止
```shell
ps aux
kill -9 PID
```

