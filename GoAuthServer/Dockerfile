FROM golang:1.16-alpine AS build

WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /app/auth

FROM alpine:latest
COPY --from=build /app/auth /app/auth
ENTRYPOINT ["/app/auth"]