package interfaces

type Controller interface {
	Handler(request *HttpRequest) *HttpResponse
}
