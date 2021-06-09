FROM golang:1.16-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /app/main

FROM alpine:latest
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]