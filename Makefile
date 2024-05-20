build:
	docker build -t dlopezlo/scds -f Dockerfile .
run:
	docker run -d --rm -h local-scds -p 3000:3000 dlopezlo/scds 
plantuml:
	docker run -d --rm -p 18080:8080 plantuml/plantuml-server:jetty
test:
	cd src && go test ./... -v -cover