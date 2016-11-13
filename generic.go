package jsonresponse // import "go.delic.rs/jsonresponse"

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MessageResponse is default wrapper structure for JSON http response.
type MessageResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Respond serializes provided response to JSON and writes it to provided writer
// with status code.
func Respond(w http.ResponseWriter, statusCode int, response interface{}) {
	if response == nil {
		response = &MessageResponse{}
	} else if r, ok := response.(MessageResponse); ok {
		response = &r
	}
	if r, ok := response.(*MessageResponse); ok {
		if r.Code == 0 {
			r.Code = statusCode
		}
		if r.Message == "" {
			r.Message = http.StatusText(statusCode)
		}
	}
	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	if defaultContentTypeHeader != "" {
		w.Header().Set("Content-Type", defaultContentTypeHeader)
	}
	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(b)+"\n")
}

// 1xx

// Continue sends response to client with HTTP status 100.
// This means that server has received the request headers and that the client
// should proceed to send the request body
func Continue(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusContinue, response)
}

// SwitchingProtocols sends response to client with HTTP status 101.
// This means the requester has asked the server to switch protocols and the
// server is acknowledging that it will do so.
func SwitchingProtocols(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusSwitchingProtocols, response)
}

// 2xx

// OK sends response to client with HTTP status 200.
// Standard response for successful HTTP requests.
func OK(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusOK, response)
}

// Created sends response to client with HTTP status 201.
// The request has been fulfilled and resulted in a new resource being created.
func Created(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusCreated, response)
}

// Accepted sends response to client with HTTP status 202.
// The request has been accepted for processing, but the processing has not
// been completed.
func Accepted(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusAccepted, response)
}

// NonAuthoritativeInfo sends response to client with HTTP status 203.
// The server successfully processed the request, but is returning information
// that may be from another source.
func NonAuthoritativeInfo(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusNonAuthoritativeInfo, response)
}

// NoContent sends response to client with HTTP status 204.
// The server successfully processed the request, but is not returning any content.
func NoContent(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusNoContent, response)
}

// ResetContent sends response to client with HTTP status 205.
// The server successfully processed the request, but is not returning any content.
// Unlike a NoContent response, this response requires that the requester reset
// the document view.
func ResetContent(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusResetContent, response)
}

// PartialContent sends response to client with HTTP status 206.
// The server is delivering only part of the resource (byte serving) due to a
// range header sent by the client.
func PartialContent(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusPartialContent, response)
}

// 3xx

// MultipleChoices sends response to client with HTTP status 300.
// Indicates multiple options for the resource that the client may follow.
func MultipleChoices(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusMultipleChoices, response)
}

// MovedPermanently sends response to client with HTTP status 301.
// This and all future requests should be directed to the given URI.
func MovedPermanently(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusMovedPermanently, response)
}

// Found sends response to client with HTTP status 302.
func Found(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusFound, response)
}

// SeeOther sends response to client with HTTP status 303.
// The response to the request can be found under another URI using a GET method.
func SeeOther(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusSeeOther, response)
}

// NotModified sends response to client with HTTP status 304.
// Indicates that the resource has not been modified since the version specified
// by the request headers If-Modified-Since or If-None-Match.
func NotModified(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusNotModified, response)
}

// UseProxy sends response to client with HTTP status 305.
// The requested resource is only available through a proxy, whose address is
// provided in the response.
func UseProxy(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusUseProxy, response)
}

// TemporaryRedirect sends response to client with HTTP status 307.
// In this case, the request should be repeated with another URI; however,
// future requests should still use the original URI.
func TemporaryRedirect(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusTemporaryRedirect, response)
}

// 4xx

// BadRequest sends response to client with HTTP status 400.
// The server cannot or will not process the request due to something that is
// perceived to be a client error
func BadRequest(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusBadRequest, response)
}

// Unauthorized sends response to client with HTTP status 401.
// Similar to 403 Forbidden, but specifically for use when authentication
// is required and has failed or has not yet been provided.
func Unauthorized(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusUnauthorized, response)
}

// PaymentRequired sends response to client with HTTP status 402.
// Reserved for future use.
func PaymentRequired(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusPaymentRequired, response)
}

