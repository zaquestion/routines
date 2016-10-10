package clienthandler

import (
	pb "github.com/zaquestion/routines/routines-service"
)

// ScrapTrelloReset implements Service.
func ScrapTrelloReset() (*pb.ScrapTrelloResetRequest, error) {

	request := pb.ScrapTrelloResetRequest{}
	return &request, nil
}

// GetRoutines implements Service.
func GetRoutines(DateStartGetRoutines string, DateEndGetRoutines string) (*pb.GetRoutinesRequest, error) {

	request := pb.GetRoutinesRequest{
		DateStart: DateStartGetRoutines,
		DateEnd:   DateEndGetRoutines,
	}
	return &request, nil
}
