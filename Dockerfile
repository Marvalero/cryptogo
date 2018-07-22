FROM golang:1.9-alpine

WORKDIR /go/src/github.com/Marvalero/cryptogo/
COPY . .

RUN apk add --update bash curl git && \
    rm /var/cache/apk/*

RUN mkdir -p $$GOPATH/bin && \
    curl https://glide.sh/get | sh

RUN glide install
RUN CGO_ENABLED=0 GOOS=linux go build -o cryptogo


FROM alpine:3.7
LABEL name="cryptogo" \
      version="latest"

RUN apk --no-cache add ca-certificates
WORKDIR /home/
COPY --from=0 /go/src/github.com/Marvalero/cryptogo/cryptogo /usr/local/bin/

ENV PORT_NUM 9200
EXPOSE ${PORT_NUM}

ENTRYPOINT ["cryptogo"]
CMD ["-h"]
