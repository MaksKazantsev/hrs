package db

import "context"

type Repository interface {
	CreateMessage(ctx context.Context)
	DeleteMessage(ctx context.Context)
	EditMessage(ctx context.Context)
}
