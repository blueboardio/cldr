go ?= GO111MODULE=on go
export go

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

# Print the Go version corresponding to Git's HEAD
go-version: go.mod $(shell $(go) list -f '{{$$Dir := .Dir}}{{range .GoFiles}}{{$$Dir}}/{{.}} {{end}}' ./...)
	@TZ=UTC git log -1 '--date=format-local:%Y%m%d%H%M%S' --abbrev=12 --pretty=tformat:$$(git describe --tags --abbrev=0 '--match=v[0-9]*')'-%cd-%h' $^

# Print the Go command to upgrade a project dependency to Git's HEAD
go-get:
	@echo $(go) get $(shell $(go) list .)@$(shell $(MAKE) -f $(firstword $(MAKEFILE_LIST)) go-version)

# Tag a new release, increasing the minor version: x.y.z -> x.(y+1).0
release.minor:
	git tag -a $$(git tag -l --sort=-v:refname 'v*' | perl -E '$$_=<>; s/\.([0-9]+)\..*$$/".".($$1+1).".0"/e; print')

# Tag a new release, increasing the patch version: x.y.z -> x.y.(z+1)
release.patch:
	git tag -a $$(git tag -l --sort=-v:refname 'v*' | perl -E '$$_=<>; s/\.([0-9]+)$$/".".($$1+1)/e; print')

# Bump last non-pushed tag to HEAD
bump-tag:
	# Check if tag has already been pushed...
	t=$$(git tag -l --sort=-v:refname 'v*' | head -n1); ! git ls-remote --exit-code --tags origin $$t
	t=$$(git tag -l --sort=-v:refname 'v*' | head -n1); git tag -f -a -m "$$(git tag -l '--format=%(contents)' $$t)" $$t

# Edit the message attached to the last tag
edit-tag:
	# Check if tag has already been pushed...
	t=$$(git tag -l --sort=-v:refname 'v*' | head -n1); ! git ls-remote --exit-code --tags origin $$t
	t=$$(git tag -l --sort=-v:refname 'v*' | head -n1); git tag -f -a $$t $$t^{}


# Dump changelog from Git tags
changelog:
	@git tag -l --sort=-v:refname "--format=[%(refname:short)] %(contents)*****************************" 'v*'
