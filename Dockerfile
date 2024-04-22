FROM golang:alpine3.19 as base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

    
FROM base as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o backend .


FROM golang:alpine3.19 as release

WORKDIR /home/node

RUN mkdir ./assets
RUN mkdir ./assets/image
RUN mkdir ./assets/video

COPY --from=builder /app/backend ./backend

EXPOSE $APP_PORT

CMD [ "./backend" ]