# Kubernetes Operator for Windows 95

Yes, you read correctly. Does anybody need that? No, but it's awesome.

This project uses the awesome [Windows 95-Electron app](https://github.com/felixrieseberg/windows95)
of Felix Rieseberg.

## Contents

The operator was created using the [Operator SDK](https://github.com/operator-framework/operator-sdk).

- `./deploy`: yaml-files to deploy the operator.
- `./docker`: Dockerfile and assets for Windows95-container.
- `./pkg`: Source code for the operator.
- `./operator-less`: yaml-file to deploy Windows95 on Kubernetes without operator.

## Deploying the operator

To deploy the operator use `kubectl` to create the operator resources:

```sh
kubectl apply -f ./deploy/role.yaml
kubectl apply -f ./deploy/role_binding.yaml
kubectl apply -f ./deploy/service_account.yaml
kubectl apply -f ./deploy/crds/win95_v1alpha1_win95_crd.yaml
kubectl apply -f ./deploy/operator.yaml
```

To create a WIndows 95 instance, modify `./deploy/crds/win95_v1alpha1_win95_cr.yaml`:

- `spec.username`: This will be used as the subdomain for your Windows 95 instance.
- `spec.password`: Password to access the VNC server.
- `spec.domain`: The Ingress URL of your cluster.

Then upload the CR:

```sh
kubectl apply -f ./deploy/crds/win95_v1alpha1_win95_cr.yaml
```

The operator will create a new instance for you. It will be available after a
while under `http://<username>.<domain>` (values configured in the CR).

## Open Issues

A lot! Foremost, the mouse pointer is not working well. Better use the keyboard.
Maybe, I will come around on spending some time on that. But feel free to send a
PR!

## License

This project is provided for educational purposes only. It is not affiliated with and has
not been approved by Microsoft.
