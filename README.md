# Lambda Architecture Implementation

This repository is for hands on experience with scalable, fault-tolerant applications deployed to open source big data technologies

## Work In Progress

## Run Hadoop in Docker
basic dev setup
```
git clone https://github.com/big-data-europe/docker-hadoop.git
cd docker-hadoop
docker-compose up
```

## Generate go code with Thrift 
```
thrift --gen go model.thrift
```

## TODO:

- Dockerize Go main app & run in k8s because it appears HDFS Go client needs to be in same network space as HDFS docker containers