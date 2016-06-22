FROM golang:1.6-onbuild

EXPOSE 9090
ENTRYPOINT ["/go/bin/app"]