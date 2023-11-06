FROM golang:1.20.4-alpine3.17
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go build --ldflags "-s -w" -race -o ./build/app ./cmd/main.go

FROM alpine:latest
LABEL description="golang,app"
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add tzdata
ENV TZ=Asia/Shanghai
WORKDIR /app
COPY --from=0 /app/conf/config-dev.yaml ./conf/
COPY --from=0 /app/build/app ./build/
CMD [ "./build/app"]

