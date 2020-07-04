package main

import data.kubernetes

deploy := "Deployment"

name = input.metadata.name

# use '#' for comments
warn[msg] {
  input.kind == deploy
  not input.spec.template.spec.securityContext.runAsNonRoot

  msg = sprintf("Containers must not run as root in Deployment %s", [name])
}

required_deployment_selectors {
  input.spec.selector.matchLabels.app
  input.spec.selector.matchLabels.release
}

deny_deployment_with_no_app_label_selector[msg] {
  kubernetes.is_deployment
  not required_deployment_selectors

  msg = sprintf("Deployment %s must provide app/release labels for pod selectors", [name])
}

required_deployment_labels {
    input.metadata.labels["app.kubernetes.io/name"]
    input.metadata.labels["app.kubernetes.io/instance"]
    input.metadata.labels["app.kubernetes.io/version"]
    input.metadata.labels["app.kubernetes.io/component"]
    input.metadata.labels["app.kubernetes.io/part-of"]
    input.metadata.labels["app.kubernetes.io/managed-by"]
}

deny[msg] {
  kubernetes.is_deployment
  not required_deployment_labels
  msg = sprintf("%s must include Kubernetes recommended labels: https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/#labels", [name])
}

deny_statefulsets[msg] {
  kubernetes.is_statefulset
  msg = sprintf("Found statefulset %s but statefulsets are not allowed", [name])
}

deprecated_deployment_version = [
  "extensions/v1beta1",
  "apps/v1beta1",
  "apps/v1beta2"
]

warn[msg] {
  input.kind == deploy
  input.apiVersion == deprecated_deployment_version[i]
  msg = "最新の APIVersion apps/v1 を指定してください"
}

deny[msg] {
  input.kind == deploy
  input.spec.template.spec.containers[_].securityContext.privileged == true
  msg = "privileged はセキュリティ上の理由で許可されていません"
}
