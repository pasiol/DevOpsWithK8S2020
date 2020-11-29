ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine AS dev

ENV APP_NAME="main" \
    APP_PATH="/var/app" \
    APP_PORT=3000

ENV APP_BUILD_NAME="${APP_NAME}"

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOFLAGS="-mod=vendor"

EXPOSE ${APP_PORT}
ENTRYPOINT ["sh"]

FROM dev as build

RUN (([ ! -d "${APP_PATH}/vendor" ] && go mod download && go mod vendor) || true)
RUN go build -ldflags="-s -w" -mod vendor -o ${APP_BUILD_NAME} main.go
RUN chmod +x ${APP_BUILD_NAME}
RUN mkdir pvc

FROM scratch AS prod

ENV APP_BUILD_PATH="/var/app" \
    APP_BUILD_NAME="main" \
    APP_LOG_FILE="/var/app/pvc/reader.log" \
    APP_OUTPUT_FILE="/var/app/pvc/uuids.txt" 
WORKDIR ${APP_BUILD_PATH}
COPY --from=build ${APP_BUILD_PATH}/${APP_BUILD_NAME} ${APP_BUILD_PATH}/

EXPOSE ${APP_PORT}
ENTRYPOINT ["/var/app/main"]
CMD ""
