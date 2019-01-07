FROM       quay.io/prometheus/busybox:latest
MAINTAINER Ryota Suginaga <transnano.jp@gmail.com>

COPY local-ssl-exporter /bin/local-ssl-exporter

ENTRYPOINT ["/bin/local-ssl-exporter"]
CMD        ["-version"]
