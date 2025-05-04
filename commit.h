#ifndef COMMIT_H
#define COMMIT_H

char* build_commit(char** funcs, int funcs_count);

int git_commit(const char* message);

#endif
