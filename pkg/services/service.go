package services

import "database/sql"

type Service struct {
	UserService    UserService
	PostService    PostService
	SessionService SessionService
}

func NewService(db *sql.DB) *Service {
	return &Service{
		UserService:    *NewUserService(db),
		PostService:    *NewPostService(db),
		SessionService: *NewSessionService(db),
	}
}
