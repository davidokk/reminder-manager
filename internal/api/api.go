package api

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// New returns an implementation of proto AdminServer
func New(storage storage.RemindersStorage) pb.AdminServer {
	return &implementation{
		storage: storage,
	}
}

type implementation struct {
	storage storage.RemindersStorage
	pb.UnsafeAdminServer
}

func (i *implementation) ReminderList(ctx context.Context, in *pb.ReminderListRequest) (*pb.ReminderListResponse, error) {
	reminders, err := i.storage.ListReminders(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := make([]*pb.Reminder, 0, len(reminders))
	for _, rem := range reminders {
		result = append(result, &pb.Reminder{
			Id:   rem.ID,
			Text: rem.Text,
			Date: utils.TimeToTimestamp(rem.Date),
		})
	}
	return &pb.ReminderListResponse{
		Reminders: result,
	}, nil
}

func (i *implementation) ReminderGet(ctx context.Context, in *pb.ReminderGetRequest) (*pb.ReminderGetResponse, error) {
	rem, err := i.storage.GetReminder(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderGetResponse{
		Reminder: &pb.Reminder{
			Id:   rem.ID,
			Text: rem.Text,
			Date: utils.TimeToTimestamp(rem.Date),
		},
	}, nil
}

func (i *implementation) ReminderCreate(ctx context.Context, in *pb.ReminderCreateRequest) (*pb.ReminderCreateResponse, error) {
	rem, err := i.storage.CreateReminder(ctx, utils.TimestampToTime(in.GetDate()), in.GetText())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderCreateResponse{
		Id: rem.ID,
	}, nil
}

func (i *implementation) ReminderUpdate(ctx context.Context, in *pb.ReminderUpdateRequest) (*pb.ReminderUpdateResponse, error) {
	if err := i.storage.UpdateReminder(ctx, in.GetId(), in.GetText()); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderUpdateResponse{}, nil
}

func (i *implementation) ReminderRemove(ctx context.Context, in *pb.ReminderRemoveRequest) (*pb.ReminderRemoveResponse, error) {
	if err := i.storage.RemoveReminder(ctx, in.GetId()); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderRemoveResponse{}, nil
}
