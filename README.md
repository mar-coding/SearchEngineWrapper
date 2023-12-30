# ElasticSearch Wrapper and API Integration for Search Engine

The project aims to develop a robust search engine wrapper that leverages the power of ElasticSearch as its core component. The wrapper will provide a seamless interface between the ElasticSearch engine and the front-end application through a well-defined API.

## Run
### Backend
on port `8000` Based on REST

How to run:
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

### FrontEnd
on port `8080`

How to run:
```azure
sudo npm install http-server -g
cd others/frontend/search
http-server
```

### ElasticSearch-Kibana(Infra)
on port `9000` and `5601`

it will need a `.env` file like the below template
```azure
# Password for the 'elastic' user (at least 6 characters)
ELASTIC_PASSWORD=changeme

# Password for the 'kibana_system' user (at least 6 characters)
KIBANA_PASSWORD=changeme

# Version of Elastic products
STACK_VERSION=8.8.2

# Set the cluster name
CLUSTER_NAME=docker-cluster

# Set to 'basic' or 'trial' to automatically start the 30-day trial
LICENSE=basic
#LICENSE=trial

# Port to expose Elasticsearch HTTP API to the host
ES_PORT=9200
#ES_PORT=127.0.0.1:9200

# Port to expose Kibana to the host
KIBANA_PORT=5601
#KIBANA_PORT=80

# Increase or decrease based on the available host memory (in bytes)
MEM_LIMIT=1073741824

# Project namespace (defaults to the current folder name if not set)
COMPOSE_PROJECT_NAME=elastic

# Parent Directory for where we want to place our configuration for Elastic
PARENT_DIR=.
```

How to run:
```azure
cd /others/infra
chmod +x setup.sh
./setup.sh
docker-compose up --build -d
```

### Crawler

How to run:
```azure
cd others/crawler
python3 -m venv venv
source venv/bin/activate
pip install -r req.txt
scrapy crawl umScraper
```