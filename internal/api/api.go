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
func New() pb.AdminServer {
	return &implementation{}
}

type implementation struct {
	pb.UnsafeAdminServer
}

func (i *implementation) ReminderList(ctx context.Context, in *pb.ReminderListRequest) (*pb.ReminderListResponse, error) {
	reminders, err := storage.Data(ctx)
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

func (i *implementation) ReminderCreate(ctx context.Context, in *pb.ReminderCreateRequest) (*pb.ReminderCreateResponse, error) {
	if err := storage.Add(ctx, storage.NewReminder(in.GetText(), utils.TimestampToTime(in.GetDate()))); err != nil {
		if errors.Is(err, storage.ErrorIDAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		} else if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderCreateResponse{}, nil
}

func (i *implementation) ReminderUpdate(ctx context.Context, in *pb.ReminderUpdateRequest) (*pb.ReminderUpdateResponse, error) {
	if err := storage.Edit(ctx, in.GetId(), in.GetText()); err != nil {
		if errors.Is(err, storage.ErrorIDNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderUpdateResponse{}, nil
}

func (i *implementation) ReminderRemove(ctx context.Context, in *pb.ReminderRemoveRequest) (*pb.ReminderRemoveResponse, error) {
	if err := storage.RemoveByID(ctx, in.GetId()); err != nil {
		if errors.Is(err, storage.ErrorIDNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderRemoveResponse{}, nil
}
