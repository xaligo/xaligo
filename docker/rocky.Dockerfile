FROM rust:1.85.1-bookworm AS wasm-builder

RUN rustup target add wasm32-wasip1

WORKDIR /build

COPY external/exporter/Cargo.toml external/exporter/Cargo.lock ./external/exporter/
COPY external/exporter/src ./external/exporter/src

RUN cargo build --manifest-path external/exporter/Cargo.toml --package xaligo-pptx-exporter --bin xaligo-exporter --target wasm32-wasip1 --release --locked \
  && test -s external/exporter/target/wasm32-wasip1/release/xaligo-exporter.wasm

FROM rockylinux:9

ARG GO_VERSION=1.26.5
ARG RUST_VERSION=1.85.1
ARG TARGETARCH

ENV PATH=/usr/local/go/bin:/root/go/bin:/root/.cargo/bin:${PATH}
ENV PREBUILT_WASM=/opt/xaligo/xaligo.wasm

RUN dnf install -y \
    ca-certificates \
    curl \
    gcc \
    gcc-c++ \
    git \
    glibc-devel \
    gzip \
    make \
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

RUN curl --proto '=https' --tlsv1.2 -fsSL https://sh.rustup.rs -o /tmp/rustup-init.sh \
  && sh /tmp/rustup-init.sh -y --profile minimal --default-toolchain "${RUST_VERSION}" \
  && rm /tmp/rustup-init.sh \
  && rustc --version \
  && cargo --version

COPY --from=wasm-builder /build/external/exporter/target/wasm32-wasip1/release/xaligo-exporter.wasm /opt/xaligo/xaligo.wasm

WORKDIR /workspace
