#include <stdio.h>
#include <stdlib.h>
#include <stdlib.h>
#include <string.h>

#include "define.h"
#include "detect.h"
#include "commit.h"
#include "parser.h"
#include "diff.h"

int main() {
	int file_count = 0;
    char** files = get_staged_files(&file_count);

    if (file_count == 0) {
        printf("No files staged for commit.\n");
        return 0;
    }

    char funcs[MAX_FUNC_COUNT][MAX_FUNC_NAME];
    int func_count = 0;

    for (int i = 0; i < file_count; i++) {
        const char* lang = detect_language(files[i]);
        if (lang == NULL) continue;

        char* diff = get_diff(files[i]);
        if (diff != NULL) {
            extract_functions(diff, lang, funcs, &func_count);
            free(diff);
        }
    }

    char* p_commit_msg = tb_keywords(funcs, file_count);
    if (p_commit_msg && strlen(p_commit_msg) > 0) {
        printf("[git auto-commit] commit is: %s\n", p_commit_msg);

        int result = git_commit(p_commit_msg);
        free(p_commit_msg);
    } else {
        char* commit_msg = build_commit(funcs, func_count);
        printf("[git auto-commit] commit is: %s\n", commit_msg);

        int result = git_commit(commit_msg);
        free(commit_msg);
    }

    for (int i = 0; i < file_count; i++) {
        free(files[i]);
    }
    free(files);

    return 0;
}
