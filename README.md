# whicinth-Steganography-bd

风信——用图片承载无法言说的秘密

* 这里是whicinth后端服务源码
* 基于HTTP框架Hertz
* 配套[前端](https://github.com/xhdd123321/whicinth-Steganography-fd)戳这里

> 开发文档：[Re：从零开始的golang-hertz项目实战](https://zhu-an.cn/todo/Re：从零开始的golang-hertz项目实战/)

## 简介

![image-20221021004444216](https://img.zhu-an.cn/img/20221021004444.png)

## 服务部署

### 一、本地开发

#### 配置文件
> 配置文件说明：参考`/pkg/viper/config.go`
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
> 配置文件说明：参考`/pkg/viper/config.go`
1. 在`/conf`目录下创建`prod.config.yaml`根据`default.config.yaml`完成配置
2. 修改`.env`文件`RUN_ENV = PROD`
3. 启动项目会自动注入`.env`中的环境变量并读取`prod.config.yaml`中的配置

#### 部署流程
> Linux 生产环境
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
make stop # 找到 ./whicinth-steganography-bd 进程PID并将其kill
```

#### 重启服务
```shell
make restart
```

## 服务升级
停止服务：`make stop`

上传本地项目至生产环境

更新依赖：`go mod tidy`

启动项目：`make start`

## 日志收集
使用`make start`启动项目会将日志输出至项目本地根目录`output`

- 服务启动日志：`output/start_YYYY-mm-dd.txt`
- 服务运行日志：`output/run_YYYY-mm-dd.txt`

## 性能分析

系统内置了pprof帮助完成性能分析，启动服务器后访问路由`/admin/pprof`查看当前项目的采样信息，注意生产环境下不要将该路由暴露给用户，建议配置Nginx将该路由return403

![image-20221022134233379](https://img.zhu-an.cn/img/20221022134233.png)

