# CMDB

云资产管理平台



## API使用

使用Postman或者curl进行调试：

### API Root

显示所有支持的API接口

```sh
curl http://localhost:8050/
```

### 示例

- 查询单个应用配置项

```
curl 'http://localhost:8050/cmdb/api/v1/configs/c668p08i9qr13q5pnle0'
```

- 查看所有应用配置项列表（带分页）

```
curl 'http://localhost:8050/cmdb/api/v1/configs?page_size=10&page_number=1'
```

- 创建应用配置项

```sh
curl -X POST 'http://localhost:8050/cmdb/api/v1/configs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "describe": {
        "application_id": "c65pksgi9qr1qv4ogokg",
        "name": "webhook.mongodb.host",
        "host": "127.0.0.1",
        "port": 8080,
        "env": "test",
        "type": "mongodb",
        "source": "gitlab"
    }
}'
```

- 删除应用配置项

```
curl -X DELETE 'http://localhost:8050/cmdb/api/v1/configs/c668p08i9qr13q5pnle0'
```

- 关键字搜索应用配置项

```
curl 'http://localhost:8050/cmdb/api/v1/configs?keywords=webhook'
```



## SDK使用

```go
package main
import (
  "context"
  "fmt"
  "gitlab.yewifi.com/DevOps/cmdb"
  "gitlab.yewifi.com/DevOps/cmdb/app/resource"
)
```

```
func main() {
  // 配置cmdb grpc服务调用地址和凭证
  conf := client.NewConfig("localhost:18060")
  // 创建CMDB客户端
  cmdb, err := client.NewClient(conf)
  if err != nil {
​    panic(err)
  }

  // 服务调用
  rs, err := cmdb.Resource().Search(context.Background(), resource.NewSearchRequest())
  if err != nil {
​    panic(err)
  }
  fmt.Println(rs)
}
```



## 部署

###  依赖

- MySQL

  

### 配置

./etc/cmdb.toml

```
[app]
name = "cmdb"
http_host = "0.0.0.0"
http_port = "8060"
grpc_host = "0.0.0.0"
grpc_port = "18060"

[mysql]
host = "xxx"
port = "xxxx"
username = "cmdb"
password = "xxxx"
database = "cmdb"

[log]
level = "debug"
path = "logs"
format = "text"
to = "stdout"
```


本地开发的启动方式:

```
go run main.go start -f etc/cmdb.toml
```

or

```
make run
```

### 编译

```
make linux
```

### 运行

```
nohup cmdb start -f /etc/cmdb.toml &> cmdb.log &
```

