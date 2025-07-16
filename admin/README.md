# Usage

These are admin scripts meant to be executed in the context of a CockroachDB Docker container. They are used to initialize the database and load data. They're not meant to be run directly on the host machine. They are already included in the Terraform configuration and will be executed automatically when the CockroachDB instance is created. Change them here to update the schema of the database or the data loading process.
