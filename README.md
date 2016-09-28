# respoke-webhook-tester-golang

Test Respoke webhooks with a simple server.

First, go to the [Respoke developer portal](https://portal.respoke.io)
and add a webhook endpoint for your app.

## Running

Download a [release](https://github.com/ruffrey/respoke-webhook-tester-golang/releases/latest) for your platform.

The first argument should be the port for the server to listen on.

```
unzip <my platform and architecture>-rhook.zip
./rhook 4000
```

## Running / developing with golang


```
go run rhook.go 4000
```
