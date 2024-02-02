docker compose -f docker-compose.yaml --project-name fetroshop down
docker volume rm fetroshop_api_volume fetroshop_web_volume
docker compose -f docker-compose.yaml --project-name fetroshop up -d