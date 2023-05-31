FROM spiksius/go-bash-1.19 AS build

RUN apk update && apk add ca-certificates
WORKDIR /src
COPY main.go .
COPY be be
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /service
COPY index.html .
FROM scratch AS bin
COPY --from=build /service /service
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/index.html /index.html
WORKDIR /
ENTRYPOINT ["/service"]