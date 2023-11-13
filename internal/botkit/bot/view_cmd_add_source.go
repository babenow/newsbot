package bot

import (
	"context"
	"fmt"
	"github.com/babenow/newsbot/internal/botkit/markup"

	"github.com/babenow/newsbot/internal/botkit"
	"github.com/babenow/newsbot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SourceStorage interface {
	Add(ctx context.Context, source model.Source) (int64, error)
}

func ViewCmdAddSource(ss SourceStorage) botkit.ViewFunc {
	type addSourceArgs struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		args, err := botkit.ParseJSON[addSourceArgs](update.Message.CommandArguments())
		if err != nil {
			// TODO: Send user message
			return err
		}

		source := model.Source{
			Name:    args.Name,
			FeedURL: args.URL,
		}

		sourceID, err := ss.Add(ctx, source)
		if err != nil {
			// TODO: Send user message
			return err
		}

		var (
			msgText = fmt.Sprintf(
				"Источник *%s* добавлен c ID: `%d`\\. Этот ID можно использовать для управления источником\\.",
				markup.EscapeForMarkdown(args.URL),
				sourceID,
			)
			reply = tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		)

		reply.ParseMode = "MarkdownV2"

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}
