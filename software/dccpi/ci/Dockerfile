FROM --platform=linux/x86_64 golang:1.20 as base

ENV APPLICATION_PATH="/var/app" \
    GOOS="linux" \
    GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOFLAGS="-mod=vendor" \
    GOPATH="/usr/go" \
    GOPRIVATE="github.com/alexbegoon"

COPY . ${APPLICATION_PATH}
WORKDIR ${APPLICATION_PATH}

FROM base as development

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN $(go env GOPATH)/bin/air -v

CMD ${GOPATH}/bin/air -c ${APPLICATION_PATH}/conf/.air.toml