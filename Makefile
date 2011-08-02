include $(GOROOT)/src/Make.inc

TARG=github.com/kuroneko/configureit

GOFILES=\
	configureit.go\
	string.go\
	int.go\
	user.go\


include $(GOROOT)/src/Make.pkg
