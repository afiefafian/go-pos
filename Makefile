build: 
	docker build -t afiefafian95/go-pos:$(tag) .
run:
	docker run --env-file .env -p 3030:3030 afiefafian95/go-pos:$(tag)
push:
	docker push afiefafian95/go-pos:$(tag)
test:
	docker run -e API_URL=172.17.0.1:3030 indraaarmy/posapp-be-test