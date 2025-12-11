Пример для генерации прото файла:
- ~/GolandProjects/MarketShop/cmd/product-service$ protoc   -I proto   --go_out=proto/generated --go_opt=paths=source_relative   --go-grpc_out=proto/generated --go-grpc_opt=paths=source_relative   proto/product/product.proto
