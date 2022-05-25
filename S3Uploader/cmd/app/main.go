package app

import (
	"fmt"
	"github.com/monferon/fsm/loader/config"
	"github.com/monferon/fsm/loader/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	app.Run(cfg)
}
