#!/bin/bash

set -euxo

# See docs on how to make
KUBE_NAMESPACE=mini-stage1

#kubectl label namespace $KUBE_NAMESPACE istio-injection=enabled



GOPDF_BINARY_HASH=$(openssl dgst -md5 -binary gopdf  | xxd -pb)

cat > deployment.yaml <<EOF
---
# Inspiration from Kelsey Hightower hello world
# See https://github.com/kelseyhightower/kubernetes-initializer-tutorial/blob/master/deployments/helloworld.yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  labels:
    app: gopdf
  name: gopdf
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: gopdf
        metrics: gopdf-custom-metrics
      name: gopdf
      annotations:
        sidecar.istio.io/inject: "true"
        prometheus.io/path: /metrics
        prometheus.io/port: "8081"
        prometheus.io/scrape: "true"
    spec:
      containers:
      - name: gopdf
        image: repository.hybris.com:5005/upsurge/gopdf:$GOPDF_BINARY_HASH
        imagePullPolicy: Always
        ports:
          - containerPort: 8080
          - containerPort: 8081

      imagePullSecrets:
        - name: "hybris-artifactory"

---
kind: Service
apiVersion: v1
metadata:
  name: gopdf
spec:
  selector:
    app: gopdf
  ports:
  - protocol: TCP
    port: 8080
---
kind: Service
apiVersion: v1
metadata:
  name: gopdf-metrics-8081
  labels:
    k8s-app: metrics
    metrics: gopdf-custom-metrics
spec:
  selector:
    app: gopdf
  ports:
  - name: web
    port: 8081
    protocol: TCP
---
EOF


kubectl -n $KUBE_NAMESPACE apply -f deployment.yaml


######
# This doesn't work yet.
# I'm getting permission errors when I try to create a ServiceMonitor
# which will automatically collect metrics from our code.
#
# + kubectl -n mini-stage1 apply -f -
# Error from server (Forbidden): error when retrieving current configuration of:
# &{0xc42009db00 0xc42035ce00 mini-stage1 metrics STDIN 0xc421e49bd8  false}
# from server for: "STDIN": Forbidden (user=jonathan.hess@sap.com, verb=get, resource=servicemonitors, subresource=) (get servicemonitors.monitoring.coreos.com metrics)

KUBE_NAMESPACE=mini-stage1

#kubectl -n $KUBE_NAMESPACE apply -f - <<EOF
#apiVersion: monitoring.coreos.com/v1
#kind: ServiceMonitor
#metadata:
#  name: metrics
#  namespace: $KUBE_NAMESPACE
#  labels:
#    prometheus: core
#    metrics: gopdf-custom-metrics
#spec:
#  selector:
#    matchLabels:
#      k8s-app: metrics
#  targetLabels:
#    - k8s-app
#  endpoints:
#  - port: web
#    interval: 10s
#  namespaceSelector:
#    matchNames:
#      - $KUBE_NAMESPACE
#EOF

#
# + kubectl -n mini-stage1 apply -f -
# Error from server (Forbidden): error when retrieving current configuration of:
# &{0xc42009db00 0xc42035ce00 mini-stage1 metrics STDIN 0xc421e49bd8  false}
# from server for: "STDIN": Forbidden (user=jonathan.hess@sap.com, verb=get, resource=servicemonitors, subresource=) (get servicemonitors.monitoring.coreos.com metrics)
