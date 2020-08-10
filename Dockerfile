FROM golang:1.14 AS builder

# # Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/APoniatowski/routes
RUN git clone https://github.com/APoniatowski/routes.git .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /routes .

FROM scratch
COPY --from=builder /routes ./
ENTRYPOINT ["./routes"]



## copy/pasta from kubesails examples. Will edit if this does not work