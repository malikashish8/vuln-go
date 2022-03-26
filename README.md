## Develop

Ensure docker compose is installed. vuln-go uses `gow` to help build in docker compose.
```bash
docker-compose -f docker-compose-dev.yml up --build
```

Stop and delete volume for DB to recreate DB:
```bash
docker-compose down --remove-orphans --volumes --rmi local
```