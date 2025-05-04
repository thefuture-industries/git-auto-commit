#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#include "define.h"
#include "parser.h"

char* ct_prepare(const char* str) {
    for (size_t i = 0; i < keyword_count; i++) {
        if (strstr(str, keywords[i]) != NULL) {
            size_t len = snprintf(NULL, 0, "added new %s module in %s", keywords[i], str);
            char* result = malloc(len + 1);

            if (result) {
                snprintf(result, len + 1, "added new %s module in %s", keywords[i], str);
            }

            return result;
        }
    }

    return NULL;
}

char* tb_keywords(char funcs[][MAX_FUNC_NAME], size_t func_count) {
    size_t total_len = 1;
    char* result = malloc(total_len);
    if (!result) return NULL;
    result[0] = '\0';

    for (size_t i = 0; i < func_count; i++) {
        char* commit_msg = ct_prepare(funcs[i]);
        if (commit_msg) {
            size_t new_len = total_len + strlen(commit_msg) + 2;
            char* temp = realloc(result, new_len);
            if (!temp) {
                free(result);
                free(commit_msg);
                return NULL;
            }

            result = temp;
            if (total_len > 1) {
                strcat(result, ", ");
            }

            strcat(result, commit_msg);
            free(commit_msg);
            total_len = new_len - 1;
        }
    }

    return result;
}
