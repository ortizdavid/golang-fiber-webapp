
FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux go build -o golang-fiber-webapp 

# small image
FROM alpine
WORKDIR /app
COPY --from=builder /app /app/
EXPOSE 5000
CMD [ "./golang-fiber-webapp" ]