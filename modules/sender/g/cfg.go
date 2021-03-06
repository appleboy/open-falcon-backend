package g

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/toolkits/file"
	"sync"
)

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type RedisConfig struct {
	Addr    string `json:"addr"`
	MaxIdle int    `json:"maxIdle"`
}

type QueueConfig struct {
	Sms        string `json:"sms"`
	Mail       string `json:"mail"`
	QQ         string `json:"qq"`
	Serverchan string `json:"serverchan"`
}

type WorkerConfig struct {
	Sms        int `json:"sms"`
	Mail       int `json:"mail"`
	QQ         int `json:"qq"`
	Serverchan int `json:"serverchan"`
}

type ApiConfig struct {
	Sms        string `json:"sms"`
	Mail       string `json:"mail"`
	QQ         string `json:"qq"`
	Serverchan string `json:"serverchan"`
}

type GlobalConfig struct {
	Debug  bool          `json:"debug"`
	Http   *HttpConfig   `json:"http"`
	Redis  *RedisConfig  `json:"redis"`
	Queue  *QueueConfig  `json:"queue"`
	Worker *WorkerConfig `json:"worker"`
	Api    *ApiConfig    `json:"api"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c
	log.Println("read config file:", cfg, "successfully")
}
