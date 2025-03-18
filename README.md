# busykube

A busybox style all-in-one binary for Kubernetes. All your favourite Kubernetes tools in a single executable.

This project builds `kubectl`, `helm`, `kind` [and more](https://github.com/jacobtomlinson/busykube/issues/new?template=suggest-new-tool.yml) into a single binary. 

You can use them as subcommands directly from `busykube`.

```console
$ busykube kubectl get pods
No resources found in default namespace.
```

You can also use symlinks to directly invoke subcommands.

```console
$ cd bin && ln -s busykube helm

$ which helm
bin/helm

$ ls -l bin/helm
lrwxrwxrwx 1 jtomlinson 62 Mar 18 15:46 bin/helm -> bin/busykube

$ helm
helm        
The Kubernetes package manager

Common actions for Helm:

- helm search:    search for charts
- helm pull:      download a chart to your local directory to view
- helm install:   upload the chart to Kubernetes
- helm list:      list releases of charts
...
```

## Installation

```bash
go install github.com/jacobtomlinson/busykube@main
```

Busykube can also create all the symlinks for you automatically.

```console
$ busykube install
Created symlink: bin/helm -> bin/busykube
Created symlink: bin/kind -> bin/busykube
Created symlink: bin/kubectl -> bin/busykube
```
