# whicinth-Steganography-bd

风信——用图片承载无法言说的秘密

* 这里是whicinth后端服务源码
* 基于HTTP框架Hertz
* 配套[前端](https://github.com/xhdd123321/whicinth-Steganography-fd)戳这里

> 开发文档：[Re：从零开始的golang-hertz项目实战](https://zhu-an.cn/todo/Re：从零开始的golang-hertz项目实战/)

## 服务部署

### 一、本地开发

#### 配置文件
1. 在`/conf`目录下创建`dev.config.yaml`根据`default.config.yaml`完成配置
2. 修改`.env`文件`RUN_ENV = DEV`
3. 启动项目会自动注入`.env`中的环境变量并读取`dev.config.yaml`中的配置

#### 部署流程

> Go >= 1.19
1. git clone & cd
2. go mod tidy
3. make build&run

### 二、线上部署

#### 配置文件
1. 在`/conf`目录下创建`prod.config.yaml`根据`default.config.yaml`完成配置
2. 修改`.env`文件`RUN_ENV = PROD`
3. 启动项目会自动注入`.env`中的环境变量并读取`prod.config.yaml`中的配置

#### 部署流程
1. 下载golang安装包并上传至Linux: https://studygolang.com/dl
2. 安装golang1.19
```shell
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
```
3. 设置go env
```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
4. 上传项目至服务器, 安装go.mod依赖
```shell
go mod tidy
```
5. 构建二进制文件&后台启动服务
```shell
make build
make start
```

#### 终止服务
```shell
ps f | grep ./whicinth-steganography-bd # 找到 ./whicinth-steganography-bd 进程PID
kill -9 PID
```

