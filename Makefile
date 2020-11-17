IMAGE=xjjo/gotpl
all:
	@echo make build
	@echo make push
	@echo make run
build:
	docker build -t $(IMAGE) .
push:
	docker push $(IMAGE)
run:
	printf '{{.first_name}} {{.last_name}} is {{.age}} years old.\n' > template.txt
	printf 'first_name: Max\nlast_name: Mustermann\nage: 30\n' > user.yml
	docker run -v $(CURDIR):/work -i xjjo/gotpl /work/template.txt < user.yml
