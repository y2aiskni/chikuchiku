package usecase

import (
	"fmt"
	"time"

	"github.com/y2aiskni/chikuchiku/internal/adapter/discord"
	"github.com/y2aiskni/chikuchiku/internal/entity"
	"gorm.io/gorm"
)

type Chikuchiku interface {
	PostTodayToDiscord(postURL string) error
}

type chikuchiku struct {
	db *gorm.DB
}

func NewChikuchiku(db *gorm.DB) Chikuchiku {
	return &chikuchiku{
		db: db,
	}
}

func (c *chikuchiku) PostTodayToDiscord(postURL string) error {
	now := time.Now()

	var chiku entity.Chikuchiku
	if err := c.db.Where("date = ?", now.Format("2006-01-02")).First(&chiku).Error; err != nil {
		return fmt.Errorf("failed to c.db.First(): %w", err)
	}

	if err := discord.PostWebhook(postURL, discord.WebHookBody{
		Embeds: []*discord.Embed{
			{
				Title:       "今日のちくちく",
				Description: chiku.Message,
				URL:         chiku.URL,
				Footer: discord.EmbedFooter{
					Text: chiku.Date.Format("2006年01月02日"),
				},
			},
		},
	}); err != nil {
		return fmt.Errorf("failed to discord.PostWebhook(): %w", err)
	}

	return nil
}
