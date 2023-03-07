FROM scratch

ADD v2ray-api ./v2ray-api
ADD v2ray ./v2ray
ADD server.json ./server.json

EXPOSE 10101

# CD to /
WORKDIR /

ENTRYPOINT ["./v2ray-api"]
