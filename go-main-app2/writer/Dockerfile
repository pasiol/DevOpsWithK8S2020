ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine AS dev

ENV APP_NAME="main" \
    APP_PATH="/var/app"

ENV APP_BUILD_NAME="${APP_NAME}"

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-mod=vendor"

ENTRYPOINT ["sh"]

FROM dev as build

RUN (([ ! -d "${APP_PATH}/vendor" ] && go mod download && go mod vendor) || true)
RUN go build -ldflags="-s -w" -mod vendor -o ${APP_BUILD_NAME} main.go
RUN chmod +x ${APP_BUILD_NAME}

FROM scratch AS prod

ENV APP_BUILD_PATH="/var/app" \
    APP_BUILD_NAME="main" \
    APP_LOG_FILE="/var/app/logs/writer.log" \
    APP_OUTPUT_FILE="/var/app/data/uuids.txt" 
WORKDIR ${APP_BUILD_PATH}
#RUN mkdir logs data
COPY --from=build ${APP_BUILD_PATH}/${APP_BUILD_NAME} ${APP_BUILD_PATH}/

ENTRYPOINT ["/var/app/main"]
CMD ""
