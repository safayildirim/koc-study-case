unit-test:
	go test ./... -v

mocks:
	mockgen -destination=mocks/mock_url_service.go -package mocks koc-digital-case/controllers URLService
	mockgen -destination=mocks/mock_url_repository.go -package mocks koc-digital-case/services URLRepository
	mockgen -destination=mocks/mock_url_auth_repository.go -package mocks koc-digital-case/services URLAuthRepository
	mockgen -destination=mocks/mock_auth_service.go -package mocks koc-digital-case/controllers AuthService
	mockgen -destination=mocks/mock_auth_repository.go -package mocks koc-digital-case/services AuthRepository

down:
	docker-compose rm -s -f -v

up:
	docker-compose up --build -d