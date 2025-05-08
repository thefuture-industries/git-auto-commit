#ifndef FUNCTION_H
#define FUNCTION_H

void extract_added_functions(const char* diff, const char* lang, char a_funcs[][MAX_FUNC_NAME], int* a_func_count);

void extract_deleted_functions(const char* diff, const char* lang, char d_funcs[][MAX_FUNC_NAME], int* d_func_count);

#endif
