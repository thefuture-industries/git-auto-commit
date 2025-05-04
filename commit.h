#ifndef COMMIT_H
#define COMMIT_H

char* build_commit(char a_funcs[][MAX_FUNC_NAME], int a_funcs_count, char d_funcs[][MAX_FUNC_NAME], int d_funcs_count);

int git_commit(const char* message);

#endif
