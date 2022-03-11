# Build
FROM golang:1.17-alpine AS build

#ARG VERSION=snapshot
#ENV VENDORED=true

RUN apk --update --no-cache add make

WORKDIR /app
ADD . /app

RUN make build

# Run
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=build /app/go-app /go-app
USER 65532:65532

ENTRYPOINT ["/go-app"]
