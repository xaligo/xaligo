FROM ubuntu:24.04 AS wasm-builder

ARG JAVY_VERSION=9.0.0
ARG TARGETARCH

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    gzip \
    nodejs \
    npm \
  && rm -rf /var/lib/apt/lists/*

RUN case "${TARGETARCH}" in \
    amd64) javy_arch="x86_64" ;; \
    arm64) javy_arch="arm" ;; \
    *) echo "unsupported TARGETARCH: ${TARGETARCH}" >&2; exit 1 ;; \
  esac \
  && javy_asset="javy-${javy_arch}-linux-v${JAVY_VERSION}.gz" \
  && curl -fsSL "https://github.com/bytecodealliance/javy/releases/download/v${JAVY_VERSION}/${javy_asset}" -o /tmp/javy.gz \
  && curl -fsSL "https://github.com/bytecodealliance/javy/releases/download/v${JAVY_VERSION}/${javy_asset}.sha256" -o /tmp/javy.sha256 \
  && printf '%s  %s\n' "$(cut -d ' ' -f 1 /tmp/javy.sha256)" /tmp/javy.gz | sha256sum -c - \
  && gzip -dc /tmp/javy.gz > /usr/local/bin/javy \
  && chmod 0755 /usr/local/bin/javy \
  && rm /tmp/javy.gz /tmp/javy.sha256 \
  && javy --version

WORKDIR /build/external

COPY external/package.json external/tsconfig.json external/tsup.config.ts external/command.ts ./
COPY external/controller ./controller
COPY external/entity ./entity
COPY external/repository ./repository
COPY external/share ./share
COPY external/usecase ./usecase

RUN mkdir -p wasm \
  && npm install --no-audit --no-fund \
  && npm run build:pptx-exporter-wasm \
  && test -s wasm/xaligo.wasm

FROM rockylinux:9

ARG GO_VERSION=1.22.12
ARG TARGETARCH

ENV PATH=/usr/local/go/bin:/root/go/bin:${PATH}
ENV PREBUILT_WASM=/opt/xaligo/xaligo.wasm

RUN dnf install -y \
    ca-certificates \
    git \
    gzip \
    rpm-build \
    tar \
  && dnf clean all \
  && rm -rf /var/cache/dnf

RUN case "${TARGETARCH}" in \
    amd64) go_arch="amd64" ;; \
    arm64) go_arch="arm64" ;; \
    *) echo "unsupported TARGETARCH: ${TARGETARCH}" >&2; exit 1 ;; \
  esac \
  && curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${go_arch}.tar.gz" -o /tmp/go.tgz \
  && tar -C /usr/local -xzf /tmp/go.tgz \
  && rm /tmp/go.tgz \
  && go version

COPY --from=wasm-builder /build/external/wasm/xaligo.wasm /opt/xaligo/xaligo.wasm

WORKDIR /workspace
