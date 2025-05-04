#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* get_git_root() {
    FILE* fp = _popen("git rev-parse --show-toplevel", "r");
    if (!fp) return NULL;

    char* root = malloc(512);
    if (!root) return NULL;

    if (fgets(root, 512, fp) == NULL) {
        free(root);
        _pclose(fp);
        return NULL;
    }

    root[strcspn(root, "\n")] = '\0';

    _pclose(fp);
    return root;
}
