package config

import (
	"github.com/spf13/viper"
	"github.com/Yamiyo/account/utils/log"
	"os"
)

// Config ...
var (
	Config *Setup
)

// LoadConfig ...
func LoadConfig(file string) error {
	if Config != nil {
		return nil
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return err
	}

	Config = new(Setup)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s", err)
	}

	viper.Unmarshal(&Config)

	return nil
}

// InitConfig ...
func InitConfig() {
	if err := LoadConfig("app.yaml"); err != nil {
		panic(err)
	}
}

// Monitor ...
type Monitor struct {
	Brands                 []string `yaml:"Brands"`
	HourMonitorBeforeHour  int      `yaml:"HourMonitorBeforeHour"`
	TenMonitorBeforeMinute int      `yaml:"TenMonitorBeforeMinute"`
}

// Setup ...
type Setup struct {
	ConsumerConfig ConsumerConfig `yaml:"ConsumerConfig"`
	DatabaseConfig DatabaseConfig `yaml:"DatabaseConfig"`
	LogConfig      LogConfig      `yaml:"LogConfig"`
	NotifyConfig   NotifyConfig   `yaml:"NotifyConfig"`
	JobConfig      JobConfig      `yaml:"JobConfig"`
	GINConfig      GINConfig      `yaml:"GINConfig"`
}

// ConsumerConfig ...
type ConsumerConfig struct {
	Brokers   []string `yaml:"Brokers"`
	Increment string   `yaml:"Increment"`
}

// Databases ...
type Databases struct {
	Password string   `yaml:"Password"`
	DataBase string   `yaml:"DataBase"`
	Name     string   `yaml:"Name"`
	Address  []string `yaml:"Address"`
	Username string   `yaml:"Username"`
}

// NotifyConfig ...
type NotifyConfig struct {
	Slack Slack `yaml:"Slack"`
}

// Slack ...
type Slack struct {
	API     string `yaml:"API"`
	Channel string `yaml:"Channel"`
	Hook    bool   `yaml:"Hook"`
}

// JobConfig ...
type JobConfig struct {
	Refresh Refresh `yaml:"Refresh"`
	Monitor Monitor `yaml:"Monitor"`
	Clean   []Clean `yaml:"Clean"`
}

// DatabaseConfig ...
type DatabaseConfig struct {
	Databases []Databases `yaml:"Databases"`
}

// LogConfig ...
type LogConfig struct {
	HistoryPath   string `yaml:"HistoryPath"`
	FullColor     bool   `yaml:"FullColor"`
	FullTimestamp bool   `yaml:"FullTimestamp"`
	Name          string `yaml:"Name"`
	Env           string `yaml:"Env"`
	Level         string `yaml:"Level"`
	Duration      string `yaml:"Duration"`
}

// Clean ...
type Clean struct {
	Topic          string `yaml:"Topic"`
	CleanDailyHour int    `yaml:"CleanDailyHour"`
	BeforeHours    int    `yaml:"BeforeHours"`
	Name           string `yaml:"Name"`
}

// GINConfig ...
type GINConfig struct {
	Address []string `yaml:"Address"`
}

// Refresh
type Refresh struct {
	Size        int `yaml:"Size"`
	IntervalMin int `yaml:"IntervalMin"`
}
