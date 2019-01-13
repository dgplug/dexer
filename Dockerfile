FROM golang:rc-alpine
MAINTAINER Kuntal Majumder (hellozee@disroot.org)

ENV APP_PATH=/go/src/github.com/dgplug/dexer/
COPY . ${APP_PATH}
RUN cd ${APP_PATH} && go install ./...
COPY ui /go/ui
COPY config.json /go/config.json
COPY logs /go/logs
CMD ["dexer"]
