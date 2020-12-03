IMAGE=xjjo/gotpl
all:
	@echo make docker-build
	@echo make docker-push
	@echo make docker-run

build:
	go build ./...

test:
	go test -v ./...

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)

docker-run:
	printf '{{.first_name}} {{.last_name}} is {{.age}} years old.\n' > template.txt
	printf 'first_name: Max\nlast_name: Mustermann\nage: 30\n' > user.yml
	docker run -v $(CURDIR):/work -i xjjo/gotpl /work/template.txt < user.yml
