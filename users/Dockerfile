FROM golang:alpine
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk update && apk add --no-cache git

WORKDIR /app


## Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# RUN go get -v github.com/cosmtrek/air

#
## Copy the code into the container
COPY . .

## Build the application
RUN go build -o /main main.go

CMD ["/main"]