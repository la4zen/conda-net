FROM alpine

RUN mkdir /app
COPY configs /app
COPY main /app

CMD ["/app/main"]