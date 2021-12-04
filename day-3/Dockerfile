FROM golang:1.17.3 AS build

WORKDIR /go/src/aoc/day3
COPY . .
RUN go get -d -v ./...
RUN GCO_ENABLED=0 && go build -o day3.git

FROM alpine:3.15.0
RUN apk add --no-cache libc6-compat
COPY --from=build /go/src/aoc/day2/day3.git /day3
COPY ./input.txt /input.txt
CMD ["/day3"]