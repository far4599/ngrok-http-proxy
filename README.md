# ngrok-http-proxy
Docker image to easily start http proxy on the host with no direct ip address to connect. Ngrok will provide an ip.

# How to run
Start a docker container on the host that you want to use as a proxy server.
It will connect to ngrok using you auth token and if successfully connected you will see on the screen the address you may use to connect to your HTTP Proxy.

You will get an auth token for free use when registered at https://dashboard.ngrok.com/signup

```shell
docker run -t -e NGROK_AUTHTOKEN=... far4599/ngrok-http-proxy:latest
```

Look for the line starting with "Forwarding". And use address starting with "tcp://" as an address for your HTTP Proxy.
![ngrok output](https://user-images.githubusercontent.com/4191145/175918164-3f05473e-ccda-4325-9bf1-e520594683db.png)