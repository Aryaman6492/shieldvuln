FROM --platform=$BUILDPLATFORM golang:1.23-bookworm AS builder

ENV GO111MODULE=on CGO_ENABLED=0
WORKDIR /work
ARG TARGETOS TARGETARCH

RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /out/shieldvuln cmd/http/main.go

FROM gcr.io/distroless/static-debian12:nonroot

USER nonroot
WORKDIR /home/nonroot/

COPY --from=builder /out/shieldvuln /usr/bin/shieldvuln

ARG image_version
ENV RELEASE=$image_version

ENTRYPOINT ["shieldvuln"]
