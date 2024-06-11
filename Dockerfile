# golang image

FROM golang:1.22.2-alpine3.13

# working dir
WORKDIR /backend

# Copy this code to inside image
COPY . .

# Deps
RUN go get -d -v ./...

# Build
RUN go build -o msync .

EXPOSE 5000

CMD ["./api"]
