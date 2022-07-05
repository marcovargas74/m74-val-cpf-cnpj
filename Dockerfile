FROM golang:1.17.0-alpine3.14 AS builder

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache build-base git 
 

# Set the working directory
WORKDIR $GOPATH/src/validator-app

COPY . .

# Fetch dependencies
#RUN go get -v ./...

# Build and strip the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/validator-app src/validator/main.go

FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/validator-app /go/bin/validator-app

# Run the validator-app binary.
CMD ["/go/bin/validator-app"]


