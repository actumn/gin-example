## Stage :: build
FROM golang:1.10 AS Builder
# copy the code from the host
WORKDIR $GOPATH/src/gin-example
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

## Stage :: run
# run
FROM ubuntu
COPY --from=builder /app ./
COPY config ./config
RUN apt-get update
RUN apt-get install -y ca-certificates
CMD ["./app"]
ENV PORT 80
EXPOSE 80
