FROM golang:1.22.1 AS build  
WORKDIR /app  
COPY main.go .  
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -a -installsuffix cgo -ldflags '-extldflags "-static"' main.go
  
FROM alpine:3.16  
WORKDIR /opt  
COPY --from=build /app/main .
CMD ["/opt/main"]