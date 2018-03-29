// Code generated by truss.
// Rerunning truss will overwrite this file.
// DO NOT EDIT!

package handlers

import (
	pb "github.com/zaquestion/routines"
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