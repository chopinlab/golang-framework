package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitLogConfig() error {

	if viper.GetString("profile") == "prod" {

	} else {

		/*
			Config{
				Level:            NewAtomicLevelAt(DebugLevel),
				Development:      true,
				Encoding:         "console",
				EncoderConfig:    NewDevelopmentEncoderConfig(),
				OutputPaths:      []string{"stderr"},
				ErrorOutputPaths: []string{"stderr"},
			}
		*/
		//ab := zap.NewAtomicLevelAt(zapcore.DebugLevel)
		//ab := zap.NewAtomicLevelAt(1)
		//a := zap.NewDevelopment()
		logConfig := zap.NewDevelopmentConfig()
		logConfig.Development = true
		logConfig.DisableCaller = true
	}

	return nil
}
