package twitch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var scopes = []string{"openid"}

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	CallbackURL  string
	Secret       string
}

type Client struct {
	Config
}

func New(cfg Config) *Client {
	return &Client{
		Config: cfg,
	}
}

func (c *Client) GetSignedUrl() string {
	return fmt.Sprintf(
		`https://id.twitch.tv/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=token&scope=%s`,
		c.ClientID,
		c.RedirectURL,
		url.QueryEscape(strings.Join(scopes, " ")),
	)
}

func (c *Client) GetCurrentUser(ctx context.Context, accessToken string) (string, error) {
	req, err := http.NewRequest("GET", "https://id.twitch.tv/oauth2/userinfo", nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New("invalid access token")
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var sessionResp struct {
		UserID   string `json:"sub"`
		UserName string `json:"preferred_username"`
	}
	if decodeErr := json.Unmarshal(resBytes, &sessionResp); decodeErr != nil {
		return "", decodeErr
	}

	return sessionResp.UserName, nil
}

func (c *Client) SubscribeToUserFollows(ctx context.Context, streamerName string) error {
	client := http.Client{}

	reqData := struct {
		CallBack  string `json:"hub.callback"`
		Mode      string `json:"hub.mode"`
		Topic     string `json:"hub.topic"`
		ExpireSec int    `json:"hub.lease_seconds"`
		Secret    string `json:"hub.secret"`
	}{
		CallBack:  c.CallbackURL,
		Mode:      "subscribe",
		Topic:     "https://api.twitch.tv/helix/users/follows?first=1&to_id=" + streamerName,
		ExpireSec: 0,
		Secret:    c.Secret,
	}

	wr := bytes.NewBuffer(nil)
	if err := json.NewEncoder(wr).Encode(&reqData); err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.twitch.tv/helix/webhooks/hub",
		wr,
	)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusAccepted {
		return errors.New("not good response from twitch")
	}

	return nil
}
