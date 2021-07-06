Example of using [Werf](https://werf.io) to deploy backend in Go and front-end SPA.

## Local Cluster

```shell
# prefix `k3d-` will be automatically added to registry name, 
# don't forget to add `127.0.0.1 k3d-registry.localhost` to /etc/hosts
k3d registry create registry.localhost --port 5000
k3d cluster create local --registry-use k3d-registry.localhost:5000 --port "8080:80@loadbalancer"
kubectl config use-context k3d-local
```

Please note — only k3d 5.x is able to edit port configuration on the fly, 4.x requires cluster recreation.

## Developer Environment

Export envs for local development. Don't forget to change `werf-spa-go-example-dev` to your app name.

```shell
export WERF_REPO=k3d-registry.localhost:5000/werf-spa-go-example-dev
export WERF_INSECURE_REGISTRY=1
export WERF_ENV=dev
```

See [Environment](https://werf.io/documentation/v1.2/advanced/helm/configuration/templates.html#environment) docs about `WERF_ENV=dev`.