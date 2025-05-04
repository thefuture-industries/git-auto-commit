#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "stdlib/strings.h"
#include "strings.h"
#include "define.h"
#include "file.h"

char* build_commit(char a_funcs[][MAX_FUNC_NAME], int a_funcs_count, char d_funcs[][MAX_FUNC_NAME], int d_funcs_count) {
    int add_count, del_count, rn_count, ch_count;

    char** added = ad_f(&add_count);
    char** deleted = del_f(&del_count);
    char** renamed = rn_f(&rn_count);
    char** changed = ch_f(&ch_count);

    char* commit_message = malloc(MAX_LINE_LENGTH * sizeof(char));
    if (!commit_message) return NULL;

    commit_message[0] = '\0';

    if (a_funcs_count > 0) {
        if (a_funcs_count == 1) {
            // strcat(commit_message, "| added ");
            if (commit_message[0] != '\0') {
                strcat(commit_message, " | added ");
            } else {
                strcat(commit_message, "added ");
            }

            strcat(commit_message, a_funcs[0]);
            strcat(commit_message, " functionality");
        } else {
            char* funcs_ptr[MAX_FUNC_COUNT];
            for (int i = 0; i < a_funcs_count; ++i) {
                funcs_ptr[i] = a_funcs[i];
            }
            char* funcs_str = join_strings(funcs_ptr, a_funcs_count - 1, ", ");
            char* last_func = a_funcs[a_funcs_count - 1];

            strcat(commit_message, "added ");
            strcat(commit_message, funcs_str);
            free(funcs_str);

            strcat(commit_message, " and ");
            strcat(commit_message, last_func);
            strcat(commit_message, " functionality");
        }
    }

    if (add_count > 0) {
        char* added_str = join_strings(added, add_count, ", ");
        if (commit_message[0] != '\0') {
            strcat(commit_message, " | including ");
        } else {
            strcat(commit_message, "including ");
        }

        strcat(commit_message, added_str);

        remove_all_spaces(added_str);
        free(added_str);
    }

    if (del_count > 0) {
        char* deleted_str = join_strings(deleted, del_count, ", ");
        if (commit_message[0] != '\0') {
            strcat(commit_message, " | deleted ");
        } else {
            strcat(commit_message, "deleted ");
        }

        strcat(commit_message, deleted_str);

        remove_all_spaces(deleted_str);
        free(deleted_str);
    }

    if (rn_count > 0) {
        char* renamed_str = join_strings(renamed, rn_count, ", ");
        if (commit_message[0] != '\0') {
            strcat(commit_message, " | renamed ");
        } else {
            strcat(commit_message, "renamed ");
        }

        strcat(commit_message, renamed_str);

        remove_all_spaces(renamed_str);
        free(renamed_str);
    }

    if (ch_count > 0) {
        char* changed_str = join_strings(changed, ch_count, ", ");
        if (commit_message[0] != '\0') {
            strcat(commit_message, " | changed ");
        } else {
            strcat(commit_message, "changed ");
        }

        strcat(commit_message, changed_str);

        remove_all_spaces(changed_str);
        free(changed_str);
    }

    if (commit_message[0] == '\0') {
        free(commit_message);
        return strdup("auto commit (github@git-auto-commit)");
    }

    if (strlen(commit_message) > COMMIT_LENGTH) {
        free(commit_message);

        if (a_funcs_count > 0) {
            char* funcs_ptr[MAX_FUNC_COUNT];
            for (int i = 0; i < a_funcs_count; ++i) {
                funcs_ptr[i] = a_funcs[i];
            }
            char* funcs_str = join_strings(funcs_ptr, a_funcs_count, ", ");
            char* short_commit = malloc(strlen("added ") + strlen(funcs_str) + 12);
            if (!short_commit) return NULL;
            sprintf(short_commit, "added %s functionality", funcs_str);
            free(funcs_str);

            remove_all_spaces(short_commit);
            return short_commit;
        }
        if (add_count > 0) {
            char* added_str = join_strings(added, add_count, ", ");
            char* short_commit = malloc(strlen("including ") + strlen(added_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "including %s", added_str);
            free(added_str);

            remove_all_spaces(short_commit);
            return short_commit;
        }
        if (del_count > 0) {
            char* del_str = join_strings(deleted, del_count, ", ");
            char* short_commit = malloc(strlen("deleted ") + strlen(del_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "deleted %s", del_str);
            free(del_str);

            remove_all_spaces(short_commit);
            return short_commit;
        }
        if (rn_count > 0) {
            char* renamed_str = join_strings(renamed, rn_count, ", ");
            char* short_commit = malloc(strlen("renamed ") + strlen(renamed_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "renamed %s", renamed_str);
            free(renamed_str);

            remove_all_spaces(short_commit);
            return short_commit;
        }
        if (ch_count > 0) {
            char* changed_str = join_strings(changed, ch_count, ", ");
            char* short_commit = malloc(strlen("changed ") + strlen(changed_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "changed %s", changed_str);
            free(changed_str);

            remove_all_spaces(short_commit);
            return short_commit;
        }
    }

    if (d_funcs_count > 0) {
        char* d_ptrs[MAX_FUNC_COUNT];
        for (int i = 0; i < d_funcs_count; ++i) {
            d_ptrs[i] = d_funcs[i];
        }

        char* d_str = join_strings(d_ptrs, d_funcs_count, ", ");
        const char* suffix = " | deleted functionality: ";
        size_t extra_len = strlen(suffix) + strlen(d_str);

        if (strlen(commit_message) + extra_len < COMMIT_LENGTH) {
            strcat(commit_message, suffix);
            strcat(commit_message, d_str);
        }

        free(d_str);
    }

    remove_all_spaces(commit_message);
    return commit_message;
}

int git_commit(const char* message) {
    char cmd[1024];
    snprintf(cmd, sizeof(cmd), "git commit -m \"%s\"", message);

    return system(cmd);
}
