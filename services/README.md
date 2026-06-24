# Docker Compose Commands

## Build all containers

```bash
DOCKER_BUILDKIT=1 docker compose --env-file .env build 
```

## Start all containers

```bash
docker compose --env-file .env up -d 
```

## Build and start all containers

```bash
DOCKER_BUILDKIT=1 docker compose --env-file .env up --build -d 
```

## Show running containers

```bash
docker compose ps 
```

## View all logs

```bash
docker compose logs -f 
```

## View planner logs

```bash
docker compose logs -f planner 
```

## View Memgraph logs

```bash
docker compose logs -f memgraph 
```

## Restart planner

```bash
docker compose restart planner 
```

## Stop all containers

```bash
docker compose stop 
```

## Stop and remove containers

```bash
docker compose down 
```

## Stop and remove containers and volumes

```bash
docker compose down -v 
```

## Rebuild planner only

```bash
DOCKER_BUILDKIT=1 docker compose build planner 
```

## Rebuild and restart planner

```bash
DOCKER_BUILDKIT=1 docker compose up -d --build planner 
```

## Verify planner API

```bash
curl http://localhost:8000/v1/models 
```

## Test planner

```bash
curl http://localhost:8000/v1/chat/completions \   -H "Content-Type: application/json" \   -d '{     "model":"architect-planner",     "messages":[       {         "role":"user",         "content":"Create a Go REST API project plan."       }     ],     "temperature":0.2   }' 
```

## Open Memgraph Lab

```text
http://localhost:3000 
```