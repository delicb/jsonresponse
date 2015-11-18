package jsonresponse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var defaultTransformerBodies = []map[string]interface{}{
	map[string]interface{}{"data": 1.0},
	map[string]interface{}{"data": true},
	map[string]interface{}{"data": map[string]interface{}{"key": 42.0}},
}

func TestCustomTransformerCalled(t *testing.T) {
	isCalled := false
	customTransformer := func(r Response, httpCode int) (headers map[string]string, response interface{}) {
		isCalled = true
		return map[string]string{}, map[string]string{"value": "do not care what I got"}
	}
	SetTransformer(customTransformer)
	recorder := httptest.NewRecorder()
	New("does not matter").OK(recorder)
	if !isCalled {
		fmt.Println("Custom transformer not called")
		t.Fail()
	}
	ResetTransformer()
}

func TestCustomTransformerSettingHeaders(t *testing.T) {
	customHeaderKey := "X-My-Custom-Header"
	customHeaderValue := "random header value"
	customTransformer := func(r Response, httpCode int) (headers map[string]string, response interface{}) {
		return map[string]string{customHeaderKey: customHeaderValue}, map[string]string{"value": "do not care what I got"}
	}

	SetTransformer(customTransformer)
	recorder := httptest.NewRecorder()
	New("does not matter").OK(recorder)
	if recorder.Header().Get(customHeaderKey) != customHeaderValue {
		fmt.Println("Custom transformer header not applied.")
		t.Fail()
	}
	ResetTransformer()
}

func TestCreationRaw(t *testing.T) {
	response := New("data")
	if response.Data != "data" {
		t.Fail()
	}
	if len(response.Headers) != 0 {
		t.Fail()
	}
	if response.Excuse != "" {
		t.Fail()
	}
}

func TestSettingHeaders(t *testing.T) {
	response := Empty().Header("X-Custom-Header", "value")

	if len(response.Headers) != 1 {
		t.Fail()
	}

	if _, ok := response.Headers["X-Custom-Header"]; !ok {
		t.Fail()
	}
}

func TestSettingProgrammingExcuse(t *testing.T) {
	response := New("").WithProgrammingExcuse()
	if response.Excuse == "" {
		t.Fail()
	}
}

func TestResponseSerializationDefaultTransformer(t *testing.T) {

	for _, v := range defaultTransformerBodies {
		recorder := httptest.NewRecorder()

		New(v["data"]).OK(recorder)

		// decode from json
		unmarshaled := make(map[string]interface{})

		if err := json.Unmarshal(recorder.Body.Bytes(), &unmarshaled); err != nil {
			fmt.Println("Failed do unmarshal response!: ", err)
			t.Fail()
		}
		if !reflect.DeepEqual(unmarshaled, v) {
			fmt.Printf("Expected %#v\nbut got  %#v\n", v, unmarshaled)
			fmt.Println("Serialized values to not match!")
			t.Fail()
		}
	}
}

func TestResponseCodesResponseObject(t *testing.T) {
	for k, v := range map[int]func(r Response, w http.ResponseWriter){
		// 1xx
		100: Response.Continue,
		101: Response.SwitchingProtocols,

		// 2xx
		200: Response.OK,
		201: Response.Created,
		202: Response.Accepted,
		203: Response.NonAuthoritativeInfo,
		204: Response.NoContent,
		205: Response.ResetContent,
		206: Response.PartialContent,

		// 3xx
		300: Response.MultipleChoices,
		301: Response.MovedPermanently,
		302: Response.Found,
		303: Response.SeeOther,
		304: Response.NotModified,
		305: Response.UseProxy,
		307: Response.TemporaryRedirect,

		// 4xx
		400: Response.BadRequest,
		401: Response.Unauthorized,
		402: Response.PaymentRequired,
		403: Response.Forbidden,
		404: Response.NotFound,
		405: Response.MethodNotAllowed,
		406: Response.NotAcceptable,
		407: Response.ProxyAuthRequired,
		408: Response.RequestTimeout,
		409: Response.Conflict,
		410: Response.Gone,
		411: Response.LengthRequired,
		412: Response.PreconditionFailed,
		413: Response.RequestEntityTooLarge,
		414: Response.RequestURITooLong,
		415: Response.UnsupportedMediaType,
		416: Response.RequestedRangeNotSatisfiable,
		417: Response.ExpectationFailed,
		418: Response.Teapot,

		// 5xx
		500: Response.InternalServerError,
		501: Response.NotImplemented,
		502: Response.BadGateway,
		503: Response.ServiceUnavailable,
		504: Response.GatewayTimeout,
		505: Response.HTTPVersionNotSupported,
	} {
		recorder := httptest.NewRecorder()
		response := Empty()
		v(response, recorder)
		if recorder.Code != k {
			fmt.Printf("HTTP code did not match, got %d, expected: %d\n", recorder.Code, k)
		}
	}
}

func TestResponseCodesGeneric(t *testing.T) {
	for k, v := range map[int]func(w http.ResponseWriter, response interface{}){
		// 1xx
		100: Continue,
		101: SwitchingProtocols,

		// 2xx
		200: OK,
		201: Created,
		202: Accepted,
		203: NonAuthoritativeInfo,
		204: NoContent,
		205: ResetContent,
		206: PartialContent,

		// 3xx
		300: MultipleChoices,
		301: MovedPermanently,
		302: Found,
		303: SeeOther,
		304: NotModified,
		305: UseProxy,
		307: TemporaryRedirect,

		// 4xx
		400: BadRequest,
		401: Unauthorized,
		402: PaymentRequired,
		403: Forbidden,
		404: NotFound,
		405: MethodNotAllowed,
		406: NotAcceptable,
		407: ProxyAuthRequired,
		408: RequestTimeout,
		409: Conflict,
		410: Gone,
		411: LengthRequired,
		412: PreconditionFailed,
		413: RequestEntityTooLarge,
		414: RequestURITooLong,
		415: UnsupportedMediaType,
		416: RequestedRangeNotSatisfiable,
		417: ExpectationFailed,
		418: Teapot,

		// 5xx
		500: InternalServerError,
		501: NotImplemented,
		502: BadGateway,
		503: ServiceUnavailable,
		504: GatewayTimeout,
		505: HTTPVersionNotSupported,
	} {
		recorder := httptest.NewRecorder()
		v(recorder, "")
		if recorder.Code != k {
			fmt.Printf("HTTP code did not match, got %d, expected: %d\n", recorder.Code, k)
		}
	}
}
