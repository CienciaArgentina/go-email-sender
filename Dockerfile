FROM golang:1.13.6 as builder
# install xz
RUN apt-get update && apt-get install -y \
    xz-utils \
&& rm -rf /var/lib/apt/lists/*
# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

# create a working directory
WORKDIR /go/src/github.com/CienciaArgentina/go-email-sender/
ADD . /go/src/github.com/CienciaArgentina/go-email-sender/

COPY go.mod go.sum ./
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main
# strip and compress the binary
RUN strip --strip-unneeded main
RUN upx main

# use scratch (base for a docker image)
FROM scratch
# set working directory
WORKDIR /root
# copy the binary from builder
COPY --from=builder /go/src/github.com/CienciaArgentina/go-email-sender/ .
# run the binary
CMD ["./main"]