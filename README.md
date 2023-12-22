# ElasticSearch Wrapper and API Integration for Search Engine

The project aims to develop a robust search engine wrapper that leverages the power of ElasticSearch as its core component. The wrapper will provide a seamless interface between the ElasticSearch engine and the front-end application through a well-defined API.

### Run
```azure
git clone https://github.com/mar-coding/SearchEngineWrapper.git
cd SearchEngineWrapper
go run cmd/main.go run -c configs/config-dev.yml
```
The template for `config.yml` file

```yaml
address: "0.0.0.0"
domain: "url"
development: true

grpc:
  port: port

rest:
  port: port

database:
  elastic:
    addresses:
      - uri
    username: elastic
    password: password_for_elastic

logging:
  debug: false
  handler: 0
  enable_caller: true

extra_data:
  field_name: "value"
```
