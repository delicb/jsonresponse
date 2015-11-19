// Utility package for sending json responses.
// Able to send to client any JSON serializable object. Object that is sent
// as response can be transformed via Transfrorm function which can be modified
// via SetTransformer function. Default implementation just wraps provided value
// in map with key "data" so it is pretty much useless, but wrapping is done
// because there is no restriction on type of data in response (just string or
// int is not valid JSON).
//
// Example of usage:
//
//     func someHandler(w http.ResponseWriter, r *Request) {
//         // create some object to return
//         obj := ...
//	       jsonResponse.New(obj).OK(w)
//     }
//
// Example of usage with custom header:
//
//     func someHandler(w http.ResponseWriter, r *Request) {
//         // create some object to return
//         obj := ...
//	       jsonResponse.New(obj).Header("X-Custom-Header", "this is cool").OK(w)
//     }

package jsonresponse

import (
	"encoding/json"
	"net/http"
	"fmt"
	"bytes"
)

// Response object, only contains object to return.
type Response struct {
	// Object to return, should be json serializable, or serialization will panic.
	Data    interface{}
	Headers map[string]string
	Excuse  string
}

func serializeToString(data interface{}) (s string ) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	if indent {
		var out bytes.Buffer
		json.Indent(&out, b, "", "\t")
		return out.String()
	} else {
		return string(b)
	}
}
// New creates response object with provided data and returns it.
func New(data interface{}) (r Response) {
	return Response{Data: data, Headers: map[string]string{}}
}

// Empty creates response object with not data. This can be useful since some
// http responses does not require data as response, only status code,
// like 204 (No Content)
func Empty() (r Response) {
	return Response{Data: nil, Headers: map[string]string{}}
}

// Response transforms body, sets headers and writes encoded body (to JSON) to
// provided writer.
func (r Response) Response(w http.ResponseWriter, httpCode int) {
	var headers map[string]string
	var body interface{}
	if transformer != nil && r.Data != nil {
		headers, body = transformer(r, httpCode)
	} else {
		headers = map[string]string{}
		body = r.Data
	}

	// if we have headers for this response, include it (and override transformer headers)
	for k, v := range r.Headers {
		headers[k] = v
	}

	// set headers to response writer
	responseHeaders := w.Header()
	for k, v := range headers {
		responseHeaders.Set(k, v)
	}

	// if Content-Type is not already included - add it here
	if _, ok := headers["Content-Type"]; !ok {
		responseHeaders.Set("Content-Type", defaultContentTypeHeader)
	}

	// write headers
	w.WriteHeader(httpCode)

	if body != nil {
		fmt.Fprint(w, serializeToString(body) + "\n")
	}
}

// Header adds header to response.
func (r Response) Header(key, value string) Response {
	r.Headers[key] = value
	return r
}

// WithProgrammingExcuse adds random programming excuse. This works with default
// transformer and Excuse field of Response has to be added to custom transfomer
// if it is used.
func (r Response) WithProgrammingExcuse() Response {
	r.Excuse = randomExcuse()
	return r
}

// 1xx

// Continue sends response to client with HTTP status 100.
// This means that server has received the request headers and that the client
// should proceed to send the request body
func (r Response) Continue(w http.ResponseWriter) {
	r.Response(w, http.StatusContinue)
}

// SwitchingProtocols sends response to client with HTTP status 101.
// This means the requester has asked the server to switch protocols and the
// server is acknowledging that it will do so.
func (r Response) SwitchingProtocols(w http.ResponseWriter) {
	r.Response(w, http.StatusSwitchingProtocols)
}

// 2xx

// OK sends response to client with HTTP status 200.
// Standard response for successful HTTP requests.
func (r Response) OK(w http.ResponseWriter) {
	r.Response(w, http.StatusOK)
}

// Created sends response to client with HTTP status 201.
// The request has been fulfilled and resulted in a new resource being created.
func (r Response) Created(w http.ResponseWriter) {
	r.Response(w, http.StatusCreated)
}

// Accepted sends response to client with HTTP status 202.
// The request has been accepted for processing, but the processing has not
// been completed.
func (r Response) Accepted(w http.ResponseWriter) {
	r.Response(w, http.StatusAccepted)
}

// NonAuthoritativeInfo sends response to client with HTTP status 203.
// The server successfully processed the request, but is returning information
// that may be from another source.
func (r Response) NonAuthoritativeInfo(w http.ResponseWriter) {
	r.Response(w, http.StatusNonAuthoritativeInfo)
}

