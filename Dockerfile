FROM golang:1.15

WORKDIR /productivity-monitoring

COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go build

EXPOSE 2112

CMD ["./productivity-monitoring"]
