FROM rockylinux:9

ARG GO_VERSION=1.22.12
ARG JAVY_VERSION=9.0.0
ARG TARGETARCH

ENV PATH=/usr/local/go/bin:/root/go/bin:${PATH}

RUN dnf module enable -y nodejs:20 \
  && dnf install -y \
    ca-certificates \
    git \
    gzip \
    make \
    nodejs \
    npm \
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
