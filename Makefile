##
# Important! Before running any command make sure you have setup GOPATH:
# export GOPATH="$HOME/go"
# PATH="$GOPATH/bin:$PATH"

start:
	# Start the application with postgresql database
	./scripts/start.sh

migrate-up:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

migrate-drop:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" drop

unittest:
	# Runs all unit-tests.
	go test $$(go list ./... | grep -v '/integrationtest') -v

testmocks:
	mockgen \
		-package mocks -destination=internal/mocks/mock_shortener_repository.go \
		-package mocks github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain ShortenerRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_shortener_redis_repository.go \
		-package mocks github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain ShortenerRedisRepository

	mockgen \
		-package mocks -destination=internal/mocks/mock_shortener_service.go \
		-package mocks github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain ShortenerService

