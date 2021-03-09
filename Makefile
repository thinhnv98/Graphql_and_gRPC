sync-proto:
	cp -r -f ./app/grpc/proto/ ./microservices/book-services
	cp -r -f ./app/grpc/proto/ ./microservices/sale-services
	make -C app/grpc gen
	make -C microservices/book-services gen
	make -C microservices/sale-services gen

run:
	make -j 2 run-microservices run-app

run-app:
	make -C app run

run-microservices:
	make -C microservices/book-services run
