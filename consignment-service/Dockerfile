FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD /gitrepo/go-micro-demo/consignment-service /app/consignment-service

CMD ["./consignment-service"]