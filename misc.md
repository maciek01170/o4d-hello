# Openshift local (crc)

## configuring instance

```bash
# set 6 cpus, 16GB of RAM, 64GB of disk
crc stop
crc config set cpus 6
crc config set memory 16384
crc config set disk-size 64 
crc config set enable-cluster-monitoring true # doesn't not work well
crc start
```

### k8s misc

Switch default namespace
```bash
kubectl config set-context --current --namespace=<namespace-name>
```

## app from git sources (use oc src builer)

Repo: https://github.com/maciek01170/o4d-hello.git

1. Create project (e.g. hello-go-src)
2. **Topology**: 
   1. Add page -> import from git
   2. Application name: hello-go-src-app, 
   3. Name: hello-go-src
   4. Set **low** resource limits (e.g. cpu 10m, memory 64MiB)
   5. Create
3. 
