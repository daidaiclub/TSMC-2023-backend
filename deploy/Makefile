BUSINESS1_URL=10.140.0.11
BUSINESS2_URL=10.140.0.12
INVENTORY_URL=10.140.0.13
STORAGE_URL=10.140.0.16

tools:
	apt update && apt install -y docker.io docker-compose

business1:
	INVENTORY_URL=${INVENTORY_URL} STORAGE_URL=${STORAGE_URL} OTHER_URL=${BUSINESS2_URL} docker-compose -f ../business/docker-compose.yml up -d

business2:
	INVENTORY_URL=${INVENTORY_URL} STORAGE_URL=${STORAGE_URL} OTHER_URL=${BUSINESS1_URL} docker-compose -f ../business/docker-compose.yml up -d

inventory:
	docker-compose -f ../inventory/docker-compose.yml up -d

storage:
	docker-compose -f ../storage/docker-compose.yml up -d

proxy:
	docker-compose -f ../proxy/docker-compose.yml up -d

down:
	docker-compose -f ../business/docker-compose.yml down
	docker-compose -f ../inventory/docker-compose.yml down
	docker-compose -f ../storage/docker-compose.yml down
	docker-compose -f ../proxy/docker-compose.yml down