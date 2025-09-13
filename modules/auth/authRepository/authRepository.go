package authRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	AuthRepositoryService interface {
	}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}

func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}
