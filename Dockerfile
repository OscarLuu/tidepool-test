FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /out/bootstrap .

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/bootstrap /app/bootstrap
COPY index.html /app/index.html
ENV PORT=8080
EXPOSE 8080
CMD ["/app/bootstrap"]
