
https://kro.run/docs/getting-started/Installation

```

❯ export KRO_VERSION=$(curl -sL \
    https://api.github.com/repos/kro-run/kro/releases/latest | \
    jq -r '.tag_name | ltrimstr("v")'
  )

```

```

❯ helm install kro oci://ghcr.io/kro-run/kro/kro \
  --namespace kro \
  --create-namespace \
  --version=${KRO_VERSION}
Pulled: ghcr.io/kro-run/kro/kro:0.2.1
Digest: sha256:2fb565bc9ecfdbd5885047ecc6611924d56ce17129d674cb818f091a818534c9
NAME: kro
LAST DEPLOYED: Sat Mar  1 00:21:23 2025
NAMESPACE: kro
STATUS: deployed
REVISION: 1
TEST SUITE: None

```

```

❯ helm -n kro list
NAME    NAMESPACE       REVISION        UPDATED                                 STATUS          CHART           APP VERSION
kro     kro             1               2025-03-01 00:21:23.651979 +0900 JST    deployed        kro-0.2.1       0.2.1

❯ kubectl get pods -n kro
NAME                  READY   STATUS    RESTARTS   AGE
kro-b667bd485-l55xt   1/1     Running   0          48s

```

https://kro.run/docs/getting-started/deploy-a-resource-graph-definition

`Resource Graph Definition` & `CEL`

```

❯ kubectl apply -f resourcegraphdefinition.yaml
resourcegraphdefinition.kro.run/my-application created

```

```

❯ kubectl get rgd my-application -owide
NAME             APIVERSION   KIND          STATE    TOPOLOGICALORDER                     AGE
my-application   v1alpha1     Application   Active   ["deployment","service","ingress"]   63s

```

```

❯ kubectl apply -f instance.yaml
application.kro.run/my-application-instance created

```

```

❯ kubectl get applications
NAME                      STATE    SYNCED   AGE
my-application-instance   ACTIVE   True     117s

❯ kubectl get deployments my-awesome-app
NAME             READY   UP-TO-DATE   AVAILABLE   AGE
my-awesome-app   3/3     3            3           2m25s


❯ kubectl get services my-awesome-app-service
NAME                     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
my-awesome-app-service   ClusterIP   10.96.189.221   <none>        80/TCP    2m39s


❯ kubectl get ingress my-awesome-app-ingress
NAME                     CLASS    HOSTS   ADDRESS   PORTS   AGE
my-awesome-app-ingress   <none>   *                 80      2m58s


```


