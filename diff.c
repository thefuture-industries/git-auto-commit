#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "define.h"
#include "git-root.h"

char* get_diff(const char* file) {
    char* git_root = get_git_root();

    char cmd[512];
    snprintf(cmd, sizeof(cmd), "git diff --cached -- %s/%s", git_root, file);

    FILE* fp = _popen(cmd, "r");
    if (!fp) return NULL;

    char* buffer = malloc(100000);
    buffer[0] = '\0';

    while (fgets(cmd, sizeof(cmd), fp)) {
        strcat(buffer, cmd);
    }

    _pclose(fp);
    return buffer;
}

char** get_staged_files(int* count) {
    FILE* fp = _popen("git diff --cached --name-only", "r");
    if (!fp) return NULL;

    char** files = malloc(128 * sizeof(char*));
    char line[512];
    int i = 0;

    while (fgets(line, sizeof(line), fp)) {
        line[strcspn(line, "\r\n")] = 0;
        files[i] = strdup(line);
        i++;
    }

    _pclose(fp);
    *count = i;
    return files;
}

