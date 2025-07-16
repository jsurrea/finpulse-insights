# Usage

To run locally, you can use the following command. You will need an `.env` file with the necessary environment variables set up.

```bash
docker build -t stocks-api .
docker run -p 8080:8080 --env-file ../.env stocks-api
```

If you want to use a local CockroachDB instance, you can use the docker compose file for the project instead:

```bash
docker compose -f ../docker-compose.yml up -d
```

To push new updates to the GCP container registry, run these commands. Update the variables as needed.

```bash
docker build --platform=linux/amd64 -t stocks-api:latest .
docker tag stocks-api:latest gcr.io/truora-stocks/stocks-api:latest
docker push gcr.io/truora-stocks/stocks-api:latest
```


