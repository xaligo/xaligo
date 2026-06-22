FROM ubuntu:24.04

ARG GO_VERSION=1.22.12
ARG TARGETARCH

ENV DEBIAN_FRONTEND=noninteractive
ENV PATH=/usr/local/go/bin:/root/go/bin:${PATH}

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    dpkg-dev \
    git \
    make \
    tar \
  && rm -rf /var/lib/apt/lists/*

RUN case "${TARGETARCH}" in \
    amd64) go_arch="amd64" ;; \
    arm64) go_arch="arm64" ;; \
    *) echo "unsupported TARGETARCH: ${TARGETARCH}" >&2; exit 1 ;; \
  esac \
  && curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${go_arch}.tar.gz" -o /tmp/go.tgz \
  && tar -C /usr/local -xzf /tmp/go.tgz \
  && rm /tmp/go.tgz \
  && go version

WORKDIR /workspace