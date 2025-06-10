package payments

import (
	"sumup/notifications/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type paymentServiceDependencies struct {
	userRepository *mocks.MockUserRepository
	producer       *mocks.MockProducer
}

func newPaymentServiceDependencies() *paymentServiceDependencies {
	return &paymentServiceDependencies{
		userRepository: &mocks.MockUserRepository{},
		producer:       &mocks.MockProducer{},
	}
}

func Test_ProcessPaymentNotification(t *testing.T) {
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
		mock func(deps *paymentServiceDependencies, args *args, want *want)
	}{
		// TODO implement tests when the notification sending logic is implemented
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deps := newPaymentServiceDependencies()
			service := NewPaymentService(
				deps.userRepository,
				deps.producer,
			)

			args, want := tt.args(), tt.want()

			tt.mock(deps, args, want)

			err := service.ProcessPaymentNotification(args.userID, args.amount)

			if want.err != nil {
				assert.Equal(t, want.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
