# v2ray-api
v2ray server controlled by an API

## Setup

`git clone https://github.com/v2fly/v2ray-core`
`cd v2ray-core/main`
`export CGO_ENABLED=0`
`go build`
`mv main /path/to/this/repo/v2ray`
`cd /path/to/this/repo`
`go build`
`docker build -t v2ray .`
`docker run v2ray:latest`
