docker compose -f docker-compose.yaml down --volumes
docker compose -f docker-compose.yaml build --no-cache
docker compose -f docker-compose.yaml --project-name fetroshop up -d