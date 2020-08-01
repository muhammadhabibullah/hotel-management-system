FROM golang:1.14-alpine AS build

RUN apk add --no-cache git gcc libc-dev curl

WORKDIR /build

ADD . .

RUN go get -v

RUN CGO_ENABLED=0 go build -o hotel_mgmt_svc

FROM alpine

WORKDIR /usr/local/bin

COPY --from=build /build/hotel_mgmt_svc .
COPY --from=build /build/.env .

RUN chmod +x hotel_mgmt_svc

CMD ["hotel_mgmt_svc", "serve"]