// NoContent sends response to client with HTTP status 204.
// The server successfully processed the request, but is not returning any content.
func (r Response) NoContent(w http.ResponseWriter) {
	r.Response(w, http.StatusNoContent)
}

// ResetContent sends response to client with HTTP status 205.
// The server successfully processed the request, but is not returning any content.
// Unlike a NoContent response, this response requires that the requester reset
// the document view.
func (r Response) ResetContent(w http.ResponseWriter) {
	r.Response(w, http.StatusResetContent)
}

// PartialContent sends response to client with HTTP status 206.
// The server is delivering only part of the resource (byte serving) due to a
// range header sent by the client.
func (r Response) PartialContent(w http.ResponseWriter) {
	r.Response(w, http.StatusPartialContent)
}

// 3xx

// MultipleChoices sends response to client with HTTP status 300.
// Indicates multiple options for the resource that the client may follow.
func (r Response) MultipleChoices(w http.ResponseWriter) {
	r.Response(w, http.StatusMultipleChoices)
}

// MovedPermanently sends response to client with HTTP status 301.
// This and all future requests should be directed to the given URI.
func (r Response) MovedPermanently(w http.ResponseWriter) {
	r.Response(w, http.StatusMovedPermanently)
}

// Found sends response to client with HTTP status 302.
func (r Response) Found(w http.ResponseWriter) {
	r.Response(w, http.StatusFound)
}

// SeeOther sends response to client with HTTP status 303.
// The response to the request can be found under another URI using a GET method.
func (r Response) SeeOther(w http.ResponseWriter) {
	r.Response(w, http.StatusSeeOther)
}

// NotModified sends response to client with HTTP status 304.
// Indicates that the resource has not been modified since the version specified
// by the request headers If-Modified-Since or If-None-Match.
func (r Response) NotModified(w http.ResponseWriter) {
	r.Response(w, http.StatusNotModified)
}

// UseProxy sends response to client with HTTP status 305.
// The requested resource is only available through a proxy, whose address is
// provided in the response.
func (r Response) UseProxy(w http.ResponseWriter) {
	r.Response(w, http.StatusUseProxy)
}

// TemporaryRedirect sends response to client with HTTP status 307.
// In this case, the request should be repeated with another URI; however,
// future requests should still use the original URI.
func (r Response) TemporaryRedirect(w http.ResponseWriter) {
	r.Response(w, http.StatusTemporaryRedirect)
}

// 4xx

// BadRequest sends response to client with HTTP status 400.
// The server cannot or will not process the request due to something that is
// perceived to be a client error
func (r Response) BadRequest(w http.ResponseWriter) {
	r.Response(w, http.StatusBadRequest)
}

// Unauthorized sends response to client with HTTP status 401.
// Similar to 403 Forbidden, but specifically for use when authentication
// is required and has failed or has not yet been provided.
func (r Response) Unauthorized(w http.ResponseWriter) {
	r.Response(w, http.StatusUnauthorized)
}

// PaymentRequired sends response to client with HTTP status 402.
// Reserved for future use.
func (r Response) PaymentRequired(w http.ResponseWriter) {
	r.Response(w, http.StatusPaymentRequired)
}

// Forbidden sends response to client with HTTP status 403.
// The request was a valid request, but the server is refusing to respond to it.
// Unlike a 401 Unauthorized response, authenticating will make no difference.
func (r Response) Forbidden(w http.ResponseWriter) {
	r.Response(w, http.StatusForbidden)
}

// NotFound sends response to client with HTTP status 404.
// The requested resource could not be found but may be available again in the future.
func (r Response) NotFound(w http.ResponseWriter) {
	r.Response(w, http.StatusNotFound)
}

// MethodNotAllowed sends response to client with HTTP status 405.
// A request was made of a resource using a request method not supported by that
// resource; for example, using GET on a form which requires data to be presented
// via POST, or using PUT on a read-only resource.
func (r Response) MethodNotAllowed(w http.ResponseWriter) {
	r.Response(w, http.StatusMethodNotAllowed)
}

// NotAcceptable sends response to client with HTTP status 406.
// The requested resource is only capable of generating content not acceptable
// according to the Accept headers sent in the request.
func (r Response) NotAcceptable(w http.ResponseWriter) {
	r.Response(w, http.StatusNotAcceptable)
}

// ProxyAuthRequired sends response to client with HTTP status 407.
// The client must first authenticate itself with the proxy.
func (r Response) ProxyAuthRequired(w http.ResponseWriter) {
	r.Response(w, http.StatusProxyAuthRequired)
}

