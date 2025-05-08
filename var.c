#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "define.h"
#include "var.h"

typedef struct {
    char type[32];
    char name[32];
    char value[MAX_LINE_LENGTH];
} Var;

int var_parse(const char* line, Var* var) {
    return sscanf(line, "%31s %31s = %127[^;];", var->type, var->name, var->value) == 3;
}

char* vars_msg(const char* file, char* diff) {
    char* line = strtok(diff, "\n");
    Var old_var = {"", "", ""}, new_var = {"", "", ""};
    int old_found = 0, new_found = 0;
    char commit_msg[MAX_STRING_LENGTH] = "";

    while (line) {
        if (line[0] == '-') {
            if (var_parse(line + 1, &old_var)) old_found = 1;
        } else if (line[0] == '+') {
            if (var_parse(line + 1, &new_var)) new_found = 1;
        }

        if (old_found && new_found) {
            if (strcmp(old_var.name, new_var.name) == 0) {
                if (strcmp(old_var.type, new_var.type) != 0) {
                    snprintf(commit_msg, sizeof(commit_msg), "changed %s -> %s %s", old_var.type, new_var.type, old_var.name);
                } else if (strcmp(old_var.value, new_var.value) != 0) {
                    snprintf(commit_msg, sizeof(commit_msg), "changed value in %s", old_var.name);
                }
            } else if (strcmp(old_var.type, new_var.type) == 0 && strcmp(old_var.value, new_var.value) == 0) {
                snprintf(commit_msg, sizeof(commit_msg), "changed %s -> %s", old_var.name, new_var.name);
            }

            old_found = new_found = 0;
            memset(&old_var, 0, sizeof(old_var));
            memset(&new_var, 0, sizeof(new_var));
        }

        line = strtok(NULL, "\n");
    }

    return strdup(commit_msg);
}
