FROM golang:1.17.3 AS build

WORKDIR /go/src/aoc/day2
COPY . .
RUN go get -d -v ./...
RUN GCO_ENABLED=0 && go build -o day2.git

FROM alpine:3.15.0
RUN apk add --no-cache libc6-compat
COPY --from=build /go/src/aoc/day2/day2.git /day2
COPY ./input.txt /input.txt
CMD ["/day2"]