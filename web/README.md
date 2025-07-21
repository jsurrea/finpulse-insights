# Usage

To set up the project, run:

```sh
npm install
```

For linting and formatting, run:

```sh
npm run lint
npm run format
```

For development, run:

```sh
npm run dev
```

For production, run:

```sh
npm run build
```

For unit testing, run:

```sh
npm run test:unit
```

For e2e testing, run:

```sh
npx playwright install
npm run test:e2e
```

To push new updates to the GCP container registry, run these commands. Update the variables as needed.

```bash
docker build --platform=linux/amd64 -t stocks-web:latest .
docker tag stocks-web:latest gcr.io/truora-stocks/stocks-web:latest
docker push gcr.io/truora-stocks/stocks-web:latest
```

