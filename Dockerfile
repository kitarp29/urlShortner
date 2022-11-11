FROM golang:1.19.3-alpine3.16
WORKDIR /k8s-api
COPY . .
RUN go mod tidy
CMD ["go","run","main.go"]
EXPOSE 8000