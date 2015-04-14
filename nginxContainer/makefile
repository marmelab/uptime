docker_build__nginx:
	docker build --tag=marmelab/nginx .

docker_run__nginx:
	docker run \
		--name=marmelab__nginx \
		--publish=8080:80 \
		--rm \
		-v nginx.conf:/etc/nginx/nginx.conf:ro \
		--volume=${CURDIR}:/usr/share/nginx/html \
		marmelab/nginx 
