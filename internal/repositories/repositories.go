package repositories

import (
	"sumup/notifications/internal/entities"
)

//go:generate mockgen -source=repositories.go -destination=../mocks/repositories.go -package=mocks
type UserRepository interface {
	GetUserDataByID(id int) (*entities.User, error)
	GetUserNotificationTypesByUserID(id int) ([]entities.NotificationType, error)
}
