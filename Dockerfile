FROM golang:alpine as build

WORKDIR /checkout

COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /checkout/server ./cmd/server/

FROM scratch

WORKDIR /app

COPY --from=build /checkout/products.json .
COPY --from=build /checkout/server .

ENTRYPOINT ["/app/server"]