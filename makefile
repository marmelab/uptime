docker_build:
	docker build --tag=marmelab/go .

docker_run:
	docker run \
		--interactive \
		--rm \
		--tty \
		--volume=${CURDIR}:/srv \
		marmelab/go
		