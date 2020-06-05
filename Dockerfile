FROM golang:1.14.2 as compiler
WORKDIR /src/app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./server.out .

FROM scratch
COPY --from=compiler /src/app/server.out /server
ENTRYPOINT ["/server"]
