FROM iron/go:dev
WORKDIR /app
#
ENV SRC_DIR=/home/srini/gows/src/github.com/muly/research/golang/docker/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/
ENTRYPOINT ["./myapp"]