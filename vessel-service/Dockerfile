FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD /gitrepo/go-micro-demo/vessel-service /app/vessel-service

CMD ["./vessel-service"]