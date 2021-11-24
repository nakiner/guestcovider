FROM golang:1.17-alpine AS build

ENV APP=./cmd/app
ENV BIN=/bin/guestcovider
ENV PATH_ROJECT=${GOPATH}/src/github.com/nakiner/guestcovider
ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOFLAGS=-mod=vendor
ARG VERSION
ENV VERSION ${VERSION:-0.1.0}
ARG BUILD_TIME
ENV BUILD_TIME ${BUILD_TIME:-unknown}
ARG COMMIT
ENV COMMIT ${COMMIT:-unknown}

ENV TZ Europe/Moscow
RUN apk add --update --no-cache tzdata

WORKDIR ${PATH_ROJECT}
COPY . ${PATH_ROJECT}

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w \
        -X github.com/nakiner/guestcovider/pkg/health.Version=${VERSION} \
        -X github.com/nakiner/guestcovider/pkg/health.Commit=${COMMIT} \
        -X github.com/nakiner/guestcovider/pkg/health.BuildTime=${BUILD_TIME}" \
    -a -o ${BIN} ${APP}

FROM alpine:3.14 as production

ENV TZ Europe/Moscow
RUN apk add --update --no-cache tzdata

COPY --from=build /bin/guestcovider /bin/guestcovider
ENTRYPOINT ["/bin/guestcovider"]
