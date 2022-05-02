FROM alpine

COPY ./bin/amd64/http-server /http-server

EXPOSE 8000

ENTRYPOINT /http-server

