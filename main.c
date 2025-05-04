#include <stdio.h>
#include <stdlib.h>
#include <stdlib.h>
#include <string.h>

#include "get-staged.h"
#include "define.h"
#include "detect.h"
#include "commit.h"
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

    char* commit_msg = build_commit(funcs, func_count);
    printf("[git auto-commit] commit is: %s\n", commit_msg);

    // int result = git_commit(commit_msg);
    // free(commit_msg);

    for (int i = 0; i < file_count; i++) {
        free(files[i]);
    }
    free(files);

    return 0;
}