// RequestTimeout sends response to client with HTTP status 408.
// The server timed out waiting for the request.
func (r Response) RequestTimeout(w http.ResponseWriter) {
	r.Response(w, http.StatusRequestTimeout)
}

// Conflict sends response to client with HTTP status 409.
// Indicates that the request could not be processed because of conflict
// in the request.
func (r Response) Conflict(w http.ResponseWriter) {
	r.Response(w, http.StatusConflict)
}

// Gone sends response to client with HTTP status 410.
// Indicates that the resource requested is no longer available and will not
// be available again.
func (r Response) Gone(w http.ResponseWriter) {
	r.Response(w, http.StatusGone)
}

// LengthRequired sends response to client with HTTP status 411.
// The request did not specify the length of its content, which is
// required by the requested resource.
func (r Response) LengthRequired(w http.ResponseWriter) {
	r.Response(w, http.StatusLengthRequired)
}

// PreconditionFailed sends response to client with HTTP status 412.
// The server does not meet one of the preconditions that the requester put
// on the request.
func (r Response) PreconditionFailed(w http.ResponseWriter) {
	r.Response(w, http.StatusPreconditionFailed)
}

// RequestEntityTooLarge sends response to client with HTTP status 413.
// The request is larger than the server is willing or able to process.
func (r Response) RequestEntityTooLarge(w http.ResponseWriter) {
	r.Response(w, http.StatusRequestEntityTooLarge)
}

// RequestURITooLong sends response to client with HTTP status 414.
// The URI provided was too long for the server to process.
func (r Response) RequestURITooLong(w http.ResponseWriter) {
	r.Response(w, http.StatusRequestURITooLong)
}

// UnsupportedMediaType sends response to client with HTTP status 415.
// The request entity has a media type which the server or resource does
// not support.
func (r Response) UnsupportedMediaType(w http.ResponseWriter) {
	r.Response(w, http.StatusUnsupportedMediaType)
}

// RequestedRangeNotSatisfiable sends response to client with HTTP status 416.
// The client has asked for a portion of the file (byte serving), but the
// server cannot supply that portion.
func (r Response) RequestedRangeNotSatisfiable(w http.ResponseWriter) {
	r.Response(w, http.StatusRequestedRangeNotSatisfiable)
}

// ExpectationFailed sends response to client with HTTP status 417.
// The server cannot meet the requirements of the Expect request-header field.
func (r Response) ExpectationFailed(w http.ResponseWriter) {
	r.Response(w, http.StatusExpectationFailed)
}

// Teapot sends response to client with HTTP status 418.
// This code should be returned by tea pots requested to brew coffee.
func (r Response) Teapot(w http.ResponseWriter) {
	r.Response(w, http.StatusTeapot)
}

// 5xx

// InternalServerError sends response to client with HTTP status 500.
// A generic error message, given when an unexpected condition was
// encountered and no more specific message is suitable.
func (r Response) InternalServerError(w http.ResponseWriter) {
	r.Response(w, http.StatusInternalServerError)
}

// NotImplemented sends response to client with HTTP status 501.
// The server either does not recognize the request method, or it lacks the
// ability to fulfill the request.
func (r Response) NotImplemented(w http.ResponseWriter) {
	r.Response(w, http.StatusNotImplemented)
}

// BadGateway sends response to client with HTTP status 502.
// The server was acting as a gateway or proxy and received an invalid
// response from the upstream server.
func (r Response) BadGateway(w http.ResponseWriter) {
	r.Response(w, http.StatusBadGateway)
}

// ServiceUnavailable sends response to client with HTTP status 503.
// The server is currently unavailable (because it is overloaded or down
// for maintenance). Generally, this is a temporary state.
func (r Response) ServiceUnavailable(w http.ResponseWriter) {
	r.Response(w, http.StatusServiceUnavailable)
}

// GatewayTimeout sends response to client with HTTP status 504.
// The server was acting as a gateway or proxy and did not receive a timely
// response from the upstream server.
func (r Response) GatewayTimeout(w http.ResponseWriter) {
	r.Response(w, http.StatusGatewayTimeout)
}

// HTTPVersionNotSupported sends response to client with HTTP status 505.
// The server does not support the HTTP protocol version used in the request.
func (r Response) HTTPVersionNotSupported(w http.ResponseWriter) {
	r.Response(w, http.StatusHTTPVersionNotSupported)
}
