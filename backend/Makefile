build:
	@echo "Building image..."
	@( \
		docker login; \
		docker build -t $${LINGUANSKI_IMAGE}:$${CI_COMMIT_SHORT_SHA} . --no-cache; \
		docker push $${LINGUANSKI_IMAGE}:$${CI_COMMIT_SHORT_SHA}; \
		docker tag $${LINGUANSKI_IMAGE}:$${CI_COMMIT_SHORT_SHA} $${LINGUANSKI_IMAGE}:latest; \
		docker push $${LINGUANSKI_IMAGE}:latest; \
	)

deploy:
	@echo "Deploying imagen..."
	@( \
		docker pull $${LINGUANSKI_IMAGE}:$${CI_COMMIT_SHORT_SHA}; \
		docker-compose up -d --no-build; \
	)
