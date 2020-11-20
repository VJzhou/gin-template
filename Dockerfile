FROM alpine

WORKDIR /go/src/gin-demo
COPY . /go/src/gin-demo
EXPOSE 8088
CMD ["./main"]

