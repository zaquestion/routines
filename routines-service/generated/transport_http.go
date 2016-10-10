package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	//stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	//"github.com/go-kit/kit/endpoint"
	//"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"

	// This service
	pb "github.com/zaquestion/routines/routines-service"
)

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.RegisterRoutinesServiceServer
	_ = io.Copy
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(ctx context.Context, endpoints Endpoints, logger log.Logger) http.Handler {
	//func MakeHTTPHandler(ctx context.Context, endpoints Endpoints, /*tracer stdopentracing.Tracer,*/ logger log.Logger) http.Handler {
	/*options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}*/
	m := http.NewServeMux()

	m.Handle("/scrap/trello", httptransport.NewServer(
		ctx,
		endpoints.ScrapTrelloResetEndpoint,
		HttpDecodeLogger(DecodeHTTPScrapTrelloResetZeroRequest, logger),
		EncodeHTTPGenericResponse,
	))

	m.Handle("/routines", httptransport.NewServer(
		ctx,
		endpoints.GetRoutinesEndpoint,
		HttpDecodeLogger(DecodeHTTPGetRoutinesZeroRequest, logger),
		EncodeHTTPGenericResponse,
	))
	return m
}

func HttpDecodeLogger(next httptransport.DecodeRequestFunc, logger log.Logger) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		logger.Log("method", r.Method, "url", r.URL.String())
		rv, err := next(ctx, r)
		if err != nil {
			logger.Log("method", r.Method, "url", r.URL.String(), "Error", err)
		}
		return rv, err
	}
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	msg := err.Error()

	/*if e, ok := err.(httptransport.Error); ok {
		msg = e.Err.Error()
		switch e.Domain {
		case httptransport.DomainDecode:
			code = http.StatusBadRequest

		case httptransport.DomainDo:
			switch e.Err {
			case ErrTwoZeroes, ErrMaxSizeExceeded, ErrIntOverflow:
				code = http.StatusBadRequest
			}
		}
	}*/

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorWrapper{Error: msg})
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// Server Decode

// DecodeHTTPScrapTrelloResetZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded scraptrelloreset request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPScrapTrelloResetZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.ScrapTrelloResetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	// err = io.EOF if r.Body was empty
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "decoding body of http request")
	}

	pathParams, err := PathParams(r.URL.Path, "/scrap/trello")
	_ = pathParams
	if err != nil {
		fmt.Printf("Error while reading path params: %v\n", err)
		return nil, errors.Wrap(err, "couldn't unmarshal path parameters")
	}
	queryParams, err := QueryParams(r.URL.Query())
	_ = queryParams
	if err != nil {
		fmt.Printf("Error while reading query params: %v\n", err)
		return nil, errors.Wrapf(err, "Error while reading query params: %v", r.URL.Query())
	}

	return &req, err
}

// DecodeHTTPGetRoutinesZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded getroutines request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPGetRoutinesZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.GetRoutinesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	// err = io.EOF if r.Body was empty
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "decoding body of http request")
	}

	pathParams, err := PathParams(r.URL.Path, "/routines")
	_ = pathParams
	if err != nil {
		fmt.Printf("Error while reading path params: %v\n", err)
		return nil, errors.Wrap(err, "couldn't unmarshal path parameters")
	}
	queryParams, err := QueryParams(r.URL.Query())
	_ = queryParams
	if err != nil {
		fmt.Printf("Error while reading query params: %v\n", err)
		return nil, errors.Wrapf(err, "Error while reading query params: %v", r.URL.Query())
	}

	DateStartGetRoutinesStr := queryParams["date_start"]
	DateStartGetRoutines := DateStartGetRoutinesStr
	// TODO: Better error handling
	if err != nil {
		fmt.Printf("Error while extracting DateStartGetRoutines from query: %v\n", err)
		fmt.Printf("queryParams: %v\n", queryParams)
		return nil, err
	}
	req.DateStart = DateStartGetRoutines

	DateEndGetRoutinesStr := queryParams["date_end"]
	DateEndGetRoutines := DateEndGetRoutinesStr
	// TODO: Better error handling
	if err != nil {
		fmt.Printf("Error while extracting DateEndGetRoutines from query: %v\n", err)
		fmt.Printf("queryParams: %v\n", queryParams)
		return nil, err
	}
	req.DateEnd = DateEndGetRoutines

	return &req, err
}

