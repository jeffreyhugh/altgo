package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/qbxt/altgo/constants/tokens"
	"github.com/qbxt/altgo/handlers"
	"github.com/qbxt/altgo/helpers"
	"github.com/qbxt/altgo/logger"
	"github.com/qbxt/altgo/structures"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {

	logger.Init()
	helpers.InitLoginManagerChans()

	bot, err := discordgo.New("Bot " + tokens.TOKEN_DISCORD)
	if err != nil {
		logger.Fatal("could not init discord bot", err, nil)
	}
	go runDiscordBot(bot)
	go manageMinecraftLoginChannels()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func runDiscordBot(bot *discordgo.Session) {

	bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)

	if err := bot.Open(); err != nil {
		logger.Fatal("error starting Discord bot", err, nil)
	}

	bot.AddHandler(handlers.Ready)
	bot.AddHandler(handlers.MessageCreate)

	logger.Info("bot is running", logrus.Fields{"name": bot.State.User.Username, "discriminator": bot.State.User.Discriminator})
}

func manageMinecraftLoginChannels() {
	logins := make(map[string]*structures.MinecraftLogin)
	for _, login := range tokens.GetMinecraftLogins() {
		logins[strings.ToLower(login.IGN)] = login
	}
	inbox, outbox := helpers.GetLoginManagerChans()
	for {
		req := <-inbox
		if !req.Write { // read
			if login, ok := logins[strings.ToLower(req.IGN)]; ok {
				outbox <- login
			} else {
				outbox <- nil
			}
		} else { // Write to temporary cache
			logins[strings.ToLower(req.IGN)] = req
			outbox <- nil
		}
	}
}
