# Sources
SRC_DIR=cmd/
SCRAPPER_SRC=$(SRC_DIR)scrapper/main.go

# Destionations
DIR_DEST=bin/
SCRAPPER_NAME=$(DIR_DEST)scrapper

# Compiler flags
GO_FLAGS=-o

all: build-scrapper

build-scrapper:
	go build $(GO_FLAGS) $(SCRAPPER_NAME) $(SCRAPPER_SRC)

run-scrapper: build-scrapper
	./$(SCRAPPER_NAME)

clean: clean-scrapper
clean-scrapper:
	rm -f $(SCRAPPER_NAME)
fclean: clean
	rm -fr $(DIR_DEST)

re: fclean all
