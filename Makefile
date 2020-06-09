project=clusterfan
docker_image=nousefreak/clusterfan
secret=123

collector:
	CLUSTERFAN_MASTERURL=http://127.0.01:8080 \
		CLUSTERFAN_MASTERSECRET=$(secret) \
			go run -race cmd/clusterfan/main.go

master:
	CLUSTERFAN_ISMASTER=true \
		CLUSTERFAN_MASTERSECRET=$(secret) \
			go run -race cmd/clusterfan/main.go

publish:
	@if [ "${version}" == "" ]; then echo "Missing version"; exit 1; fi
	# git tag ${version}
	docker buildx use $(project)-builder \
		|| (docker buildx create --name $(project)-builder && docker buildx use $(project)-builder)
	
	docker buildx inspect --bootstrap
	docker buildx build -t $(docker_image):$(version) \
	 	--platform linux/amd64,linux/arm64,linux/arm --push .
	docker buildx rm $(project)-builder
	@echo "Published version $(version)"