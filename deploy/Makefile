BUSINESS1_URL=10.0.0.4
BUSINESS2_URL=10.0.0.5
INVENTORY_URL=10.0.0.6
STORAGE_URL=10.0.0.7

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
	docker-compose -f ../LoadBalance/docker-compose.yml up -d

down:
	docker-compose -f ../business/docker-compose.yml down
	docker-compose -f ../inventory/docker-compose.yml down
	docker-compose -f ../storage/docker-compose.yml down
	docker-compose -f ../LoadBalance/docker-compose.yml down