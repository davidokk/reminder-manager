package api

import (
	"context"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/pkg/api"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// New returns an implementation of proto DataServer
func New(storage storage.RemindersStorage) *implementation {
	return &implementation{
		storage: storage,
	}
}

type implementation struct {
	storage storage.RemindersStorage
	api.UnsafeDataServer
}

func (i *implementation) ReminderList(ctx context.Context, in *api.ReminderListRequest) (*api.ReminderListResponse, error) {
	reminders, err := i.storage.ListReminders(ctx)
	if err != nil {
		grpcErrors.Inc(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*api.Reminder, 0, len(reminders))
	for _, rem := range reminders {
		result = append(result, &api.Reminder{
			Id:   rem.ID,
			Text: rem.Text,
			Date: utils.TimeToTimestamp(rem.Date),
		})
	}
	successGRPcRequests.Inc()
	return &api.ReminderListResponse{
		Reminders: result,
	}, nil
}

func (i *implementation) ReminderGet(ctx context.Context, in *api.ReminderGetRequest) (*api.ReminderGetResponse, error) {
	rem, err := i.storage.GetReminder(ctx, in.GetId())
	if err != nil {
		grpcErrors.Inc(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	successGRPcRequests.Inc()
	return &api.ReminderGetResponse{
		Reminder: &api.Reminder{
			Id:   rem.ID,
			Text: rem.Text,
			Date: utils.TimeToTimestamp(rem.Date),
		},
	}, nil
}

func (i *implementation) ReminderCreate(ctx context.Context, in *api.ReminderCreateRequest) (*api.ReminderCreateResponse, error) {
	rem, err := i.storage.CreateReminder(ctx, utils.TimestampToTime(in.GetDate()), in.GetText())
	if err != nil {
		grpcErrors.Inc(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	successGRPcRequests.Inc()
	return &api.ReminderCreateResponse{
		Id: rem.ID,
	}, nil
}

func (i *implementation) ReminderUpdate(ctx context.Context, in *api.ReminderUpdateRequest) (*api.ReminderUpdateResponse, error) {
	if err := i.storage.UpdateReminder(ctx, in.GetId(), in.GetText()); err != nil {
		grpcErrors.Inc(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	successGRPcRequests.Inc()
	return &api.ReminderUpdateResponse{}, nil
}

func (i *implementation) ReminderRemove(ctx context.Context, in *api.ReminderRemoveRequest) (*api.ReminderRemoveResponse, error) {
	if err := i.storage.RemoveReminder(ctx, in.GetId()); err != nil {
		grpcErrors.Inc(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	successGRPcRequests.Inc()
	return &api.ReminderRemoveResponse{}, nil
}
