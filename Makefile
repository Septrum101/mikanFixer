# parameters
REMOTE_HOST=root@10.0.0.249
TARGET_DIR=/opt/mikanfixer

# Targets
build: clean
	docker build -f Dockerfile -t ghcr.io/thank243/mikanfixer:dev .

clean:
	docker image prune -f

update:
	ssh $(REMOTE_HOST) "docker-compose -f $(TARGET_DIR)/docker-compose.yml up -d"
	ssh $(REMOTE_HOST) "docker image prune -f"

deploy:
	docker save ghcr.io/thank243/mikanfixer:dev | ssh $(REMOTE_HOST) "docker load"

all: build deploy update
