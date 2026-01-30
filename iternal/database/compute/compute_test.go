package compute

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewCompute(t *testing.T) {

	tests := map[string]struct {
		logger *zap.Logger

		expectedErr    error
		expectedNilObj bool
	}{
		"create compute without logger": {
			expectedErr:    errors.New("logger is invalid"),
			expectedNilObj: true,
		},
		"create compute": {
			logger:      zap.NewNop(),
			expectedErr: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {

			compute, err := NewCompute(test.logger)
			assert.Equal(t, test.expectedErr, err)
			if test.expectedNilObj {
				assert.Nil(t, compute)
			} else {
				assert.NotNil(t, compute)
			}
		})
	}
}
