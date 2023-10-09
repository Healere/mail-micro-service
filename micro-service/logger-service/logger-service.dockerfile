FROM alpine:latest

RUN mkdir /app

COPY loggerServiceApp /app

COPY cmd/api/config/application.yml /app

WORKDIR /app

CMD [ "/app/loggerServiceApp"]