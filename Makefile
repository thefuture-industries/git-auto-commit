bindir = bin

BUILDIN_MAIN = main.c
BUILDIN_COMMIT = commit.c
BUILDIN_DETECT = detect.c
BUILDIN_DIFF = diff.c
BUILDIN_FILE = file.c
BUILDIN_GIT_ROOT = git-root.c
BUILDIN_PARSER = parser.c
BUILDIN_STRINGS = stdlib/strings.c

MAIN_OUT = "$(bindir)/auto-commit"

UNAME_S := $(shell uname -s)

CC = gcc

ifeq ($(OS),Windows_NT)
    MKDIR = mkdir $(bindir) || echo "Directory already exists"
    OS_TYPE = Windows
else
    MKDIR = mkdir -p $(bindir)
    OS_TYPE = $(shell uname -s)
endif

build:
	$(MKDIR)
	$(CC) $(BUILDIN_MAIN) -o $(MAIN_OUT) $(BUILDIN_COMMIT) \
	$(BUILDIN_DETECT) $(BUILDIN_DIFF) $(BUILDIN_FILE) \
	$(BUILDIN_GET_STAGED) $(BUILDIN_STRINGS) $(BUILDIN_GIT_ROOT) \
	$(BUILDIN_PARSER)
