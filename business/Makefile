main: main.go
	docker build -t jimmyhealer/business-service:latest .
	docker push jimmyhealer/business-service:latest

run:
	docker run -p 8100:8100 business-service