#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#include "diff.h"
#include "var.h"
#include "detect.h"
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

void parser(char** files, int file_count) {
    char commit_msg[COMMIT_LENGTH] = "";

    for (int i = 0; i < file_count; ++i) {
        const char* lang = detect_language(files[i]);
        if (lang == NULL) continue;

        char* diff = get_diff(files[i]);
        if (diff == NULL) continue;

        char* msg = vars_msg(files[i], diff);
        printf("%s\n", msg);
        if (msg && strlen(msg) > 0) {
            if (strlen(commit_msg) + strlen(msg) < COMMIT_LENGTH) {
                strcat(commit_msg, msg);
            } else {
                free(msg);
                break;
            }
        }

        free(msg);
        free(diff);
    }

    printf("commit is: %s\n", commit_msg);
    // commit(commit_msg);
}
