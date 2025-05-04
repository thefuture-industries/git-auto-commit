#ifndef LANGUAGE_H
#define LANGUAGE_H

const char* detect_language(const char* filename);

void extract_functions(const char* diff, const char* lang, char funcs[][MAX_FUNC_NAME], int* func_count);

#endif
