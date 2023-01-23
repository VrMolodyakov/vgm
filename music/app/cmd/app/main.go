package main

import (
	_ "context"
	_ "fmt"

	_ "github.com/VrMolodyakov/vgm/music/app/internal/config"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

func main() {
	logging.Init("info", "log.txt")
	logger := logging.GetLogger()
	logger.Infow("album", "sql", "args")
	logger.Info("msg")
	// if err != nil {
	// 	fmt.Printf("cannot start app due to %v", err)
	// 	return
	// }
	// logger.Info("music service start")
	// //cfg := config.GetConfig()
	// //logger.Sugar().Info(cfg)
	// ctx := context.Background()
	// logging.ContextWithLogger(ctx,logging.New())
	// logger.Info("music service end")

}
