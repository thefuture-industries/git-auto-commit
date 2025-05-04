#ifndef DIFF_H
#define DIFF_H

char* get_diff(const char* file);

void extract_added_functions(const char* diff, const char* lang, char a_funcs[][MAX_FUNC_NAME], int* a_func_count);

void extract_deleted_functions(const char* diff, const char* lang, char d_funcs[][MAX_FUNC_NAME], int* d_func_count);

char** get_staged_files(int* count);

#endif
