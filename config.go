package jsonresponse

import "sync"

var (
	transformerLock = &sync.Mutex{}

	// Transformer for response. Default implementation wraps response in
	// SBG envelope (with status and message).
	transformer ResponseTransformer = defaultTransformer
)

var (
	contentTypeHeaderLock = &sync.Mutex{}

	// Default Content-Type header for Json responses.
	defaultContentTypeHeader string = "application/json; charset=utf-8"
)

var (
	indentLock = &sync.Mutex{}

	// indent sets flag that indicates that all json responses will be
	// indented before returned. This can be useful for debugging when
	// consumer of JSON response is developer and not other service.
	indent bool = false
)

// SetTransformer sets function that will process response additionally.
// Default implementation just wraps response value in map with key "data".
func SetTransformer(t ResponseTransformer) {
	transformerLock.Lock()
	defer transformerLock.Unlock()
	transformer = t
}

// ResetTransformer resets current transformer to default one.
func ResetTransformer() {
	transformerLock.Lock()
	defer transformerLock.Unlock()
	transformer = defaultTransformer
}

// SetDefaultContentTypeHeader sets string that will be included in header
// under Content-Type header. This will only be included if transformer function
// does not already set Content-Type header.
func SetDefaultContentTypeHeader(contentType string) {
	contentTypeHeaderLock.Lock()
	defer contentTypeHeaderLock.Unlock()
	defaultContentTypeHeader = contentType
}

// SetIndent sets flat that indicates that JSON response messages should
// be indented before sending them to client. This can be useful for
// debugging when consumer of JSON response is developer and not
// other service.
func SetIndent(flag bool) {
	indentLock.Lock()
	defer indentLock.Unlock()
	indent = flag
}
