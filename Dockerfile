FROM golang:1.20.1

RUN mkdir /app

WORKDIR $GOPATH/src/github.com/Payload-solutions/iot-microservice

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...


# here
# COPY --from=client_builder /app/dist ./client/dist

EXPOSE 8000

# ENTRYPOINT measurements-realtime --build="go build main.go" --command=./main

CMD ["iot-microservice"]
