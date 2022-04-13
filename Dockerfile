FROM node:14.19 as web-builder

WORKDIR /builder
COPY ./web/package-*.json .
RUN npm install

COPY ./web .
RUN npm run build

FROM golang:1.18 as service-builder
ARG VERSION=vNext
ARG COMMIT_HASH="n/a"
ARG BUILD_TIMESTAMP="n/a"


WORKDIR /builder
COPY go.* .
RUN go mod download

COPY . .
RUN echo "ts: $BUILD_TIMESTAMP" > ts.txt
RUN go build -v \
  -ldflags="-s -w  \
    -X github.com/avanibbles/flowflow/internal.Version=${VERSION} \
    -X github.com/avanibbles/flowflow/internal.CommitHash=${COMMIT_HASH} \
    -X github.com/avanibbles/flowflow/internal.BuildTimestamp=${BUILD_TIMESTAMP}" \
  ./cmd/flowflow

FROM debian:buster-slim as runtime

WORKDIR /app

ENV FLOWFLOW_HTTP_SITEDATA=/app/web
COPY --from=service-builder /builder/flowflow .
COPY --from=web-builder /builder/build/ ./web/


CMD ["/app/flowflow"]