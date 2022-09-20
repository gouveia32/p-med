FROM golang:1.15
WORKDIR ./p-med
COPY . . 
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
# Install server application
CMD ["go", "run", "main.go"]