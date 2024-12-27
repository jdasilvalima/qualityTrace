FROM golang:1.23.4-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go get github.com/99designs/gqlgen

COPY go.* ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]