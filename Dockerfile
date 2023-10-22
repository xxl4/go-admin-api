FROM alpine

# ENV GOPROXY https://goproxy.cn/

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update --no-cache
RUN apk add --update gcc g++ libc6-compat
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

COPY ./go-admin-api /go-admin-api
COPY ./config/settings.demo.yml /config/settings.yml
COPY ./go-admin-db.db /go-admin-db.db
EXPOSE 8000
RUN  chmod +x /go-admin-api
CMD ["/go-admin-api","server","-c", "/config/settings.yml"]