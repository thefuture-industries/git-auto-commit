#ifndef DIFF_H
#define DIFF_H

char* get_diff(const char* file);

void extract_functions(const char* diff, const char* lang, char funcs[][MAX_FUNC_NAME], int* func_count);

char** get_staged_files(int* count);

#endif
