
ln_in_gopath:
	ln -s $$(pwd) $$GOPATH/src/github.com/skarab7/cracking-coding-challenges

test:
	go test  -timeout 2s $$(glide novendor);

test-infinity:
	fswatch -0 -event solution | xargs -0 -I {} -n 1 make test
	
