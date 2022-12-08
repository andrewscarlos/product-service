ui:
	grpcui --plaintext localhost:5000

test-cover:
	go test --coverprofile coverfile.out product-service/$(path) && \
	go tool cover -html=coverfile.out -o coverage.html