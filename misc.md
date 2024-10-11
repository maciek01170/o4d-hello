# Openshift local (crc)

## from git sources (use oc src builer)

Repo: git@github.com:maciek01170/o4d-hello.git

1. Create project (e.g. hello-go-src)
2. **Topology**: 
   1. Add page -> import from git
   2. Application name: hello-go-src-app, 
   3. Name: hello-go-src
   2. Set **low** resource limits (e.g. cpu 10m, memory 64MiB)
   3. Create
   4. *crc crashed* -> restart
5. 
