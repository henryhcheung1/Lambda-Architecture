# Lambda Architecture Implementation

This repository is for hands on experience with scalable, fault-tolerant applications deployed to open source big data technologies

## Run Hadoop in Docker
basic development setup
```
git clone https://github.com/big-data-europe/docker-hadoop.git
cd docker-hadoop
docker-compose up
```

## Generate go code with Thrift 
```
thrift --gen go model.thrift
```