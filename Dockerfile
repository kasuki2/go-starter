FROM golang

WORKDIR /app


COPY backend/*.mod backend/*.sum ./

RUN go mod download


COPY backend/. .

RUN CGO_ENABLED=0 GOOS=linux go build -o main /app/cmd/http

EXPOSE 4002

CMD ["/app/main"]

