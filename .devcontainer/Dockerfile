FROM mcr.microsoft.com/vscode/devcontainers/base:debian-11

# install fish terminal
RUN apt-get update && apt-get install -y \
    fish \
    && rm -rf /var/lib/apt/lists/*

# install golang
ARG GO_VERSION=1.20.1
RUN set -x \
    && curl -fsSL "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -xz -C /usr/local
ENV PATH $PATH:/usr/local/go/bin
ENV PATH $PATH:/home/vscode/go/bin

# install go tools (for golang.go extention)
# cgo depends on gcc
RUN apt-get update && apt-get install -y \
    gcc \
    && rm -rf /var/lib/apt/lists/*
ENV GOPATH=/home/vscode/go
USER vscode
RUN    go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/cweill/gotests/gotests@latest \
    && go install -v github.com/fatih/gomodifytags@latest \
    && go install -v github.com/josharian/impl@latest \
    && go install -v github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v honnef.co/go/tools/cmd/staticcheck@latest
