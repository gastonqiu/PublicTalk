```go
POST /internal/api/request-locks

Request
{
	"idempotency_key": "<string>"
	"lock_duration": "<int>"
}

Response
{
	"lock_value": "<string>"
}

Error
- concurrent_process_error (409)
```


```go
GET /response/{idempoency_key}

HEADER
- X-Request-Fingerprint: "<string>"

Response:
{
    "status_code": "<int>",
    "headers": "<object>",
    "body": "<object>" 
}

Error:
- idempotent_key_not_found (404)
- fingerprint_conflict (422)
```

```go
POST /response

Request
{
    "idempotency_key": "<string>",
    "request_fingerprint": "<string>",
    "response": {
        "headers": "<object>",
        "body": "<object>" ,
        "status_code": "<int>"
    },
    "lock_meta_data": {
        "lock_value": "<string>"
    }
}
```


```go
Method: POST
Resource: /api/v1/payment
Header: 
- X-Idempotency-Key: Uniqueness id (uuid)

Body: 
{
    "amount": 100,
    "currency": "USD"
}
```

```go
func main () {
  server.StartServer(New, Version, Priority)
}

func New() interface{} {
    return &MyConfig{}
}

func (conf *MyConfig) Access (kong *pdk.PDK) {
  ...
}

type PDK struct {
	Client          client.Client
	Ctx             ctx.Ctx
	Log             log.Log
	Nginx           nginx.Nginx
	Request         request.Request
	Response        response.Response
	Router          router.Router
	IP              ip.Ip
	Node            node.Node
	Service         service.Service
	ServiceRequest  service_request.Request
	ServiceResponse service_response.Response
}
```


```go
key = req.header["Idempotency-Key"]
fingerprint = Fingerprint(url+body+whitelisted(headers))
Lock(key)
defer Unlock(key)
cachedResponse = GetFromCache(key, fingerprint)
if cachedResponse != nil {
    return cachedResponse
}
resp = Process(req)
SetResponseToCache(key, resp, fingerprint)
```

```go
// lock
SET resource_name my_random_value NX PX 30000

// unlock
if redis.call("get", key) == value then
    return redis.call("del", key)
else
    return 0
end
```

```dockerfile
# Build Golang plugins

FROM golang:1.20 AS plugin-builder

WORKDIR /builder

COPY ./hello ./go_plugins/hello

RUN find ./go_plugins -maxdepth 1 -mindepth 1 -type d -not -path "*/.git*" | \
    while read dir; do \
        cd $dir && go build -o /builds/$dir main.go  ; \
    done

# Build Kong
FROM kong:3.4.0-ubuntu

COPY --from=plugin-builder ./builds/go_plugins/  ./kong/

USER kong
```

```shell
KONG_PLUGINS: "bundled,hello"
KONG_PLUGINSERVER_NAMES: hello
KONG_PLUGINSERVER_HELLO_START_CMD: /kong/hello
KONG_PLUGINSERVER_HELLO_QUERY_CMD: /kong/hello -dump
```


```yaml
plugins:
    - config:
        replace:
          uri: /service/test
      enabled: true
      name: request-transformer
    - config:
        message: HELLO!
      enabled: true
      name: hello
```

```yaml
env:
    plugins: “bundle,hello"

dblessConfig:
    configMap: ""
    config: |
        _format_version: "3.0"
        _transform: true

        services:
        - name: my-service
        url: https://example.com
        plugins:
        - name: hello
        routes:
        - name: my-route
        paths:
        - /
```