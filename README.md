mockery --dir=internal/repositories  --name=ProductRepositoryInterface --filename=product_repository_interface.go --output=internal/mocks/repomocks --outpkg=repomocks

go test -v internal/services/product_service_test.go