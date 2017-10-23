# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
ENV SRC_DIR=/go/src/polosate/dockergo/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR && go build -o myapp && cp myapp /app/
RUN chmod +x ./myapp
ENTRYPOINT ["./myapp"]