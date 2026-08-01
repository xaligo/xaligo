FROM node:24-bookworm-slim AS node-runtime

FROM ubuntu:24.04

ARG GO_VERSION=1.26.5
ARG RUST_VERSION=1.85.1
ARG JAVY_VERSION=9.0.0
ARG TARGETARCH

ENV DEBIAN_FRONTEND=noninteractive
ENV PATH=/usr/local/go/bin:/root/go/bin:/root/.cargo/bin:${PATH}

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
    ca-certificates \
    build-essential \
    curl \
    dpkg-dev \
    git \
    gzip \
    make \
    tar \
  && rm -rf /var/lib/apt/lists/*

COPY --from=node-runtime /usr/local/ /usr/local/

RUN case "${TARGETARCH}" in \
    amd64) go_arch="amd64" ;; \
    arm64) go_arch="arm64" ;; \
    *) echo "unsupported TARGETARCH: ${TARGETARCH}" >&2; exit 1 ;; \
  esac \
  && curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${go_arch}.tar.gz" -o /tmp/go.tgz \
  && tar -C /usr/local -xzf /tmp/go.tgz \
  && rm /tmp/go.tgz \
  && go version

RUN curl --proto '=https' --tlsv1.2 -fsSL https://sh.rustup.rs -o /tmp/rustup-init.sh \
  && sh /tmp/rustup-init.sh -y --profile minimal --default-toolchain "${RUST_VERSION}" \
  && rm /tmp/rustup-init.sh \
  && rustc --version \
  && cargo --version

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

WORKDIR /workspace
