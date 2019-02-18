
# Show this message
help:
	@sed -n '\
		/^#/{ s/../	/ ; H ; d ;};\
		/^[a-z][-a-z0-9.]*:/{ s/:.*/:/p ; x ; s/^\n//; p ; d ;};\
		s/.*// ;\
		x ;\
		d ;\
		' Makefile | sed 's/^\([a-z].*\):/'$$(printf '\033')'[32m\1'$$(printf '\033')'[m:/'

# Run all unit tests
test:
	go test -v ./...

# Tag a new release, increasing the minor version: x.y.z -> x.(y+1).0
release.minor:
	git tag -a $$(git tag -l --sort=-v:refname 'v*' | perl -E '$$_=<>; s/\.([0-9]+)\..*$$/".".($$1+1).".0"/e; print')

# Tag a new release, increasing the patch version: x.y.z -> x.y.(z+1)
release.patch:
	git tag -a $$(git tag -l --sort=-v:refname 'v*' | perl -E '$$_=<>; s/\.([0-9]+)$$/".".($$1+1)/e; print')
