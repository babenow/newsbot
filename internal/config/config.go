package config

import (
	"log"
	"sync"
	"time"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfighcl"
)

type Config struct {
	TelegramBotToken     string        `hcl:"telegram_bot_token" env:"TELEGRAM_BOT_TOKEN" required:"true"`
	TelegramChannelID    int64         `hcl:"telegram_channel_id" env:"TELEGRAM_CHANNEL_ID" required:"true"`
	DatabaseDNS          string        `hcl:"database_dns" env:"DATABASEW_DNS" default:"postgres://postgres:postgres@localhost:5432/news_feed_bot?sslmode=disable"`
	FetchInterval        time.Duration `hcl:"fetch_interval" env:"FETCH_INTERVAL" default:"10m"`
	NotificationInterval time.Duration `hcl:"notification_interval" env:"NOTIFICATION_INTERVAL" default:"1m"`
	FilterKeywords       []string      `hcl:"filter_keywords" env:"FILTER_KEYWORDS"`
	OpenAIKey            string        `hcl:"openai_key" env:"OPENAI_KEY"`
	OpenAIPrompt         string        `hcl:"openai_propt" env:"OPENAI_PROMPT"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		loader := aconfig.LoaderFor(&cfg, aconfig.Config{
			EnvPrefix: "NFB",
			Files:     []string{"./config.hcl", "./config.local.hcl"},
			FileDecoders: map[string]aconfig.FileDecoder{
				".hcl": aconfighcl.New(),
			},
		})
		if err := loader.Load(); err != nil {
			log.Printf("[ERROR] Failed to load config: %v", err)
		}
	})

	return cfg
}
