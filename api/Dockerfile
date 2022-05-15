FROM golang:1.16-alpine

ARG PORT=8081

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN go build -o /api

EXPOSE ${PORT}

CMD [ "/api" ]
