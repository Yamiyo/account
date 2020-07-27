package server

import (
	"github.com/Yamiyo/account/glob"
	"github.com/Yamiyo/account/glob/config"
	"github.com/Yamiyo/account/transport/http"
	"github.com/Yamiyo/account/utils/log"

	"github.com/k0kubun/pp"
)

// Run ...
func Run() error {
	if err := glob.Inject(); err != nil {
		return err
	}

	// Error handling setup
	log.Init(config.Config.LogConfig.Env,
		config.Config.LogConfig.Level,
		config.Config.LogConfig.HistoryPath,
		config.Config.LogConfig.Duration,
		config.Config.NotifyConfig.Slack.API,
		config.Config.NotifyConfig.Slack.Channel,
		config.Config.NotifyConfig.Slack.Hook,
		config.Config.LogConfig.FullColor,
		config.Config.LogConfig.FullTimestamp)

	conf := pp.Sprintln(config.Config)
	log.Debug(conf)

	if err := glob.InitListener(); err != nil {
		return err
	}

	if err := glob.InitJob(); err != nil {
		return err
	}
	//job.InitJobImpl()

	if err := glob.InitController(); err != nil {
		return err
	}

	router, err := http.InitController()
	if err != nil {
		return err
	}

	log.Info("Server startup")
	for _, port := range config.Config.GINConfig.Address {
		if err := router.Run(port); err != nil {
			continue
		} else {
			break
		}
	}

	return nil
}
