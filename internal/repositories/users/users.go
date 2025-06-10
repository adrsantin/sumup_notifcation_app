package users

import (
	"database/sql"
	"fmt"
	"sumup/notifications/internal/entities"
	"sumup/notifications/internal/repositories"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

const (
	userNotificationsListQuery = "SELECT n.type FROM user_notification un JOIN notification n ON un.notification_id = n.id WHERE un.user_id = ?"
	userDataQuery              = "SELECT id, name, email, phone FROM user WHERE id = ?"
)

func (r *userRepositoryImpl) GetUserDataByID(id int) (*entities.User, error) {
	user := &entities.User{}
	row := r.db.QueryRow(userDataQuery, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, fmt.Errorf("failed to query user data: %w", err)
	}
	return user, nil
}

func (r *userRepositoryImpl) GetUserNotificationTypesByUserID(id int) ([]entities.NotificationType, error) {
	notificationTypes := make([]entities.NotificationType, 0)
	rows, err := r.db.Query(userNotificationsListQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query user notification types: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var notificationType string
		rows.Scan(&notificationType)
		if entities.IsValidNotificationType(notificationType) {
			notificationTypes = append(notificationTypes, entities.NotificationType(notificationType))
		}
	}
	return notificationTypes, nil
}
