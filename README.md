# Introduce
This is a project(https://github.com/yogeshlonkar/pod-dependency-init-container) rewritten using go, thanks this useful and great project. Some changes have been made when rewriting this project.  First, in the current version of k8s, pod will not automatically inject the environment variable KUBE_NAMESPACE, so you have to create a clusterrolebinding, this is unnessary. The second is the judgment about the status of the pod.

### Settings

| Environment Variable | Required | Default | Description |
| --- | --- | --- | --- |
| POD_LABELS | Yes | - | This is comma (,) separated string of labels of dependency pods which will be checked for `Running` phase. |
| MAX_RETRY | NO | 5 | Maximum number of times for which init container will try to check if dependency pods are `Running`. |

Example usage:
```yaml
spec:
  containers:
  ...
  serviceAccountName: {{ .Values.serviceAccount }} #optional
  initContainers:
  - name: pod-dependency
    image: ylonkar/pod-dependency-init-container:1.0.2
    env:
    - name: POD_LABELS
      value: app=nodeapp,name=mongo-1
    - name: MAX_RETRY
      value: "10"
    - name: NAMESPACE_NAME
      valueFrom: 
        fieldRef: "metadata.namespace"
```

## RBAC
In case of RBAC this container requires `pods` resource `get`, `list`, `watch` access. Which can be provided by below yaml
```yaml
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: pod-dependency-init-container-role
  namespace: test-ns
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
      - list
      - watch
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: pod-dependency-init-container-sa
  namespace: test-ns
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: pod-dependency-init-container-rolebinding
  namespace: test-ns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: pod-dependency-init-container-role
  namespace: test-ns
subjects:
- kind: ServiceAccount
  name: pod-dependency-init-container-sa
  namespace: test-ns
```