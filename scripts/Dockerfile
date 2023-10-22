FROM alpine

COPY ./go-admin-api /
EXPOSE 8000

ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

CMD ["/go-admin-api","server","-c", "/config/settings.yml"]
