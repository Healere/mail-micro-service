FROM alpine:latest

RUN mkdir /app

COPY authApp /app

COPY cmd/api/config/application.yml /app

WORKDIR /app

CMD [ "/app/authApp" ]

