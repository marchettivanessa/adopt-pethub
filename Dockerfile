FROM golang:1.21-alpine

WORKDIR /adopt-pethub

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

WORKDIR /adopt-pethub/backend/main
ENTRYPOINT ["go", "run", "main.go"]

