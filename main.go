package main

import (
	"github.com/atef1374/video-note/bot"
	"github.com/atef1374/video-note/common"
)

func main() {
	configuration := common.LoadConfiguration()
	go common.WebServer()
	bot.LoadBot(configuration)
}
