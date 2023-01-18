package main

import (
	"fmt"

	"github.com/VrMolodyakov/vgm/music/pkg/logging"
)

func main() {
	fmt.Println("music service start")
	logging.NewLogger("debug", nil, nil)

}
