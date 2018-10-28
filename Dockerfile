FROM golang:1.11-alpine as build

ENV APP_PATH=/go/src/github.com/farhaanbukhsh/file-indexer/

COPY . ${APP_PATH}
RUN cd ${APP_PATH} &&\
  go build

FROM alpine:latest

COPY --from=build /go/src/github.com/farhaanbukhsh/file-indexer/file-indexer /usr/local/bin
COPY ui /ui
COPY config.json /config.json
COPY logs /logs

EXPOSE 8000
CMD [ "file-indexer" ]
