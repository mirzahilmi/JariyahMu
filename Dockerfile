    FROM golang:1.22.0-alpine3.19 AS Builder

LABEL authors="mirza.hilmi@gmail.com"

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/app

FROM alpine:3.19.1

WORKDIR /bin

ARG USER=runner
ARG GROUP=$USER

RUN addgroup -g 1000 runner && \
adduser -DH -g '' -G runner -u 1000 runner

COPY --from=Builder --chown=$USER:$GROUP --chmod=500 /src/app /src/.env ./

USER $USER:$GROUP

ENTRYPOINT ./app
