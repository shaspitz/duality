# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
FROM golang:1.18-bullseye as build-env

# install additional OS packages.
RUN apt update && \
    apt upgrade -y

RUN apt-get install -y \
    build-essential \
    ca-certificates \
    # must install cross compiler for arm64
    gcc-aarch64-linux-gnu

WORKDIR /usr/src

# Get Go dependencies
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Copy rest of files
COPY . .

# compile dualityd to ARM64 architecture for final image
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    CGO_ENABLED=1 \
    CC=aarch64-linux-gnu-gcc \
    GOOS=linux \
    GOARCH=arm64 \
    go build -o build/dualityd_arm64 ./cmd/dualityd


# Final image build on small stable release of ARM64 Linux
FROM arm64v8/alpine:20220715

# Install ca-certificates
RUN apk add --update \
    # required for dualityd to work
    libc6-compat \
    # allow JSON parsing in startup shell scripts
    jq \
    # required for HTTPS to connect properly
    ca-certificates

# install TOML editing tool dasel for complicated TOML edits \
RUN wget https://github.com/TomWright/dasel/releases/download/v1.27.3/dasel_linux_arm64.gz; \
    gzip -d dasel_linux_arm64.gz; \
    chmod 755 dasel_linux_arm64; \
    mv ./dasel_linux_arm64 /usr/local/bin/dasel;

WORKDIR /usr/src

# Copy over binaries from the build-env
COPY --from=build-env /usr/src/build/dualityd_arm64 /usr/bin/dualityd

ARG NETWORK=duality-1
# Initialize configuration
RUN dualityd init --chain-id $NETWORK duality
# Add our configuration settings, starting with enabling the API in non-production
RUN dasel put bool -f /root/.duality/config/app.toml ".api.enable" $([[ ! "$NETWORK" =~ "^duality-\d+$" ]] && echo "true" || echo "false"); \
    # add Prometheus telemetry \
    dasel put bool -f /root/.duality/config/app.toml ".telemetry.enable" "true"; \
    dasel put int -f /root/.duality/config/app.toml ".telemetry.prometheus-retention-time" "60"; \
    # todo: the following line has an error \
    # dasel put document -f /root/.duality/config/app.toml -r json -w toml ".telemetry.global-labels" '[["chain_id", "duality"]]'; \
    # ensure listening to the RPC port doesn't block outgoing RPC connections \
    dasel put string -f /root/.duality/config/config.toml ".rpc.laddr" "tcp://0.0.0.0:26657";

# expose ports
# rpc
EXPOSE 26657
# p2p
EXPOSE 26656
# grpc-web
EXPOSE 9091
# grpc
EXPOSE 9090
# prof
EXPOSE 6060
# api
EXPOSE 1317

# add startup scripts and their dependencies
COPY scripts scripts
COPY networks networks

# default to serving the chain with default data and name
CMD ["sh", "./scripts/startup.sh"]
