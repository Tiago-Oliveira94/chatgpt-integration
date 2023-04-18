package gateway

import (
	"context"

	"github.com/Tiago-Oliveira94/chatgpt-integration/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatByID(ctx context.Context, chatId string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
