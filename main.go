package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang-framework/configs"
	"net/http"
)

func main() {

	if err := configs.InitAppConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//logger, _ := zap.NewProduction()
	//logger, _ := zap.NewDevelopment()
	//configaaa := zap.NewDevelopmentConfig()
	//configaaa.
	//a, _ := configaaa.Build()

	//print(configaaa.Encoding)
	devConfig := zap.NewDevelopmentConfig()
	devConfigEncoder := zap.NewDevelopmentEncoderConfig()
	devConfigEncoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
	devConfig.EncoderConfig = devConfigEncoder
	logger, _ := devConfig.Build()

	//print(devConfig.CallerKey)
	//

	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()
	//slogger := logger.Sugar()
	sugar.Info("Info() uses sprint")
	sugar.Infof("Infof() uses %s", "sprintf")
	sugar.Infow("Infow() allows tags", "name", "Legolas", "type", 1)
	/*sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "hihihi",
		"attempt", 3,
		"backoff", time.Second,
	)*/
	sugar.Infof("Failed to fetch URL: %s", "hihi")

	/*if err := configs.InitLogConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}*/

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Access-Control-Allow-Origin
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//viper.SetDefault("test", "Hyungwoo")
	//print(viper.GetString("test"))
	//log.Print(viper.GetString("test"))

	//viper.SetConfigName("config_local")
	//viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	//viper.AddConfigPath(".")
	/*err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}*/

	//log.Print(viper.GetString("profile"))
	//log.Print(os.Getenv("profile"))

	/*g := new(errgroup.Group)
	g.Go(func() error {
		return e.Start(":8080") })
	//g.Go(func() error { return serveFileServer() })
	err := g.Wait()
	if err != nil {
		panic(err.Error())
	}*/
	//e.Logger.Fatal(e.Start(":8080"))

	//println(viper.GetString("dbconnection"))
}
