FROM alpine:latest

ADD ./web /web
ADD ./public /public/
ADD ./result.csv /result.csv
ADD ./trains.csv /trains.csv
ADD ./updated /updated

ENTRYPOINT ["./web"]
