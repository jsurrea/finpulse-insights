# Usage

To apply the Terraform configuration, run the following commands in this directory:

```bash
terraform init
terraform apply
```

You will also need to set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to point to your Google Cloud service account key file.

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/sa-key.json"
```

