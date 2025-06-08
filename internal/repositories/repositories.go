package repositories

import (
	"sumup/notifications/internal/entities"
)

type UserRepository interface {
	GetUserDataByID(id int) (*entities.User, error)
	GetUserNotificationTypesByUserID(id int) ([]entities.NotificationType, error)
}
