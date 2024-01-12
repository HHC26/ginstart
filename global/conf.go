package global

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	System struct {
		AppName        string `json:"appname" mapstructure:"appname"`
		Resource       string `json:"resource" mapstructure:"resource"`
		Port           int    `json:"port" mapstructure:"port"`
		RunMode        string `json:"run-mode" mapstructure:"run-mode"`
		DbType         string `json:"db-type" mapstructure:"db-type"`
		OssType        string `json:"oss-type" mapstructure:"oss-type"`
		UseRedis       bool   `json:"use-redis" mapstructure:"use-redis"`
		UberlimitCount int    `json:"uberlimit-count" mapstructure:"uberlimit-count"`
		IplimitCount   int    `json:"iplimit-count" mapstructure:"iplimit-count"`
		IplimitTime    int    `json:" iplimit-time" mapstructure:" iplimit-time"`
	}

	FileUpload struct {
		Directory string `json:"directory" mapstructure:"directory"`
		ImageSize int    `json:"image-size" mapstructure:"image-size"`
		ImageExt  string `json:" image-ext" mapstructure:" image-ext"`
		VideoSize int    `json:"video-size" mapstructure:"video-size"`
		VideoExt  string `json:"video-ext" mapstructure:"video-ext"`
	}

	Jwt struct {
		SigningKey  string `json:"signing-key" mapstructure:"signing-key"`
		Issuer      string `json:"Issuer" mapstructure:"Issuer"`
		ExpiresTime int    `json:"expires-time" mapstructure:"expires-time"`
		BufferTime  int    `json:"buffer-time" mapstructure:"buffer-time"`
	}

	Mysql struct {
		Username     string `json:"username" mapstructure:"username"`
		Password     string `json:"password" mapstructure:"password"`
		Path         string `json:"path" mapstructure:"path"`
		DbName       string `json:"db-name" mapstructure:"db-name"`
		Config       string `json:"config" mapstructure:"config"`
		MaxIdleConns int    `json:" max-idle-conns" mapstructure:" max-idle-conns"`
		MaxOpenConns int    `json:"max-open-conns" mapstructure:"max-open-conns"`
		LogMode      bool   `json:"log-mode" mapstructure:"log-mode"`
	}

	Sqlite struct {
		DbName  string `json:"db-name" mapstructure:"db-name"`
		Config  string `json:"config" mapstructure:"config"`
		LogMode bool   `json:"log-mode" mapstructure:"log-mode"`
	}

	Redis struct {
		Addr     string `json:"addr" mapstructure:"addr"`
		Password string `json:"password" mapstructure:"password"`
		DB       int    `json:"db" mapstructure:"db"`
		Prefix   string `json:"prefix" mapstructure:"prefix"`
	}

	Captcha struct {
		KeyLong     int    `json:"key-long" mapstructure:"key-long"`
		ImgWidth    int    `json:"img-width" mapstructure:"img-width"`
		ImgHeightng int    `json:"img-height" mapstructure:"img-height"`
		ImgType     string `json:"img-type" mapstructure:"img-type"`
	}

	Zap struct {
		Level      string `json:"level" mapstructure:"level"`
		Format     string `json:"format" mapstructure:"format"`
		Path       string `json:"path" mapstructure:"path"`
		MaxSize    int    `json:"max-size" mapstructure:"max-size"`
		MaxBackups int    `json:"max-backups" mapstructure:"max-backups"`
		MaxAge     int    `json:"max-age" mapstructure:"max-age"`
		ShowLine   bool   `json:"show-line" mapstructure:"show-line"`
		Compress   bool   `json:"compress" mapstructure:"compress"`
	}

	Local struct {
		Path      string `json:"path" mapstructure:"path"`
		StorePath string `json:"store-path" mapstructure:"store-path"`
	}

	Qiniu struct {
		Zone          string `json:"zone" mapstructure:"zone"`
		Bucket        string `json:"bucket" mapstructure:"bucket"`
		ImgPath       string `json:"img-path" mapstructure:"img-path"`
		UseHttps      bool   `json:"use-https" mapstructure:"use-https"`
		AccessKey     string `json:"access-key" mapstructure:"access-key"`
		SecretKey     string `json:"secret-key" mapstructure:"secret-key"`
		UseCdnDomains bool   `json:"use-cdn-domains" mapstructure:"use-cdn-domains"`
	}
}

func InitConfig(configYaml string) *AppConfig {
	v := viper.New()
	v.SetConfigFile(configYaml)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err))
	}

	//将配置文件读到结构体中
	var _config *AppConfig
	err = v.Unmarshal(&_config)
	if err != nil {
		fmt.Println("配置文件转实体出现错误" + err.Error())
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&_config); err != nil {
			fmt.Println(err)
		}
	})
	return _config
}
