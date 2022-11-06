############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git bash wget curl
WORKDIR /build
RUN git clone --progress https://github.com/v2fly/v2fly-core.git . && \
    bash ./release/user-package.sh nosource noconf codename=$(git describe --abbrev=0 --tags) buildname=docker-fly abpathtgz=/tmp/v2fly.tgz

############################
# STEP 2 build a small image
############################
FROM alpine

LABEL maintainer "V2Fly Community <admin@v2fly.org>"
COPY --from=builder /tmp/v2fly.tgz /tmp
RUN apk update && apk add ca-certificates && \
    mkdir -p /usr/bin/v2fly && \
    tar xvfz /tmp/v2fly.tgz -C /usr/bin/v2fly

#ENTRYPOINT ["/usr/bin/v2fly/v2fly"]
ENV PATH /usr/bin/v2fly:$PATH
CMD ["v2fly", "-config=/etc/v2fly/config.json"]
