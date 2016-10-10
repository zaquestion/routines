package handler

// This file contains the Service definition, and a basic service
// implementation. It also includes service middlewares.

import (
	_ "errors"
	_ "time"

	"golang.org/x/net/context"

	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/metrics"

	_ "github.com/zaquestion/routines/internal/routines"
	pb "github.com/zaquestion/routines/routines-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() Service {
	return routinesService{}
}

type routinesService struct {
}

// ScrapTrelloReset implements Service.
func (s routinesService) ScrapTrelloReset(ctx context.Context, inpb *pb.ScrapTrelloResetRequest) (*pb.ScrapTrelloResetReply, error) {
	_ = ctx
	response := pb.ScrapTrelloResetReply{}
	return &response, nil
}

// GetRoutines implements Service.
func (s routinesService) GetRoutines(ctx context.Context, in *pb.GetRoutinesRequest) (*pb.GetRoutinesReply, error) {
	_ = ctx
	_ = in
	response := pb.GetRoutinesReply{
	// Routines:
	// Err:
	}
	return &response, nil
}

// OauthCallback implements Service.

// Err:

// TrelloAuth implements Service.

type Service interface {
	ScrapTrelloReset(ctx context.Context, in *pb.ScrapTrelloResetRequest) (*pb.ScrapTrelloResetReply, error)
	GetRoutines(ctx context.Context, in *pb.GetRoutinesRequest) (*pb.GetRoutinesReply, error)
}
