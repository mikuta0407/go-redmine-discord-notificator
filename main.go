/*
Copyright (c) 2016 emersion
Copyright (c) 2024 mikuta0407
Released under the MIT license
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/emersion/go-smtp"
	"github.com/jhillyerd/enmime"
	"github.com/mikuta0407/go-redmine-discord-notificator/discord"
	"github.com/mikuta0407/go-redmine-discord-notificator/member"
)

var addr = "127.0.0.1:1025"

func init() {
	flag.StringVar(&addr, "l", addr, "Listen address")
}

type backend struct{}

func (bkd *backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &session{}, nil
}

type session struct{}

func (s *session) AuthPlain(username, password string) error {
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	return nil
}

func (s *session) Data(r io.Reader) error {
	e, err := enmime.ReadEnvelope(r)
	if err != nil {
		return err
	}

	// To -> DiscordID
	toAddr := e.GetHeader("To")
	fmt.Println(toAddr)
	discordID, err := member.ConvertToDiscordID(toAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(discordID)

	// Send DiscordDM
	subject := e.GetHeader("Subject")
	text := e.Text
	if err = discordClient.PostDM(discordID, "**"+subject+"**"+"\n"+"```\n"+text+"\n```"); err != nil {
		return err
	}
	return nil
}

func (s *session) Reset() {}

func (s *session) Logout() error {
	return nil
}

func main() {
	flag.Parse()

	token := os.Getenv("DISCORD_TOKEN")
	discordClient = discord.NewClient(token)
	s := smtp.NewServer(&backend{})

	s.Addr = addr
	s.Domain = "localhost"
	s.AllowInsecureAuth = true
	s.Debug = os.Stdout

	log.Println("Starting SMTP server at", addr)
	log.Fatal(s.ListenAndServe())

}

var discordClient *discord.Client
