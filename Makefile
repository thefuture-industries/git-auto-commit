bindir = bin

BUILDIN_MAIN = main.c
BUILDIN_COMMIT = commit.c
BUILDIN_DETECT = detect.c
BUILDIN_DIFF = diff.c
BUILDIN_FILE = file.c
BUILDIN_GIT_ROOT = git-root.c
BUILDIN_PARSER = parser.c
BUILDIN_VAR= var.c
BUILDIN_LOGIC = logic.c
BUILDIN_FUNCTION = function.c
BUILDIN_CLASS = class.c
BUILDIN_STRINGS = stdlib/strings.c

MAIN_OUT = "$(bindir)/auto-commit"

UNAME_S := $(shell uname -s)

CC = gcc

ifeq ($(OS),Windows_NT)
    MKDIR = mkdir $(bindir) || echo "Directory already exists"
    REMOVE_EXT = mv $(MAIN_OUT).exe $(MAIN_OUT)
else
    MKDIR = mkdir -p $(bindir)
    OS_TYPE = $(shell uname -s)
	REMOVE_EXT = true
endif

build:
	$(MKDIR)
	$(CC) $(BUILDIN_MAIN) -o $(MAIN_OUT) $(BUILDIN_COMMIT) \
	$(BUILDIN_DETECT) $(BUILDIN_DIFF) $(BUILDIN_FILE) \
	$(BUILDIN_GET_STAGED) $(BUILDIN_STRINGS) $(BUILDIN_GIT_ROOT) \
	$(BUILDIN_PARSER) $(BUILDIN_VAR) $(BUILDIN_LOGIC) \
	$(BUILDIN_FUNCTION) $(BUILDIN_CLASS)

	$(REMOVE_EXT)

buildt:
	$(MKDIR)
	$(CC) $(BUILDIN_MAIN) -o $(MAIN_OUT) $(BUILDIN_COMMIT) \
	$(BUILDIN_DETECT) $(BUILDIN_DIFF) $(BUILDIN_FILE) \
	$(BUILDIN_GET_STAGED) $(BUILDIN_STRINGS) $(BUILDIN_GIT_ROOT) \
	$(BUILDIN_PARSER)
