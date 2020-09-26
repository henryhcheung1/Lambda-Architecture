# Lambda Architecture Implementation

This repository is for hands on experience with scalable, fault-tolerant applications deployed to open source big data technologies

## Run Hadoop in Kubernetes
```
helm install --set yarn.nodeManager.resources.limits.memory=4096Mi --set yarn.nodeManager.replicas=1 stable/hadoop --generate-name
```

## Generate go code with Thrift 
```
thrift --gen go model.thrift
```