FROM golang:1.25-alpine AS building

WORKDIR /building
COPY . .
ENV PRODUCTION_BUILD="true"
RUN go mod tidy
RUN go build -o /building/app .

FROM alpine as app_bases

WORKDIR /app

COPY --from=building /backend/app /app/server

EXPOSE 3500
ENTRYPOINT ["/app/server"]