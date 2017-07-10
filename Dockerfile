FROM alpine:3.5

ENV PORT 80
ENV TO https://lenka.blog
ENV TYPE 301

RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY simple-redirect /

EXPOSE $PORT

CMD ["/simple-redirect"]
