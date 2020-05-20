FROM scratch

WORKDIR $GOPATH/src/gin-monitor
COPY . $GOPATH/src/gin-monitor

EXPOSE 9090
CMD ["./run"]