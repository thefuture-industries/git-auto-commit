#include <stdio.h>
#include <stdlib.h>
#include <string.h>

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
