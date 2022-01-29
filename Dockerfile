FROM golang:1.17-alpine
WORKDIR /src
ADD ./ ./
ADD ./web/static /app/static
ADD ./web/template /app/template
ADD ./web/musings /app/musings
RUN go build -o /app/ergo ./cmd/ergo

FROM alpine:latest
EXPOSE 80
COPY --from=0 /app/ergo /app/ergo
COPY --from=0 /app/static /app/static
COPY --from=0 /app/template /app/template
COPY --from=0 /app/musings /app/musings
WORKDIR /app
CMD ["./ergo"]
