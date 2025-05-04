#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* get_diff(const char* file) {
    char cmd[512];
    snprintf(cmd, sizeof(cmd), "git diff --cached %s", file);

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
