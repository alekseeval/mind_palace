package tgbot

import (
	"MindPalace/internal/tgbot/configuration"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
)

type TelegramBot struct {
	logger           *logrus.Entry
	telegramApiToken string
	cancelFunction   context.CancelFunc
	wg               *sync.WaitGroup
	exitChannel      chan interface{}
}

func NewTelegramBot(config *configuration.Config, logger *logrus.Entry) *TelegramBot {
	exitChan := make(chan interface{}, 1)
	return &TelegramBot{
		logger:           logger,
		telegramApiToken: config.System.Telegram.AccessToken,
		exitChannel:      exitChan,
	}
}

func (bot *TelegramBot) Run() error {
	bot.wg = &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	bot.cancelFunction = cancel
	select {
	case <-ctx.Done():
		bot.wg.Wait()
		bot.exitChannel <- nil
		return nil
	}

	//return nil
}

func (bot *TelegramBot) Shutdown(ctx context.Context) error {
	if bot.wg == nil || bot.cancelFunction == nil {
		return fmt.Errorf("bot was not started yet")
	}
	bot.cancelFunction()
	bot.logger.Info("Stopping bot gracefully..")
	select {
	case <-ctx.Done():
		return fmt.Errorf("context expired while wait to stop all bot processes")
	case <-bot.exitChannel:
		bot.logger.Info("Bot was stopped successfully")
		bot.cancelFunction = nil
		bot.wg = nil
		return nil
	}
}
