#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "stdlib/strings.h"
#include "define.h"
#include "file.h"

char* build_commit(char funcs[][MAX_FUNC_NAME], int funcs_count) {
    int add_count, del_count, rn_count, ch_count;

    char** added = AdFileFolder(&add_count);
    char** deleted = DelFileFolder(&del_count);
    char** renamed = RnFileFolder(&rn_count);
    char** changed = ChFileFolder(&ch_count);

    char* commit_message = malloc(MAX_LINE_LENGTH * sizeof(char));
    if (!commit_message) return NULL;

    commit_message[0] = '\0';

    if (funcs_count > 0) {
        if (funcs_count == 1) {
            strcat(commit_message, "added ");
            strcat(commit_message, funcs[0]);
            strcat(commit_message, " functionality");
        } else {
            char* funcs_ptr[MAX_FUNC_COUNT];
            for (int i = 0; i < funcs_count; ++i) {
                funcs_ptr[i] = funcs[i];
            }
            char* funcs_str = join_strings(funcs_ptr, funcs_count - 1, ", ");
            char* last_func = funcs[funcs_count - 1];

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
        strcat(commit_message, " including ");
        strcat(commit_message, added_str);
        free(added_str);
    }

    if (del_count > 0) {
        char* deleted_str = join_strings(deleted, del_count, ", ");
        strcat(commit_message, " deleted ");
        strcat(commit_message, deleted_str);
        free(deleted_str);
    }

    if (rn_count > 0) {
        char* renamed_str = join_strings(renamed, rn_count, ", ");
        strcat(commit_message, " renamed ");
        strcat(commit_message, renamed_str);
        free(renamed_str);
    }

    if (ch_count > 0) {
        char* changed_str = join_strings(changed, ch_count, ", ");
        strcat(commit_message, " changed ");
        strcat(commit_message, changed_str);
        free(changed_str);
    }

    if (commit_message[0] == '\0') {
        free(commit_message);
        return strdup("auto commit (github@git-auto-commit)");
    }

    if (strlen(commit_message) > COMMIT_LENGTH) {
        free(commit_message);

        if (funcs_count > 0) {
            char* funcs_ptr[MAX_FUNC_COUNT];
            for (int i = 0; i < funcs_count; ++i) {
                funcs_ptr[i] = funcs[i];
            }
            char* funcs_str = join_strings(funcs_ptr, funcs_count, ", ");
            char* short_commit = malloc(strlen("added ") + strlen(funcs_str) + 12);
            if (!short_commit) return NULL;
            sprintf(short_commit, "added %s functionality", funcs_str);
            free(funcs_str);
            return short_commit;
        }
        if (add_count > 0) {
            char* added_str = join_strings(added, add_count, ", ");
            char* short_commit = malloc(strlen("including ") + strlen(added_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "including %s", added_str);
            free(added_str);
            return short_commit;
        }
        if (del_count > 0) {
            char* del_str = join_strings(deleted, del_count, ", ");
            char* short_commit = malloc(strlen("deleted ") + strlen(del_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "deleted %s", del_str);
            free(del_str);
            return short_commit;
        }
        if (rn_count > 0) {
            char* renamed_str = join_strings(renamed, rn_count, ", ");
            char* short_commit = malloc(strlen("renamed ") + strlen(renamed_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "renamed %s", renamed_str);
            free(renamed_str);
            return short_commit;
        }
        if (ch_count > 0) {
            char* changed_str = join_strings(changed, ch_count, ", ");
            char* short_commit = malloc(strlen("changed ") + strlen(changed_str) + 2);
            if (!short_commit) return NULL;
            sprintf(short_commit, "changed %s", changed_str);
            free(changed_str);
            return short_commit;
        }
    }

    return commit_message;
}

int git_commit(const char* message) {
    char cmd[1024];
    snprintf(cmd, sizeof(cmd), "git commit -m \"%s\"", message);

    return system(cmd);
}
