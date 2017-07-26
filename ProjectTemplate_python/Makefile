.PHONY: test

mkvirtualenv:
	@echo "mkvirtualenv --python=/usr/local/bin/python3 coding_challenge"
	@echo "pip install -r requirements.txt"
	@echo "pip install -r test_requirements.txt"

test:
	PYTHONPATH=. py.test --verbose -s

test-infinity:
	PYTHONPATH=. ptw 

# to ignore use: # noqa
lint:
	flake8 solution test
