FROM golang:1.19

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

EXPOSE 3001

CMD [ "go", "run", "cmd/product_service/main.go" ]