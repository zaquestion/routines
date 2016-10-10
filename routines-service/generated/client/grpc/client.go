// Package grpc provides a gRPC client for the add service.
package grpc

import (
	//"time"

	//jujuratelimit "github.com/juju/ratelimit"
	//stdopentracing "github.com/opentracing/opentracing-go"
	//"github.com/sony/gobreaker"
	"google.golang.org/grpc"

	//"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	//"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/ratelimit"
	//"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/zaquestion/routines/routines-service"
	svc "github.com/zaquestion/routines/routines-service/generated"
	handler "github.com/zaquestion/routines/routines-service/handlers/server"
)

// New returns an AddService backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn /*, tracer stdopentracing.Tracer, logger log.Logger*/) handler.Service {
	// We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.

	//limiter := ratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(100, 100))

	var scraptrelloresetEndpoint endpoint.Endpoint
	{
		scraptrelloresetEndpoint = grpctransport.NewClient(
			conn,
			"routines.RoutinesService",
			"ScrapTrelloReset",
			svc.EncodeGRPCScrapTrelloResetRequest,
			svc.DecodeGRPCScrapTrelloResetResponse,
			pb.ScrapTrelloResetReply{},
			//grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "ScrapTrelloReset", logger)),
		).Endpoint()
		//scraptrelloresetEndpoint = opentracing.TraceClient(tracer, "ScrapTrelloReset")(scraptrelloresetEndpoint)
		//scraptrelloresetEndpoint = limiter(scraptrelloresetEndpoint)
		//scraptrelloresetEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		//Name:    "ScrapTrelloReset",
		//Timeout: 30 * time.Second,
		//}))(scraptrelloresetEndpoint)
	}

	var getroutinesEndpoint endpoint.Endpoint
	{
		getroutinesEndpoint = grpctransport.NewClient(
			conn,
			"routines.RoutinesService",
			"GetRoutines",
			svc.EncodeGRPCGetRoutinesRequest,
			svc.DecodeGRPCGetRoutinesResponse,
			pb.GetRoutinesReply{},
			//grpctransport.ClientBefore(opentracing.FromGRPCRequest(tracer, "GetRoutines", logger)),
		).Endpoint()
		//getroutinesEndpoint = opentracing.TraceClient(tracer, "GetRoutines")(getroutinesEndpoint)
		//getroutinesEndpoint = limiter(getroutinesEndpoint)
		//getroutinesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		//Name:    "GetRoutines",
		//Timeout: 30 * time.Second,
		//}))(getroutinesEndpoint)
	}

	return svc.Endpoints{

		ScrapTrelloResetEndpoint: scraptrelloresetEndpoint,
		GetRoutinesEndpoint:      getroutinesEndpoint,
	}
}
