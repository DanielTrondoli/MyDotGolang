FROM golang:1.16-alpine

WORKDIR /app

COPY . .

# RUN go clean --modcache
RUN go mod download
RUN go build -o /mydotgolang

EXPOSE 1666

CMD [ "/mydotgolang" ]
