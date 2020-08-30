# local ssl exporter ![Releases](https://github.com/transnano/local-ssl-exporter/workflows/Releases/badge.svg) ![Publish Docker image](https://github.com/transnano/local-ssl-exporter/workflows/Publish%20Docker%20image/badge.svg) ![Vulnerability Scan](https://github.com/transnano/local-ssl-exporter/workflows/Vulnerability%20Scan/badge.svg)

![License](https://img.shields.io/github/license/transnano/local-ssl-exporter?style=flat)

![Container image version](https://img.shields.io/docker/v/transnano/local-ssl-exporter?style=flat)
![Container image size](https://img.shields.io/docker/image-size/transnano/local-ssl-exporter/latest?style=flat)
![Container image pulls](https://img.shields.io/docker/pulls/transnano/local-ssl-exporter?style=flat)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/transnano/local-ssl-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/transnano/local-ssl-exporter)](https://goreportcard.com/report/github.com/transnano/local-ssl-exporter)

Check the deadline of server certificate on local.

## How to use

``` shell
$ local-ssl-exporter \
  -files=dummy1.crt,dummy2.crt \
  -out=/path/to/local_ssl_exporter.prom
```

- `-files`: Specify certificate files with comma separated values.
- `-out`: Output prom-formatted file path.

## Output

```
local_ssl_unixtime{file="dummy1.crt"} 1595516400
local_ssl_sub_days{file="dummy1.crt"} 570
local_ssl_unixtime{file="dummy2.crt"} 1598281200
local_ssl_sub_days{file="dummy2.crt"} 602
```

- `local_ssl_unixtime`: Represent the deadline of server certificate in Unixtime.
- `local_ssl_sub_days`: Represents the difference between the deadline of the server certificate and the current date-time.
