FROM golang:1.22-alpine AS bob-construtor

WORKDIR /vrcursos

COPY go.mod go.sum ./
RUN go mod download

COPY src src

RUN go build -o bin/app src/main.go



FROM scratch AS runner

# since strach don't have shit, we need some certs to do some api to other
# services in the future
COPY --from=bob-construtor /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /vrcursos

COPY --from=bob-construtor /vrcursos/bin/app app

COPY migrations migrations

EXPOSE 3000

ENTRYPOINT ["./app"]