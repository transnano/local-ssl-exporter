# local ssl exporter

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
