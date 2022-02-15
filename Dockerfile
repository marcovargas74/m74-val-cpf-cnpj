FROM golang:1.17.0-alpine3.14 AS builder

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache build-base git 
 

# Set the working directory
WORKDIR /


COPY . .

# Fetch dependencies
#RUN go get -v ./...

# Build and strip the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/api-bank src/bank/main.go

FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/api-bank /go/bin/api-bank

# Run the api-bank binary.
CMD ["/go/bin/bank-api"]


