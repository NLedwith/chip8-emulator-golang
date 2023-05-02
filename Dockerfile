# syntax=docker/dockerfile:1
FROM golang:1.19
ENV DISPLAY :99
WORKDIR /app
COPY go.mod go.sum ./
RUN apt-get update -y
RUN apt-get install -y libgl1-mesa-dev
RUN apt-get install -y xorg-dev
RUN go mod download
RUN go get "github.com/go-gl/gl/v4.1-core/gl"
RUN go get "github.com/go-gl/glfw/v3.2/glfw"
COPY *.go ./
COPY /roms/*.ch8 ./roms/
RUN CGO_ENABLED=1 GOOS=linux go build -o /emulator
CMD ["/emulator"]
