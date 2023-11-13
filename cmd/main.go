package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/babenow/newsbot/internal/botkit"
	"github.com/babenow/newsbot/internal/botkit/bot"
	"github.com/babenow/newsbot/internal/config"
	"github.com/babenow/newsbot/internal/fetcher"
	"github.com/babenow/newsbot/internal/notifier"
	"github.com/babenow/newsbot/internal/storage"
	"github.com/babenow/newsbot/internal/summary"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

func main() {
	botApi, err := tgbotapi.NewBotAPI(config.Get().TelegramBotToken)
	if err != nil {
		log.Printf("[ERROR] failed to create telegam bot: %v", err)
		return
	}

	db, err := sqlx.Connect("postgres", config.Get().DatabaseDNS)
	if err != nil {
		log.Printf("[ERROR] failed to connect database: %v", err)
		return
	}
	defer db.Close()
	var (
		aStorage = storage.NewArticlePostgresStorage(db)
		sStorage = storage.NewSourcePostgresStorage(db)
		fetcher  = fetcher.NewFetcher(
			aStorage,
			sStorage,
			config.Get().FetchInterval,
			config.Get().FilterKeywords,
		)
		notifier = notifier.NewNotifier(
			aStorage,
			summary.NewOpenAISummarizer(config.Get().OpenAIKey, config.Get().OpenAIPrompt),
			botApi,
			config.Get().NotificationInterval,
			2*config.Get().FetchInterval,
			// 24*10*time.Hour,
			config.Get().TelegramChannelID,
		)
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	newsBot := botkit.NewBot(botApi)

	newsBot.RegisterCmdView("start", bot.ViewCmdStart())
	newsBot.RegisterCmdView("addsource", bot.ViewCmdAddSource(sStorage))

	go func(ctx context.Context) {
		if err := fetcher.Start(ctx); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Printf("[ERROR] failed to start fetcher: %v", err)
				return
			}

			log.Printf("[ERROR] fetcher stopped with error: %v", err)
		}
	}(ctx)

	go func(ctx context.Context) {
		if err := notifier.Start(ctx); err != nil {
			if !errors.Is(err, context.Canceled) {
				log.Printf("[ERROR] failed to select and send articles: %v", err)
				return
			}

			log.Printf("[ERROR] notifier stopped with error^ %v", err)
		}
	}(ctx)

	if err := newsBot.Run(ctx); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("[ERROR] failed start telegram bot: %v", err)
			return
		}

		log.Printf("[ERROR] bot stopped with error: %v", err)
	}
}
