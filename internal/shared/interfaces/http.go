package interfaces

type HttpResponse struct {
	StatusCode int
	JsonBody   []byte
}

type HttpRequest struct {
	Params []byte
	Body   []byte
}
