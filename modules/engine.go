package modules

import (
	"fmt"
	"nucleus/modules/proxy"
	"nucleus/utils"
)

func Start() {

	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	proxy.Server(&config.Rotues)

}
