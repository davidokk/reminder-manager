package api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mockstorage "gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage/mocks"
)

type serviceFixture struct {
	service *implementation
	storage *mockstorage.MockRemindersStorage
	ctx     context.Context
}

func setUp(t *testing.T) serviceFixture {
	var f serviceFixture
	f.ctx = context.Background()
	f.storage = mockstorage.NewMockRemindersStorage(gomock.NewController(t))
	f.service = New(f.storage)

	return f
}
