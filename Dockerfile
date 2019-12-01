# Stage 1: Build
FROM golang:latest as builder
COPY . /app
WORKDIR /app/server
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o Recommender

RUN apt-get update && apt-get install -y \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

# install UPX for executable compression
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

RUN strip --strip-unneeded Recommender
RUN upx Recommender

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
CMD ["./Recommender"]
