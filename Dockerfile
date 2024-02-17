FROM golang:1.21.6-alpine AS build

WORKDIR /Receipt-Processor/src
COPY . .

RUN go mod download
RUN go build -o /Receipt-Processor/src/Receipt-Processor

FROM alpine:latest

WORKDIR /Receipt-Processor/src
COPY --from=build /Receipt-Processor/src/Receipt-Processor .
CMD ["./Receipt-Processor"]