// Forbidden sends response to client with HTTP status 403.
// The request was a valid request, but the server is refusing to respond to it.
// Unlike a 401 Unauthorized response, authenticating will make no difference.
func Forbidden(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusForbidden, response)
}

// NotFound sends response to client with HTTP status 404.
// The requested resource could not be found but may be available again in the future.
func NotFound(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusNotFound, response)
}

// MethodNotAllowed sends response to client with HTTP status 405.
// A request was made of a resource using a request method not supported by that
// resource; for example, using GET on a form which requires data to be presented
// via POST, or using PUT on a read-only resource.
func MethodNotAllowed(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusMethodNotAllowed, response)
}

// NotAcceptable sends response to client with HTTP status 406.
// The requested resource is only capable of generating content not acceptable
// according to the Accept headers sent in the request.
func NotAcceptable(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusNotAcceptable, response)
}

// ProxyAuthRequired sends response to client with HTTP status 407.
// The client must first authenticate itself with the proxy.
func ProxyAuthRequired(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusProxyAuthRequired, response)
}

// RequestTimeout sends response to client with HTTP status 408.
// The server timed out waiting for the request.
func RequestTimeout(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusRequestTimeout, response)
}

// Conflict sends response to client with HTTP status 409.
// Indicates that the request could not be processed because of conflict
// in the request.
func Conflict(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusConflict, response)
}

// Gone sends response to client with HTTP status 410.
// Indicates that the resource requested is no longer available and will not
// be available again.
func Gone(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusGone, response)
}

// LengthRequired sends response to client with HTTP status 411.
// The request did not specify the length of its content, which is
// required by the requested resource.
func LengthRequired(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusLengthRequired, response)
}

// PreconditionFailed sends response to client with HTTP status 412.
// The server does not meet one of the preconditions that the requester put
// on the request.
func PreconditionFailed(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusPreconditionFailed, response)
}

// RequestEntityTooLarge sends response to client with HTTP status 413.
// The request is larger than the server is willing or able to process.
func RequestEntityTooLarge(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusRequestEntityTooLarge, response)
}

// RequestURITooLong sends response to client with HTTP status 414.
// The URI provided was too long for the server to process.
func RequestURITooLong(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusRequestURITooLong, response)
}

// UnsupportedMediaType sends response to client with HTTP status 415.
// The request entity has a media type which the server or resource does
// not support.
func UnsupportedMediaType(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusUnsupportedMediaType, response)
}

// RequestedRangeNotSatisfiable sends response to client with HTTP status 416.
// The client has asked for a portion of the file (byte serving), but the
// server cannot supply that portion.
func RequestedRangeNotSatisfiable(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusRequestedRangeNotSatisfiable, response)
}

// ExpectationFailed sends response to client with HTTP status 417.
// The server cannot meet the requirements of the Expect request-header field.
func ExpectationFailed(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusExpectationFailed, response)
}

// Teapot sends response to client with HTTP status 418.
// This code should be returned by tea pots requested to brew coffee.
func Teapot(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusTeapot, response)
}

// 5xx

// InternalServerError sends response to client with HTTP status 500.
// A generic error message, given when an unexpected condition was
// encountered and no more specific message is suitable.
func InternalServerError(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusInternalServerError, response)
}

// NotImplemented sends response to client with HTTP status 501.
// The server either does not recognize the request method, or it lacks the
// ability to fulfill the request.
func NotImplemented(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusNotImplemented, response)
}

// BadGateway sends response to client with HTTP status 502.
// The server was acting as a gateway or proxy and received an invalid
// response from the upstream server.
func BadGateway(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusBadGateway, response)
}

// ServiceUnavailable sends response to client with HTTP status 503.
// The server is currently unavailable (because it is overloaded or down
// for maintenance). Generally, this is a temporary state.
func ServiceUnavailable(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusServiceUnavailable, response)
}

// GatewayTimeout sends response to client with HTTP status 504.
// The server was acting as a gateway or proxy and did not receive a timely
// response from the upstream server.
func GatewayTimeout(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusGatewayTimeout, response)
}

// HTTPVersionNotSupported sends response to client with HTTP status 505.
// The server does not support the HTTP protocol version used in the request.
func HTTPVersionNotSupported(w http.ResponseWriter, response interface{}) {
	Respond(w, http.StatusHTTPVersionNotSupported, response)
}
