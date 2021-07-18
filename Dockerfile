FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . ./

RUN go build -o /case-study

EXPOSE 8080

CMD [ "/case-study" ]