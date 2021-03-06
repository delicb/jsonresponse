package jsonresponse

// ResponseTransformer is function that transforms response before it is send
// to client. Default implementation is provided, but it can suite
// more specific needs.
type ResponseTransformer func(resp Response, httpCode int) (headers map[string]string, result interface{})

// PassthroughTransformer only returns data as they are in response without modification.
func PassthroughTransformer(resp Response, httpCode int) (headers map[string]string, result interface{}) {
	return make(map[string]string), resp.Data
}

// MessageCodeTransformer wraps response into map with data and code fields.
// Data and code fields can be defined as function parameters.
func MessageCodeTransformer(dataField string, codeField string) ResponseTransformer {
	return ResponseTransformer(func(resp Response, httpCode int) (headers map[string]string, result interface{}) {
		h := map[string]string{}
		r := map[string]interface{}{
			dataField: resp.Data,
			codeField: httpCode,
		}
		return h, r
	})
}

// MessageCodeExcuseTransformer is same as MessageCodeTransformer, except that
// it adds "programming-excuse" field with Excuse field from response.
func MessageCodeExcuseTransformer(dataField string, codeField string) ResponseTransformer {
	return ResponseTransformer(func(resp Response, httpCode int) (headers map[string]string, result interface{}) {
		h := map[string]string{}
		r := map[string]interface{}{
			dataField: resp.Data,
			codeField: httpCode,
		}
		if resp.Excuse != "" {
			r["programming-excuse"] = resp.Excuse
		}
		return h, r
	})
}

// defaultTransformer just wraps response dict with key data.
func defaultTransformer(resp Response, httpCode int) (headers map[string]string, result interface{}) {
	h := map[string]string{}
	r := map[string]interface{}{
		"data": resp.Data,
	}
	if resp.Excuse != "" {
		r["programming-excuse"] = resp.Excuse
	}
	return h, r
}
