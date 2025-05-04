#ifndef COMMIT_H
#define COMMIT_H

char* build_commit(char funcs[][MAX_FUNC_NAME], int funcs_count);

int git_commit(const char* message);

#endif
