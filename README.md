# Vault-Audit-Logging
Hashicorp Vault Audit Logging when deployed in Kubernetes

## Problem Statement:

Running Hashicorp Vault and enable its Audit logs is not straight forward as it only supports enabling the audit logging feature in a [Local File](https://github.com/hashicorp/vault-helm/blob/main/values.yaml#L732) audit device which makes the below two problems :

- Limitation that if that file is fully utilized and Vault could not write data on it, Then that will prevent Vault from servicing requests as per the [Official Documentation](https://developer.hashicorp.com/vault/docs/audit)

- The need to get the data out of that local file, may be to SIEM solution or any other Observability platform like ( Datadog, Newrelic, Loki ...etc)

## Solution:

This code will be running as a sidecar container with Vault, It will listen all the time onto the audit file as set in environment variable `AUDIT_FILE_PATH`

It will Log the data away from the file as STDOUT and then it will truncates the file content so the file is never getting utilized

## How it works ?

Build the Go code and push it to your container registry and then In the [Official Vault Helm Chart](https://github.com/hashicorp/vault-helm) we just need to update the `extraContainer` section in its values.yaml file with the below content:

```yaml
extraContainers:
    - name: vault-audit
      image: alyragab/vault-audit:v1 # Change it to your container repository
      imagePullPolicy: Always
      env:
        - name: "AUDIT_FILE_PATH"
          value: "/path/to/audit/file"
```
