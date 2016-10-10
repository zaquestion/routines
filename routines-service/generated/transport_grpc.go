package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	//stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/zaquestion/routines/routines-service"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AddServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints /*, tracer stdopentracing.Tracer, logger log.Logger*/) pb.RoutinesServiceServer {
	//options := []grpctransport.ServerOption{
	//grpctransport.ServerErrorLogger(logger),
	//}
	return &grpcServer{
		// routinesservice

		scraptrelloreset: grpctransport.NewServer(
			ctx,
			endpoints.ScrapTrelloResetEndpoint,
			DecodeGRPCScrapTrelloResetRequest,
			EncodeGRPCScrapTrelloResetResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "ScrapTrelloReset", logger)))...,
		),
		getroutines: grpctransport.NewServer(
			ctx,
			endpoints.GetRoutinesEndpoint,
			DecodeGRPCGetRoutinesRequest,
			EncodeGRPCGetRoutinesResponse,
			//append(options,grpctransport.ServerBefore(opentracing.FromGRPCRequest(tracer, "GetRoutines", logger)))...,
		),
	}
}

type grpcServer struct {
	scraptrelloreset grpctransport.Handler
	getroutines      grpctransport.Handler
}

// Methods

func (s *grpcServer) ScrapTrelloReset(ctx context.Context, req *pb.ScrapTrelloResetRequest) (*pb.ScrapTrelloResetReply, error) {
	_, rep, err := s.scraptrelloreset.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ScrapTrelloResetReply), nil
}

func (s *grpcServer) GetRoutines(ctx context.Context, req *pb.GetRoutinesRequest) (*pb.GetRoutinesReply, error) {
	_, rep, err := s.getroutines.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetRoutinesReply), nil
}

// Server Decode

// DecodeGRPCScrapTrelloResetRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC scraptrelloreset request to a user-domain scraptrelloreset request. Primarily useful in a server.
func DecodeGRPCScrapTrelloResetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ScrapTrelloResetRequest)
	return req, nil
}

// DecodeGRPCGetRoutinesRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getroutines request to a user-domain getroutines request. Primarily useful in a server.
func DecodeGRPCGetRoutinesRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetRoutinesRequest)
	return req, nil
}

// Client Decode

// DecodeGRPCScrapTrelloResetResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC scraptrelloreset reply to a user-domain scraptrelloreset response. Primarily useful in a client.
func DecodeGRPCScrapTrelloResetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ScrapTrelloResetReply)
	return reply, nil
}

// DecodeGRPCGetRoutinesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getroutines reply to a user-domain getroutines response. Primarily useful in a client.
func DecodeGRPCGetRoutinesResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetRoutinesReply)
	return reply, nil
}

// Server Encode

// EncodeGRPCScrapTrelloResetResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain scraptrelloreset response to a gRPC scraptrelloreset reply. Primarily useful in a server.
func EncodeGRPCScrapTrelloResetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ScrapTrelloResetReply)
	return resp, nil
}

// EncodeGRPCGetRoutinesResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getroutines response to a gRPC getroutines reply. Primarily useful in a server.
func EncodeGRPCGetRoutinesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetRoutinesReply)
	return resp, nil
}

// Client Encode

// EncodeGRPCScrapTrelloResetRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain scraptrelloreset request to a gRPC scraptrelloreset request. Primarily useful in a client.
func EncodeGRPCScrapTrelloResetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ScrapTrelloResetRequest)
	return req, nil
}

// EncodeGRPCGetRoutinesRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getroutines request to a gRPC getroutines request. Primarily useful in a client.
func EncodeGRPCGetRoutinesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetRoutinesRequest)
	return req, nil
}
