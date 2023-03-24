FROM golang:latest
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go get "github.com/mattn/go-tty"
COPY *.go ./
COPY ./roms/*.ch8 ./roms/
RUN go build -o /emulator
CMD ["/emulator"]

