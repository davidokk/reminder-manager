package api

import (
	"context"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/internal/producer"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/internal/models"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/interface-service/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"encoding/json"
)

// New returns an implementation of proto InterfaceServer
func New(client pb.DataClient) pb.InterfaceServer {
	return &implementation{
		dataService: client,
		producer:    producer.New(),
	}
}

type implementation struct {
	dataService pb.DataClient
	producer    sarama.AsyncProducer

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

	params := make(map[string]string)
	params["date"] = utils.TimestampToTime(in.GetDate()).Format(time.RFC3339)
	params["text"] = in.GetText()
	b, _ := json.Marshal(params)

	i.producer.Input() <- &sarama.ProducerMessage{
		Topic: config.App.Kafka.DataIncomingTopic,
		Key:   sarama.StringEncoder("create"),
		Value: sarama.StringEncoder(b),
	}

	return &pb.ReminderCreateResponse{}, nil
}

func (i *implementation) ReminderUpdate(ctx context.Context, in *pb.ReminderUpdateRequest) (*pb.ReminderUpdateResponse, error) {
	reminder := models.Reminder{
		Text: in.GetText(),
		Date: time.Now(),
	}
	if err := reminder.Valid(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	params := make(map[string]string)
	params["id"] = strconv.FormatInt(int64(in.GetId()), 10)
	params["text"] = in.GetText()
	b, _ := json.Marshal(params)

	i.producer.Input() <- &sarama.ProducerMessage{
		Topic: config.App.Kafka.DataIncomingTopic,
		Key:   sarama.StringEncoder("update"),
		Value: sarama.StringEncoder(b),
	}

	return &pb.ReminderUpdateResponse{}, nil
}

func (i *implementation) ReminderRemove(ctx context.Context, in *pb.ReminderRemoveRequest) (*pb.ReminderRemoveResponse, error) {
	params := make(map[string]string)
	params["id"] = strconv.FormatInt(int64(in.GetId()), 10)
	b, _ := json.Marshal(params)

	i.producer.Input() <- &sarama.ProducerMessage{
		Topic: config.App.Kafka.DataIncomingTopic,
		Key:   sarama.StringEncoder("remove"),
		Value: sarama.StringEncoder(b),
	}

	return &pb.ReminderRemoveResponse{}, nil
}
