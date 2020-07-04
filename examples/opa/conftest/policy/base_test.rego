package main

empty(value) {
  count(value) == 0
}

no_violations {
  empty(deny)
}

test_deployment_with_security_context {
  input := {
    "kind": "Deployment",
    "metadata": {
      "name": "sample",
      "labels": {
        "app.kubernetes.io/name": "name",
        "app.kubernetes.io/instance": "instance",
        "app.kubernetes.io/version": "version",
        "app.kubernetes.io/component": "component",
        "app.kubernetes.io/part-of": "part-of",
        "app.kubernetes.io/managed-by": "managed-by"
      }
    },
    "spec": {
      "selector": {
        "matchLabels": {
          "app": "app",
          "release": "release"
        }
      },
      "template": {
        "spec": {
          "securityContext": {
            "runAsNonRoot": true
          }
        }
      }
    }
  }

  no_violations with input as input
}

test_deprecated_deployment {
  warn["最新の APIVersion apps/v1 を指定してください"] with input as {"kind": "Deployment", "apiVersion": "extensions/v1beta1"}
}

test_deprecated_deployment {
  warn["最新の APIVersion apps/v1 を指定してください"] with input as {"kind": "Deployment", "apiVersion": "apps/v1beta1"}
}

test_deprecated_deployment {
  warn["最新の APIVersion apps/v1 を指定してください"] with input as {"kind": "Deployment", "apiVersion": "apps/v1beta2"}
}

test_deprecated_deployment_allowed {
  # 最新の API Version の場合は warn が出ない
  not warn["最新の APIVersion apps/v1 を指定してください"] with input as {"kind": "Deployment", "apiVersion": "apps/v1"}
}

test_privileged_deployment {
  deny["privileged はセキュリティ上の理由で許可されていません"] with input as
  {
    "kind": "Deployment",
    "spec": {
      "template": {
        "spec": {
          "containers": [
            {
              "name": "app1",
              "image": "nginx",
            },
            {
              "name": "app2",
              "image": "nginx",
              "securityContext": {
                "privileged": true,
              }
            }
          ],
        }
      }
    }
  }
}

test_privileged_deployment_allowed {
  not deny["privileged はセキュリティ上の理由で許可されていません"] with input as
  {
    "kind": "Deployment",
    "spec": {
      "template": {
        "spec": {
          "containers": [
            {
              "name": "app1",
              "image": "nginx",
            },
            {
              "name": "app2",
              "image": "nginx",
            }
          ],
        }
      }
    }
  }
}


