package main

import "github.com/pythonwithsean/k.gg/utils"

func main() {
	server := utils.NewServer(utils.ADDR, utils.PORT)
	server.Start()
}
