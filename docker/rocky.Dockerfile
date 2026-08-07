FROM rockylinux:9

ARG GO_VERSION=1.26.5
ARG RUST_VERSION=1.85.1
ARG TARGETARCH

ENV PATH=/usr/local/go/bin:/root/go/bin:/root/.cargo/bin:${PATH}

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

WORKDIR /workspace
