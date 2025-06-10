package users

import (
	"database/sql"
	"fmt"
	"os"
	"sumup/notifications/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func Test_GetUserDataByID(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	bytes, _ := os.ReadFile("../test_data.sql")
	sqlString := string(bytes)

	db.Exec(sqlString)

	type args struct {
		id int
	}
	type want struct {
		user *entities.User
		err  error
	}
	tests := []struct {
		name string
		args func() *args
		want func() *want
	}{
		{
			name: "Should return valid user data",
			args: func() *args {
				return &args{id: 1}
			},
			want: func() *want {
				return &want{
					user: &entities.User{
						ID:    1,
						Name:  "John Doe",
						Email: "johndoe@email.com",
						Phone: "123456789",
					},
					err: nil,
				}
			},
		},
		{
			name: "Should return error when user does not exist",
			args: func() *args {
				return &args{id: 3}
			},
			want: func() *want {
				return &want{
					user: nil,
					err:  fmt.Errorf("failed to query user data: sql: no rows in result set"),
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, want := tt.args(), tt.want()

			repo := NewUserRepository(db)
			result, err := repo.GetUserDataByID(args.id)

			assert.Equal(t, want.user, result)
			if want.err != nil {
				assert.Equal(t, want.err.Error(), err.Error())
			}
		})
	}
}

func Test_GetUserNotificationTypesByUserID(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	bytes, _ := os.ReadFile("test_data.sql")
	sqlString := string(bytes)

	db.Exec(sqlString)

	type args struct {
		id int
	}
	type want struct {
		notificationTypes []entities.NotificationType
		err               error
	}
	tests := []struct {
		name string
		args func() *args
		want func() *want
	}{
		{
			name: "Should return valid notification types and skip invalid ones",
			args: func() *args {
				return &args{id: 1}
			},
			want: func() *want {
				return &want{
					notificationTypes: []entities.NotificationType{
						entities.EmailNotification,
						entities.SMSNotification,
					},
					err: nil,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, want := tt.args(), tt.want()

			repo := NewUserRepository(db)
			result, err := repo.GetUserNotificationTypesByUserID(args.id)

			assert.Equal(t, want.notificationTypes, result)
			if want.err != nil {
				assert.Equal(t, want.err.Error(), err.Error())
			}
		})
	}
}
