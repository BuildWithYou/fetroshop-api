docker compose -f docker-compose.yaml --project-name fetroshop down --volumes
docker compose -f docker-compose.yaml --project-name fetroshop build --no-cache
docker compose -f docker-compose.yaml --project-name fetroshop up -d