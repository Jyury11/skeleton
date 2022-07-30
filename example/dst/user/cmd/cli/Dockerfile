FROM golang:1.18 as build
WORKDIR /go/src/main
COPY . .
RUN cd ./user/cmd/cli && go install .

FROM gcr.io/distroless/base
COPY --from=build /go/bin/cli /
CMD ["/cli"]
