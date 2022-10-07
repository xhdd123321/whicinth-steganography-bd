package viper

type Config struct {
	App     *App     `yaml:"App"`
	Cronjob *Cronjob `yaml:"Cronjob"`
	Qiniu   *Qiniu   `yaml:"Qiniu"`
	Redis   *Redis   `yaml:"Redis"`
}

type App struct {
	HostPorts          string `yaml:"HostPorts"`          // 服务监听的地址和端口
	MaxRequestBodySize int    `yaml:"MaxRequestBodySize"` // 最大的请求体大小
}

type Cronjob struct {
	TempFileMinute float64 `yaml:"TempFileMinute"` // 文件上传token刷新时间（默认1h过期）
	TokenMinute    float64 `yaml:"TokenMinute"`    // 临时文件夹最长生存时间
}

type Qiniu struct {
	Domain    string `yaml:"Domain"`    // 源站域名
	AccessKey string `yaml:"AccessKey"` // AK
	SecretKey string `yaml:"SecretKey"` // SK
	Bucket    string `yaml:"Bucket"`    // 空间名称
	Prefix    string `yaml:"Prefix"`    // 保存目录
}

type Redis struct {
	Addr             string `yaml:"Addr"`             // 服务所在地址和端口
	Password         string `yaml:"Password"`         // 密码
	Db               int    `yaml:"Db"`               // 数据库编号
	EncodeLockSecond int    `yaml:"EncodeLockSecond"` // 加密锁限流间隔
	DecodeLockSecond int    `yaml:"DecodeLockSecond"` // 解密锁限流间隔
}
