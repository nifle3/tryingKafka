FROM ubuntu:latest
LABEL authors="Nifie"

WORKDIR ./app
ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./main cmd/main.go

ENTRYPOINT ["./main"]