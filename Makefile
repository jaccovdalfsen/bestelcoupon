
GO_VERSION=1.24

all: build

build:
	docker build -t order-api ./order-api
	docker build -t coupon-service ./coupon-service

run:
	docker-compose up --build

clean:
	docker-compose down --volumes
