#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "define.h"

char* exec_command(const char* cmd) {
    FILE *fp;
    char *output = malloc(MAX_LINE_LENGTH);
    if (output == NULL) {
        perror("Failed to allocate memory");
        return NULL;
    }

    fp = popen(cmd, "r");
    if (fp == NULL) {
        perror("Failed to run command");
        free(output);
        return NULL;
    }

    size_t len = 0;
    while (fgets(output + len, MAX_LINE_LENGTH - len, fp)) {
        len += strlen(output + len);
        if (len >= MAX_LINE_LENGTH - 1) {
            break;
        }
    }
    fclose(fp);

    return output;
}

char** AdFileFolder(int* count) {
    char* cmd_output = exec_command("git diff --cached --name-status");
    if (cmd_output == NULL) {
        *count = 0;
        return NULL;
    }

    char** add = malloc(MAX_LINE_LENGTH * sizeof(char*));
    if (add == NULL) {
        perror("Failed to allocate memory for added files");
        *count = 0;
        free(cmd_output);
        return NULL;
    }

    *count = 0;
    char* line = strtok(cmd_output, "\n");
    while (line != NULL) {
        if (line[0] == 'A') {
            add[*count] = strdup(line + 2);
            (*count)++;
        }

        line = strtok(NULL, "\n");
    }

    free(cmd_output);
    return add;
}

char** DelFileFolder(int* count) {
    char* cmd_output = exec_command("git diff --cached --name-status");
    if (cmd_output == NULL) {
        *count = 0;
        return NULL;
    }

    char** del = malloc(MAX_LINE_LENGTH * sizeof(char*));
    if (del == NULL) {
        perror("Failed to allocate memory for deleted files");
        *count = 0;
        free(cmd_output);
        return NULL;
    }

    *count = 0;
    char* line = strtok(cmd_output, "\n");
    while (line != NULL) {
        if (line[0] == 'D') {
            del[*count] = strdup(line + 2);
            (*count)++;
        }

        line = strtok(NULL, "\n");
    }

    free(cmd_output);
    return del;
}

char** RnFileFolder(int* count) {
    char* cmd_output = exec_command("git diff --cached --name-status");
    if (cmd_output == NULL) {
        *count = 0;
        return NULL;
    }

    char** rn = malloc(MAX_LINE_LENGTH * sizeof(char*));
    if (rn == NULL) {
        perror("Failed to allocate memory for renamed files");
        *count = 0;
        free(cmd_output);
        return NULL;
    }

    *count = 0;
    char* line = strtok(cmd_output, "\n");
    while (line != NULL) {
        if (line[0] == 'R') {
            char oldFile[256], newFile[256];
            sscanf(line, "R %s %s", oldFile, newFile);

            char* rename_msg = malloc(strlen(oldFile) + strlen(newFile) + 5);
            sprintf(rename_msg, "%s -> %s", oldFile, newFile);

            rn[*count] = rename_msg;
            (*count)++;
        }

        line = strtok(NULL, "\n");
    }

    free(cmd_output);
    return rn;
}

char** ChFileFolder(int* count) {
    char* cmd_output = exec_command("git status --porcelain");
    if (cmd_output == NULL) {
        *count = 0;
        return NULL;
    }

    char** ch = malloc(MAX_LINE_LENGTH * sizeof(char*));
    if (ch == NULL) {
        perror("Failed to allocate memory for changed files");
        *count = 0;
        free(cmd_output);
        return NULL;
    }

    *count = 0;
    char* line = strtok(cmd_output, "\n");
    while (line != NULL) {
        if (strlen(line) >= 4) {
            ch[*count] = strdup(line + 3);
            (*count)++;
        }

        line = strtok(NULL, "\n");
    }

    free(cmd_output);
    return ch;
}

void free_file_array(char** arr, int count) {
    for (int i = 0; i < count; i++) {
        free(arr[i]);
    }
    free(arr);
}
