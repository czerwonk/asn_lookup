# asn_lookup
[![Build Status](https://travis-ci.org/czerwonk/asn_lookup.svg)](https://travis-ci.org/czerwonk/asn_lookup)
[![Go Report Card](https://goreportcard.com/badge/github.com/czerwonk/asn_lookup)](https://goreportcard.com/report/github.com/czerwonk/asn_lookup)

Webservice providing prefixes announced by a specific ASN based on routing daemon information

## Default port
By default the service listens on port 10179 for incoming requests

## Installation
```
go get -u github.com/czerwonk/asn_lookup
```

## Running
In this exaple the socket is located at /usr/local/var/run/. This can vary in different setups:
```
asn_lookup -bird.socket=/usr/local/var/run/bird.ctl -bird6.socket=/usr/local/var/run/bird6.ctl
```

## Usage (Example)
### Request
```
curl http://[::]:10179/asn/15169
```

### Response:
```
<AutonomousSystem><ASN>15169</ASN><IPv6>2404:6800:4008::/48</IPv6><IPv6>2404:6800:4009::/48</IPv6><IPv6>2404:6800:400a::/48</IPv6><IPv6>2620:15c::/36</IPv6><IPv6>2404:6800:4001::/48</IPv6><IPv6>2404:6800:4002::/48</IPv6><IPv6>2404:6800:4003::/48</IPv6><IPv6>2404:6800:4004::/48</IPv6><IPv6>2605:ef80:c0::/42</IPv6><IPv6>2404:6800:4005::/48</IPv6><IPv6>2404:6800:4006::/48</IPv6><IPv6>2404:6800:4007::/48</IPv6><IPv6>2620:0:1000::/40</IPv6><IPv6>2800:3f0::/32</IPv6><IPv6>2a00:1450::/32</IPv6><IPv6>2600:1900::/28</IPv6><IPv6>2404:6800::/32</IPv6><IPv6>2605:ef80:80::/42</IPv6><IPv6>2a00:79e0::/32</IPv6><IPv6>2a00:1450:401c::/48</IPv6><IPv6>2a00:1450:401a::/48</IPv6><IPv6>2a00:1450:401b::/48</IPv6><IPv6>2a00:1450:4018::/48</IPv6><IPv6>2a00:1450:4019::/48</IPv6><IPv6>2a00:1450:4016::/48</IPv6><IPv6>2800:3f0:4003::/48</IPv6><IPv6>2001:4860::/32</IPv6><IPv6>2a00:1450:4017::/48</IPv6><IPv6>2800:3f0:4002::/48</IPv6><IPv6>2a00:1450:4014::/48</IPv6><IPv6>2800:3f0:4001::/48</IPv6><IPv6>2a00:1450:4015::/48</IPv6><IPv6>2a00:1450:4012::/48</IPv6><IPv6>2a00:1450:4013::/48</IPv6><IPv6>2a00:1450:4010::/48</IPv6><IPv6>2800:3f0:4005::/48</IPv6><IPv6>2a00:1450:4011::/48</IPv6><IPv6>2800:3f0:4004::/48</IPv6><IPv6>2a00:1450:400e::/48</IPv6><IPv6>2a00:1450:400f::/48</IPv6><IPv6>2a00:1450:400c::/48</IPv6><IPv6>2a00:1450:400d::/48</IPv6><IPv6>2a00:1450:400a::/48</IPv6><IPv6>2a00:1450:400b::/48</IPv6><IPv6>2a00:1450:4008::/48</IPv6><IPv6>2a00:1450:4009::/48</IPv6><IPv6>2a00:1450:4006::/48</IPv6><IPv6>2a00:1450:4007::/48</IPv6><IPv6>2a00:1450:4004::/48</IPv6><IPv6>2a00:1450:4005::/48</IPv6><IPv6>2a00:1450:4002::/48</IPv6><IPv6>2a00:1450:4003::/48</IPv6><IPv6>2a00:1450:4001::/48</IPv6><IPv6>2a03:ace0:100::/40</IPv6><IPv6>2a03:ace0::/32</IPv6><IPv6>2607:f8b0:4011::/48</IPv6><IPv6>2607:f8b0:4010::/48</IPv6><IPv6>2607:f8b0:4013::/48</IPv6><IPv6>2607:f8b0:4012::/48</IPv6><IPv6>2605:ef80:40::/42</IPv6><IPv6>2607:f8b0:4015::/48</IPv6><IPv6>2605:ef80:140::/42</IPv6><IPv6>2607:f8b0:4014::/48</IPv6><IPv6>2607:f8b0:4016::/48</IPv6><IPv6>2607:f8b0:4009::/48</IPv6><IPv6>2607:f8b0:4008::/48</IPv6><IPv6>2607:f8b0:400b::/48</IPv6><IPv6>2607:f8b0:400a::/48</IPv6><IPv6>2607:f8b0:400d::/48</IPv6><IPv6>2607:f8b0:400c::/48</IPv6><IPv6>2607:f8b0:400f::/48</IPv6><IPv6>2607:f8b0:400e::/48</IPv6><IPv6>2607:f8b0:4001::/48</IPv6><IPv6>2607:f8b0:4000::/48</IPv6><IPv6>2607:f8b0:4003::/48</IPv6><IPv6>2c0f:fb50:4002::/48</IPv6><IPv6>2607:f8b0:4002::/48</IPv6><IPv6>2c0f:fb50:4003::/48</IPv6><IPv6>2607:f8b0:4005::/48</IPv6><IPv6>2607:f8b0:4004::/48</IPv6><IPv6>2607:f8b0:4007::/48</IPv6><IPv6>2607:f8b0:4006::/48</IPv6><IPv6>2605:ef80::/32</IPv6><IPv6>2605:ef80:200::/42</IPv6><IPv6>2c0f:fb50::/32</IPv6><IPv6>2607:f8b0::/32</IPv6><IPv6>2401:fa00::/32</IPv6><IPv6>2620:120:e000::/40</IPv6></AutonomousSystem>
```

## Supported routing daemons
* [Bird](http://bird.network.cz/)

## License
(c) Daniel Czerwonk, 2017. Licensed under [MIT](LICENSE) license.

