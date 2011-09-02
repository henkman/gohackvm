PROGRAM = gohvm

ifeq "$(GOOS)" "windows"
	OUTPUT = $(PROGRAM).exe
else
	OUTPUT = $(PROGRAM)
endif

$(OUTPUT):
	gd . -o $(OUTPUT)
clean:
	gd . -clean
	rm -f $(OUTPUT)