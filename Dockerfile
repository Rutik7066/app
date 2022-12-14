# syntax=docker/dockerfile:1
FROM golang:1.19
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD [ "/app/main" ]
EXPOSE 3000
