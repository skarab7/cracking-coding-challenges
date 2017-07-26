.PHONY: test

# TODO: go with fswatch
infinititest:
	 while :; do gradle --daemon --parallel test  | tail -n +7 | grep -B 100 ':test FAILED' ; sleep 1; done

test:
	gradle test --daemon --parallel

compile:
	gradle compile

lint:
	gradle checkstyleMain && \
	gradle checkstyleTest && \
	gradle pmdMain && \
	gradle findbugsMain
