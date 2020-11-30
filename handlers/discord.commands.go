package handlers

import (
	"fmt"
	mc "github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/yggdrasil"
	"github.com/bwmarrin/discordgo"
	"github.com/qbxt/altgo/helpers"
	"github.com/qbxt/altgo/logger"
	"github.com/qbxt/altgo/structures"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

//func LoginCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
//	args := strings.Split(m.Content, " ")
//	// /login Xx_Example_xX
//	//    0        1
//	if len(args) == 1 { // did not provide IGN
//		_, _ = s.ChannelMessageSend(m.ChannelID, "No name provided. Syntax: `/login <IGN>`")
//		return
//	}
//
//	// Get current session or create new one
//	inbox, outbox := helpers.GetLoginManagerChans()
//	inbox <- &structures.MinecraftLogin{
//		IGN:      args[1],
//		Password: "",
//	}
//	currentSession := <-outbox
//	if currentSession == nil { // username not found
//		_, _ = s.ChannelMessageSend(m.ChannelID, "Could not find provided username.")
//		return
//	}
//
//	if currentSession.Client != nil {
//		_, _ = s.ChannelMessageSend(m.ChannelID, "This user is already logged in. To relog, type `/relog`. To logout, type `/logout`")
//		return
//	}
//
//	c := mc.NewClient()
//	if currentSession.Migrated {
//		auth, err := yggdrasil.Authenticate(currentSession.Email, currentSession.Password)
//		if err != nil {
//			logger.Error("Could not log in", err, logrus.Fields{"username": currentSession.IGN})
//			_, _ = s.ChannelMessageSend(m.ChannelID, "Could not log in. Check the console for more details.")
//			return
//		}
//		c.Auth.UUID, c.Name = auth.SelectedProfile()
//		c.AsTk = auth.AccessToken()
//	} else {
//		auth, err := yggdrasil.Authenticate(currentSession.IGN, currentSession.Password)
//		if err != nil {
//			logger.Error("Could not log in", err, logrus.Fields{"username": currentSession.IGN})
//			_, _ = s.ChannelMessageSend(m.ChannelID, "Could not log in. Check the console for more details.")
//			return
//		}
//		c.Auth.UUID, c.Name = auth.SelectedProfile()
//		c.AsTk = auth.AccessToken()
//	}
//	currentSession.Client = c
//	currentSession.Password = "a"
//	inbox <- currentSession
//	_ = <-outbox
//
//	logger.Info("logged in", logrus.Fields{"ign": currentSession.IGN})
//	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Successfully logged in as **%s** (UUID %s)", currentSession.IGN, currentSession.Client.Auth.UUID))
//}

func ConnectCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	// /join Xx_Example_xX mc.hypixel.net
	//   0        1             2
	if len(args) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "No name and/or no server provided. Syntax: `/join <IGN> <server>` or `/join <IGN> <server:port>`")
		return
	}

	inbox, outbox := helpers.GetLoginManagerChans()
	inbox <- &structures.MinecraftLogin{
		IGN:      args[1],
		Password: "",
	}
	currentSession := <-outbox
	if currentSession == nil { // username not found
		_, _ = s.ChannelMessageSend(m.ChannelID, "Could not find provided username.")
		return
	}

	if currentSession.Client != nil { // logged in elsewhere
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("This user is logged in at `%s:%d`. Please type `/disconnect %s`.", currentSession.Server.Hostname, currentSession.Server.Port, args[1]))
		return
	}

	c := mc.NewClient()
	if currentSession.Migrated {
		auth, err := yggdrasil.Authenticate(currentSession.Email, currentSession.Password)
		if err != nil {
			logger.Error("Could not log in", err, logrus.Fields{"username": currentSession.IGN})
			_, _ = s.ChannelMessageSend(m.ChannelID, "Could not log in. Check the console for more details.")
			return
		}
		c.Auth.UUID, c.Name = auth.SelectedProfile()
		c.AsTk = auth.AccessToken()
	} else {
		auth, err := yggdrasil.Authenticate(currentSession.IGN, currentSession.Password)
		if err != nil {
			logger.Error("Could not log in", err, logrus.Fields{"username": currentSession.IGN})
			_, _ = s.ChannelMessageSend(m.ChannelID, "Could not log in. Check the console for more details.")
			return
		}
		c.Auth.UUID, c.Name = auth.SelectedProfile()
		c.AsTk = auth.AccessToken()
	}
	currentSession.Client = c
	currentSession.Password = "a"

	serverArgs := strings.Split(args[2], ":")
	hostname := serverArgs[0]
	port := 25565
	if len(serverArgs) == 2 {
		var err error
		port, err = strconv.Atoi(serverArgs[1])
		if err != nil {
			logger.Error("could not atoi", err, nil)
			_, _ = s.ChannelMessageSend(m.ChannelID, "Could not use `Atoi()`. Check the console for more details.")
			return
		}
	}

	if err := currentSession.Client.JoinServer(hostname, port); err != nil {
		logger.Error("could not join server", err, nil)
		_, _ = s.ChannelMessageSend(m.ChannelID, "Could not join the specified server. Check the console for more details.")
		return
	}

	connectedAt := time.Now()
	currentSession.Server.ConnectedAt = &connectedAt
	currentSession.Server.Hostname = hostname
	currentSession.Server.Port = port

	logger.Info("joined game", nil)
	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Successfully joined %s:%d", hostname, port))

	currentSession.Client.Events.GameStart = func() error { return OnGameStart(currentSession) }
	currentSession.Client.Events.Die = func() error { return OnDie(currentSession) }
	currentSession.Client.Events.Disconnect = func(msg chat.Message) error { return OnDisconnect(currentSession, msg) }

	// Handle c.Events.ReceivePacket
	// debug
	// currentSession.Client.Events.ReceivePacket = func(p packet.Packet) (bool, error) {return OnPacket(currentSession, p)}

	inbox <- currentSession
	_ = <-outbox

	if err := currentSession.Client.HandleGame(); err != nil {
		logger.Error("error handling game", err, logrus.Fields{"hostname": hostname, "port": port})
	}
}

func DisconnectCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	// /login Xx_Example_xX
	//    0        1
	if len(args) == 1 { // did not provide IGN
		_, _ = s.ChannelMessageSend(m.ChannelID, "No name provided. Syntax: `/disconnect <IGN>`")
		return
	}

	// Get current session
	inbox, outbox := helpers.GetLoginManagerChans()
	inbox <- &structures.MinecraftLogin{
		IGN:      args[1],
		Password: "",
	}
	currentSession := <-outbox
	if currentSession == nil { // username not found
		_, _ = s.ChannelMessageSend(m.ChannelID, "Could not find provided username.")
		return
	}

	if currentSession.Client == nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("This user is not connected to any servers. Please type `/join %s <server IP>`.", args[1]))
		return
	}

	_ = currentSession.Client.Disconnect()

	currentSession.Client = nil
	currentSession.Server.ConnectedAt = nil
	currentSession.Server.Hostname = ""
	currentSession.Server.Port = 0


	logger.Info("disconnected", logrus.Fields{"ign": currentSession.IGN})
	_, _ = s.ChannelMessageSend(m.ChannelID, "Disconnected from server")
}

func ChatCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	// /chat Xx_Example_xX /tpa Example69
	//   0        1          2      3
	if len(args) < 2 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "No name and/or no message provided. Syntax: `/join <IGN> <server>` or `/join <IGN> <server:port>`")
		return
	}
	message := strings.Join(args[2:], " ")
	if len(message) > 256 {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Message is longer than 256 characters.")
		return
	}

	inbox, outbox := helpers.GetLoginManagerChans()
	inbox <- &structures.MinecraftLogin{
		IGN:      args[1],
		Password: "",
	}
	currentSession := <-outbox
	if currentSession == nil { // username not found
		_, _ = s.ChannelMessageSend(m.ChannelID, "Could not find provided username.")
		return
	}

	if currentSession.Client == nil { // not logged in
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("This user is not logged in. Please type `/login %s`.", args[1]))
		return
	}

	if err := currentSession.Client.Chat(message); err != nil {
		logger.Error("could not send chat message", err, logrus.Fields{"ign": currentSession.IGN})
		_, _ = s.ChannelMessageSend(m.ChannelID, "Could not send chat message. Check the console for more details.")
		return
	}

	logger.Info("sent chat message", logrus.Fields{"ign": currentSession.IGN, "message": message})
	_, _ = s.ChannelMessageSend(m.ChannelID, "Sent message to server")

	return
}
