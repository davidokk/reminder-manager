package api

import (
	"context"

	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New() pb.AdminServer {
	return &implementation{}
}

type implementation struct {
	pb.UnsafeAdminServer
}

func (i *implementation) ReminderList(context.Context, *pb.ReminderListRequest) (*pb.ReminderListResponse, error) {
	reminders, err := storage.Data()
	if err != nil {
		if err == storage.ErrorTimeoutExceeded {
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
	if err := storage.Add(storage.NewReminder(in.GetText(), utils.TimestampToTime(in.GetDate()))); err != nil {
		if err == storage.ErrorIDAlreadyExists {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderCreateResponse{}, nil
}

func (i *implementation) ReminderUpdate(ctx context.Context, in *pb.ReminderUpdateRequest) (*pb.ReminderUpdateResponse, error) {
	if err := storage.Edit(in.GetId(), in.GetText()); err != nil {
		if err == storage.ErrorIDNotExists {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderUpdateResponse{}, nil
}

func (i *implementation) ReminderRemove(ctx context.Context, in *pb.ReminderRemoveRequest) (*pb.ReminderRemoveResponse, error) {
	if err := storage.RemoveById(in.GetId()); err != nil {
		if err == storage.ErrorIDNotExists {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.ReminderRemoveResponse{}, nil
}
