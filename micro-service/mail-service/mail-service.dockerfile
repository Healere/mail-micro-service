FROM alpine:latest

RUN mkdir /app

COPY mailerApp /app
COPY templates /templates
COPY cmd/api/config/application.yml /app

WORKDIR /app

CMD [ "/app/mailerApp"]