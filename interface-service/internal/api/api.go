package api

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/internal/models"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/interface-service/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// New returns an implementation of proto InterfaceServer
func New(client pb.DataClient) pb.InterfaceServer {
	return &implementation{
		dataService: client,
	}
}

type implementation struct {
	dataService pb.DataClient
	pb.UnsafeInterfaceServer
}

func (i *implementation) ReminderList(ctx context.Context, in *pb.ReminderListRequest) (*pb.ReminderListResponse, error) {
	response, err := i.dataService.ReminderList(ctx, &pb.ReminderListRequest{})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return response, nil
}

func (i *implementation) ReminderGet(ctx context.Context, in *pb.ReminderGetRequest) (*pb.ReminderGetResponse, error) {
	response, err := i.dataService.ReminderGet(ctx, &pb.ReminderGetRequest{
		Id: in.GetId(),
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return response, nil
}

func (i *implementation) ReminderCreate(ctx context.Context, in *pb.ReminderCreateRequest) (*pb.ReminderCreateResponse, error) {
	reminder := models.Reminder{
		Text: in.GetText(),
		Date: utils.TimestampToTime(in.GetDate()),
	}
	if err := reminder.Valid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response, err := i.dataService.ReminderCreate(ctx, &pb.ReminderCreateRequest{
		Date: in.GetDate(),
		Text: in.GetText(),
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return response, nil
}

func (i *implementation) ReminderUpdate(ctx context.Context, in *pb.ReminderUpdateRequest) (*pb.ReminderUpdateResponse, error) {
	reminder := models.Reminder{
		Text: in.GetText(),
		Date: time.Now(),
	}
	if err := reminder.Valid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	response, err := i.dataService.ReminderUpdate(ctx, &pb.ReminderUpdateRequest{
		Text: in.GetText(),
		Id:   in.GetId(),
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return response, nil
}

func (i *implementation) ReminderRemove(ctx context.Context, in *pb.ReminderRemoveRequest) (*pb.ReminderRemoveResponse, error) {
	response, err := i.dataService.ReminderRemove(ctx, &pb.ReminderRemoveRequest{
		Id: in.GetId(),
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return response, nil
}
