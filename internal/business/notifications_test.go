package business

import (
	"sumup/notifications/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type notificationServiceDependencies struct {
	userRepository *mocks.MockUserRepository
}

func newNotificationServiceDependencies() *notificationServiceDependencies {
	return &notificationServiceDependencies{
		userRepository: &mocks.MockUserRepository{},
	}
}

func Test_SendPaymentNotifications(t *testing.T) {
	type args struct {
		userID int
		amount float64
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args func() *args
		want func() *want
		mock func(deps *notificationServiceDependencies, args *args, want *want)
	}{
		// TODO implement tests when the notification sending logic is implemented
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := newNotificationServiceDependencies()
			service := NewNotificationService(deps.userRepository)

			args, want := tt.args(), tt.want()

			tt.mock(deps, args, want)

			err := service.SendPaymentNotifications(args.userID, args.amount)

			if want.err != nil {
				assert.Equal(t, want.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
