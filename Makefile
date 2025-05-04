bindir = bindir

BUILDIN_MAIN = main.c
BUILDIN_COMMIT = commit.c
BUILDIN_DETECT = detect.c
BUILDIN_DIFF = diff.c
BUILDIN_FILE = file.c
BUILDIN_GET_STAGED = get-staged.c
BUILDIN_STRINGS = stdlib/strings.c

MAIN_OUT = "$(bindir)/auto-commit"

CC = gcc

build:
	@if not exist $(bindir) mkdir $(bindir)
	$(CC) $(BUILDIN_MAIN) -o $(MAIN_OUT) $(BUILDIN_COMMIT) \
	$(BUILDIN_DETECT) $(BUILDIN_DIFF) $(BUILDIN_FILE) \
	$(BUILDIN_GET_STAGED) $(BUILDIN_STRINGS)
