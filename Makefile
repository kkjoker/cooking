gooseup:
	docker-compose exec app bash -c "goose -env local up"
goosedown:
	docker-compose exec app bash -c "goose -env local down"