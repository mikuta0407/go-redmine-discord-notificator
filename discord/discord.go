package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

type PostChannelRequestJSON struct {
	RecipientID string `json:"recipient_id"`
}

type PostChannelResponseJSON struct {
	ChannelID string `json:"id"`
}

type PostDMRequestJSON struct {
	Content string `json:"content"`
}

// DM送信
func (c *Client) PostDM(discordID, content string) error {
	// channnelid取得
	channelID, err := c.getChannelID(discordID)
	if err != nil {
		return err
	}
	fmt.Println(channelID)

	// post
	if err := c.send(channelID, content); err != nil {
		return err
	}

	return nil
}

// ChannelID取得
func (c *Client) getChannelID(discordID string) (string, error) {
	reqBody, err := json.Marshal(PostChannelRequestJSON{
		RecipientID: discordID,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		"POST",
		"https://discordapp.com/api/users/@me/channels",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("authorization", "Bot "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBodyRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var res PostChannelResponseJSON

	if err := json.Unmarshal(resBodyRaw, &res); err != nil {
		return "", err
	}

	return res.ChannelID, nil
}

func (c *Client) send(channelID, content string) error {
	reqBody, err := json.Marshal(PostDMRequestJSON{
		Content: content,
	})
	if err != nil {
		return err
	}

	url, err := url.JoinPath("https://discordapp.com/api/channels/", channelID, "messages")
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return err
	}

	req.Header.Set("authorization", "Bot "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resBodyRaw, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBodyRaw))
	if err != nil {
		return err
	}

	return nil
}
