# Balerter Core API Go Library

[Balerter Core API documentation](https://balerter.com/core-api/about)

## Example

Create config file for Balerter `config.hcl`

```hcl
datasources {
  postgres "pg1" {
    host     = "postgres"
    port     = 5432
    username = "postgres"
    password = "secret"
    database = "db"
    sslMode  = "disable"
  }
}

channels {
  log "lg1" {}
}

api {
  coreApi {
    address = ":2020"
  }
}
```

Up the balerter and postgres with docker-compose

```yaml
version: '3'
services:
  postgres:
    image: postgres:9.6.22-stretch
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: db
  balerter:
    image: balerter/balerter
    volumes:
      - ./config.hcl:/opt/config.hcl
    command:
      - "-config=/opt/config.hcl"
    ports:
      - "2020:2020"
    depends_on:
      - postgres
```

Create and run go file with core api library usage

> We pass errors check in this example. Off course, you should check errors in production.

```go
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	coreapi "github.com/balerter/coreapi-go"
)

type serviceInfo struct {
	Name string `json:"name"`
	RPS  int    `json:"rps"`
}

func main() {
	api := coreapi.New("http://localhost:2020", "")

	// Run check every 5 seconds
	t := time.NewTicker(time.Second * 5)

	for {
		// random rps value 35-45
		randonRPSValue := strconv.Itoa(rand.Intn(10) + 35)

		resp, _ := api.Datasource.Postgres("pg1").Query("SELECT 'Cache' AS name, " + randonRPSValue + " AS rps")

		var info []serviceInfo
		json.Unmarshal(resp, &info)

		fmt.Printf("%+v\n", info)

		if info[0].RPS > 40 {
			api.Alert.Error("alert-id", fmt.Sprintf("service %s has too high rps: %d", info[0].Name, info[0].RPS), nil)
		} else {
			api.Alert.Success("alert-id", fmt.Sprintf("service %s rps is ok: %d", info[0].Name, info[0].RPS), nil)
		}

		<-t.C
	}
}
```

Balerter log output after some iterations

```
balerter_1  | {"level":"info","ts":1660573586.675884,"msg":"Log channel message","channel name":"lg1","alert id":"alert-id","level":"error","message":"service Cache has too high rps: 42","image":"","fields":null}
balerter_1  | {"level":"info","ts":1660573601.6705568,"msg":"Log channel message","channel name":"lg1","alert id":"alert-id","level":"success","message":"service Cache rps is ok: 36","image":"","fields":null}
```

## API

Information about use core api modules you can find on [balerter site](https://balerter.com)

### Create new instance

```go
api := coreapi.New("<balerter core api address>", "<auth token>")
```

You should run balerter with Core API enabled. See example above or balerter documentation.

If you not configure Core API auth token in balerter config file, you can use `New` function with empty auth token.

### Modules

#### Alert

```go
api.Alert.Success(alertName, message string, opts *AlertOptions) (*Alert, bool, error)
api.Alert.Warning(alertName, message string, opts *AlertOptions) (*Alert, bool, error)
api.Alert.Error(alertName, message string, opts *AlertOptions) (*Alert, bool, error)
```

#### TLS

```go
api.TLS.Get(hostname string) ([]TLSResult, error)
```

#### Datasource

```go
// Postgres
api.Datasource.Postgres(name string).Query(query string) ([]byte, error)

// MySQL
api.Datasource.MySQL(name string).Query(query string) ([]byte, error)

// Clickhouse
api.Datasource.Clickhouse(name string).Query(query string) ([]byte, error)

// Loki
api.Datasource.Loki(name string).Query(query string, params *LokiQueryParams) ([]byte, error)
api.Datasource.Loki(name string).Range(query string, params *LokiRangeParams) ([]byte, error)

// Prometheus
api.Datasource.Prometheus(name string).Query(query string, params *PrometheusQueryParams) ([]byte, error)
api.Datasource.Prometheus(name string).Range(query string, params *PrometheusRangeParams) ([]byte, error)
```

#### KV

```go
api.KV.Put(key, value string) error
api.KV.Upsert(key, value string) error
api.KV.Delete(key string) error
api.KV.Get(key string) (string, error)
api.KV.All() (map[string]string, error)
```

#### Log

```go
api.Log.Error(message string) error
api.Log.Warn(message string) error
api.Log.Info(message string) error
api.Log.Debug(message string) error
```

#### Runtime

```go
api.Runtime.Get() (*RuntimeInfo, error) 
```

#### Chart

```go
api.Chart.Render(title string, series []DataSeries) ([]byte, error)
```