## kn service migrate

Migrate Knative services from source cluster to destination cluster

### Synopsis

Migrate Knative services from source cluster to destination cluster

```
kn service migrate [flags]
```

### Examples

```

  # Migrate Knative services from source cluster to destination cluster by export KUBECONFIG and KUBECONFIG_DESTINATION as environment variables
  kn migrate --namespace default --destination-namespace default

  # Migrate Knative services from source cluster to destination cluster by set kubeconfig as parameters
  kn migrate --namespace default --destination-namespace default --kubeconfig $HOME/.kube/config/source-cluster-config.yml --destination-kubeconfig $HOME/.kube/config/destination-cluster-config.yml

  # Migrate Knative services from source cluster to destination cluster and force replace the service if exists in destination cluster
  kn migrate --namespace default --destination-namespace default --force

  # Migrate Knative services from source cluster to destination cluster and delete the service in source cluster
  kn migrate --namespace default --destination-namespace default --force --delete
```

### Options

```
      --delete                          Delete all Knative resources after kn-migration from source cluster
      --destination-kubeconfig string   The kubeconfig of the destination Knative resources (default is KUBECONFIG_DESTINATION from environment variable)
      --destination-namespace string    The namespace of the destination Knative resources
      --force                           Migrate service forcefully, replaces existing service if any.
  -h, --help                            help for migrate
  -n, --namespace string                The namespace of the source Knative resources (default "default")
```

### Options inherited from parent commands

```
      --config string       kn config file (default is $HOME/.kn/config.yaml)
      --kubeconfig string   kubectl config file (default is $HOME/.kube/config)
      --log-http            log http traffic
```

### SEE ALSO

* [kn service](kn_service.md)	 - Service command group

