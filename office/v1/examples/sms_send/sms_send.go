package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/grokify/oauth2more/credentials"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	rc "github.com/grokify/go-ringcentral-client/office/v1/client"
	ru "github.com/grokify/go-ringcentral-client/office/v1/util"
)

type CliOptions struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
}

type RingCentralConfig struct {
	ServerURL    string `env:"RINGCENTRAL_SERVER_URL"`
	ClientID     string `env:"RINGCENTRAL_CLIENT_ID"`
	ClientSecret string `env:"RINGCENTRAL_CLIENT_SECRET"`
	Username     string `env:"RINGCENTRAL_USERNAME"`
	Extension    string `env:"RINGCENTRAL_EXTENSION"`
	Password     string `env:"RINGCENTRAL_PASSWORD"`
}

func NewRingCentralConfigEnv() (*RingCentralConfig, error) {
	appCfg := &RingCentralConfig{}
	return appCfg, env.Parse(appCfg)
}

func (cfg *RingCentralConfig) ApplicationCredentials() credentials.ApplicationCredentials {
	return credentials.ApplicationCredentials{
		ServerURL:    cfg.ServerURL,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret}
}

func (cfg *RingCentralConfig) PasswordCredentials() credentials.PasswordCredentials {
	return credentials.PasswordCredentials{
		Username: cfg.Username,
		Password: cfg.Password}
}

// example: $ go run fax_send.go -to=+16505550100 -file=$GOPATH/src/github.com/grokify/go-ringcentral-client/office/v1/examples/fax_send/test_file.pdf
func main() {
	opts := CliOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	err = config.LoadDotEnvFirst(opts.EnvFile, os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}
	rcCfg, err := NewRingCentralConfigEnv()
	if err != nil {
		log.Fatal(err)
	}

	apiClient, err := ru.NewApiClientPassword(
		rcCfg.ApplicationCredentials(),
		rcCfg.PasswordCredentials())
	if err != nil {
		log.Fatal(err)
	}

	params := rc.CreateSmsMessage{
		From: rc.MessageStoreCallerInfoRequest{
			PhoneNumber: os.Getenv("RINGCENTRAL_DEMO_SMS_FROM")},
		To: []rc.MessageStoreCallerInfoRequest{
			{
				PhoneNumber: os.Getenv("RINGCENTRAL_DEMO_SMS_TO")}},
		Text: os.Getenv("RINGCENTRAL_DEMO_SMS_TEXT")}

	fmtutil.PrintJSON(params)

	info, resp, err := apiClient.MessagesApi.SendSMS(context.Background(), "~", "~", params)

	if err != nil {
		panic(err)
	} else if resp.StatusCode > 299 {
		panic(fmt.Errorf("API Status %v", resp.StatusCode))
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
