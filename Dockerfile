FROM golang:1.22.5

WORKDIR /app
COPY . .
RUN go mod download

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /hexacrosswords-api
EXPOSE 8080

# Run
CMD ["/hexacrosswords-api"]