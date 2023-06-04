FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/event-api
RUN go build -buildvcs=false -o /api

CMD /api