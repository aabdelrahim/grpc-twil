FROM alpine

RUN apk update
RUN apk add ca-certificates
ADD app /app
ENTRYPOINT ["/app"]