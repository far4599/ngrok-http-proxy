# ngrok-http-proxy
Docker image to easily start an HTTP proxy server on the host with no direct ip address to connect. Ngrok will provide a public ip for FREE.

# How to run
Start a docker container on the host that you want to use as a proxy server.
It will connect to Ngrok using you auth token and if successfully connected you will see the address of your HTTP Proxy in container logs.

You can find your auth token at https://dashboard.ngrok.com/get-started/your-authtoken. Ngrok registration and service is free.

```shell
$ docker run -t -e NGROK_AUTHTOKEN=... far4599/ngrok-http-proxy:latest
HTTP proxy listens at '0.tcp.ap.ngrok.io:15363'
```