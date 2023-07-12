# kustomize-test
kustomize version testing

## Download docker images

```bash
docker pull registry.k8s.io/kustomize/kustomize:v4.5.7
docker pull registry.k8s.io/kustomize/kustomize:v5.0.1
```

## Build

Using go version: `go version go1.20.4 darwin/amd64`
`go build` cmd will output the `testing-kustomize` executable

## Launch

Launch the executable:
```bash
./testing-kustomize --input-dir <path-to-folder>
```
*Note: <path-to-folder> must be the parent folder containing kustomize directories*