FROM golang:alpine

RUN ["mkdir", "/app"]
WORKDIR /app

COPY ./main /app/

CMD ["./main"]