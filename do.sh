#!/bin/bash
set -x

start() {
    go run main.go
}

generate_k8s_config () {
    mkdir -p ./docker/k8s/configs
    CONFIG_ENV_FILE=${CONFIG_ENV_FILE:-.env}
    export CONFIG_ENV=$(sed 's/^/    /' ./${CONFIG_ENV_FILE})

    cat > ./docker/k8s/configs/ga-backend-secrets.yml <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: ga-backend-secrets
  labels:
    app: ga-backend
    component: microservice
    role: ga-backend
stringData:
   $CONFIG_ENV_FILE: |-
$CONFIG_ENV
EOF
}

$*

