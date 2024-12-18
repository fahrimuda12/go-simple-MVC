FROM golang:1.23

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application code
COPY . ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/docker-go-simple-struture


EXPOSE 8080

CMD ["/bin/docker-go-simple-struture"]