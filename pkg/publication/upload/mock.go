package upload

import (
	"context"
	"github.com/stretchr/testify/mock"
	"io"
)

type UploaderMock struct {
	mock.Mock
}

func (um *UploaderMock) Upload(ctx context.Context, file io.Reader, size int64, contentType string) (string, error) {
	args := um.Called(file, size, contentType)
	return args.Get(0).(string), args.Error(1)
}