// Client Decode

// DecodeHTTPScrapTrelloReset is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded ScrapTrelloResetReply response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPScrapTrelloResetResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp pb.ScrapTrelloResetReply
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

// DecodeHTTPGetRoutines is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded GetRoutinesReply response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPGetRoutinesResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp pb.GetRoutinesReply
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

// Client Encode

// EncodeHTTPScrapTrelloResetZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a scraptrelloreset request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPScrapTrelloResetZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	fmt.Printf("Encoding request %v\n", request)
	req := request.(*pb.ScrapTrelloResetRequest)
	_ = req

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"scrap",
		"trello",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()

	// Set the body parameters
	var buf bytes.Buffer
	toRet := map[string]interface{}{}
	if err := json.NewEncoder(&buf).Encode(toRet); err != nil {
		return errors.Wrapf(err, "couldn't encode body as json %v", toRet)
	}
	r.Body = ioutil.NopCloser(&buf)
	fmt.Printf("URL: %v\n", r.URL)
	return nil
}

// EncodeHTTPGetRoutinesZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a getroutines request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPGetRoutinesZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	fmt.Printf("Encoding request %v\n", request)
	req := request.(*pb.GetRoutinesRequest)
	_ = req

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"routines",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	values.Add("date_start", fmt.Sprint(req.DateStart))

	values.Add("date_end", fmt.Sprint(req.DateEnd))

	r.URL.RawQuery = values.Encode()

	// Set the body parameters
	var buf bytes.Buffer
	toRet := map[string]interface{}{}
	if err := json.NewEncoder(&buf).Encode(toRet); err != nil {
		return errors.Wrapf(err, "couldn't encode body as json %v", toRet)
	}
	r.Body = ioutil.NopCloser(&buf)
	fmt.Printf("URL: %v\n", r.URL)
	return nil
}

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// PathParams takes a url and a gRPC-annotation style url template, and
// returns a map of the named parameters in the template and their values in
// the given url.
//
// PathParams does not support the entirety of the URL template syntax defined
// in third_party/googleapis/google/api/httprule.proto. Only a small subset of
// the functionality defined there is implemented here.
func PathParams(url string, urlTmpl string) (map[string]string, error) {
	rv := map[string]string{}
	pmp := BuildParamMap(urlTmpl)

	parts := strings.Split(url, "/")
	for k, v := range pmp {
		rv[k] = parts[v]
	}

	return rv, nil
}

// BuildParamMap takes a string representing a url template and returns a map
// indicating the location of each parameter within that url, where the
// location is the index as if in a slash-separated sequence of path
// components. For example, given the url template:
//
//     "/v1/{a}/{b}"
//
// The returned param map would look like:
//
//     map[string]int {
//         "a": 2,
//         "b": 3,
//     }
func BuildParamMap(urlTmpl string) map[string]int {
	rv := map[string]int{}

	parts := strings.Split(urlTmpl, "/")
	for idx, part := range parts {
		if strings.ContainsAny(part, "{}") {
			param := RemoveBraces(part)
			rv[param] = idx
		}
	}
	return rv
}

// RemoveBraces replace all curly braces in the provided string, opening and
// closing, with empty strings.
func RemoveBraces(val string) string {
	val = strings.Replace(val, "{", "", -1)
	val = strings.Replace(val, "}", "", -1)
	return val
}

// QueryParams takes query parameters in the form of url.Values, and returns a
// bare map of the string representation of each key to the string
// representation for each value. The representations of repeated query
// parameters is undefined.
func QueryParams(vals url.Values) (map[string]string, error) {

	rv := map[string]string{}
	for k, v := range vals {
		rv[k] = v[0]
	}
	return rv, nil
}
