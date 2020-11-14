PROJECTNAME=$(shell basename $(PWD))
GOFILES = $(wildcard */*.go)
DATE="14 Nov 2020"

# Make is verbose in Linux. Made it silent
MAKEFLAGS += --silent

## usage: make [option]
## Options:
.PHONY : help
help : Makefile
	echo "\nDay 1 of #100DaysOfCode - $(DATE) 😸"
	echo "The project $(PROJECTNAME) is covering the basics mentioned in 📖 https://tour.golang.org/\n"
	sed -n 's/^## //p' $<
	sed -n 's/^###/ 🔖/p' $<

### exec: To execute specific go file - `make exec package=tour-hello-world` 
.PHONY: exec
exec:
	go run $(package)/main.go

### list: To list all go files
.PHONY: list
list: 
	echo $(GOFILES)