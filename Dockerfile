FROM golang:1.19

WORKDIR /usr/app

COPY go.mod ./

RUN go mod tidy

EXPOSE 3333

CMD [ "go", "run", "main.go" ]