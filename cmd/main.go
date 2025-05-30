package main

import (
	"os"

	"github.com/gyu-young-park/StoryShift/internal/config"
	"github.com/gyu-young-park/StoryShift/internal/injector"
	"github.com/gyu-young-park/StoryShift/pkg/file"
	"github.com/gyu-young-park/StoryShift/pkg/log"
	"github.com/gyu-young-park/StoryShift/pkg/server"
	"go.uber.org/zap"
)

func main() {
	logger := log.GetLogger()
	injector.Initialize()
	if _, err := os.Stat(file.TEMP_DIR); os.IsNotExist(err) {
		os.Mkdir(file.TEMP_DIR, 0755)
	}

	logger.Info("App starts", zap.String("hellp", "world"))
	server.Start(config.Manager.ConfigModel)
}
