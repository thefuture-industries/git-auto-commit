#ifndef PARSER_H
#define PARSER_H

static const char* keywords[] = {"test", "tests", "testing", "http", "https", "image", "resource"};
static const size_t keyword_count = sizeof(keywords) / sizeof(keywords[0]);

typedef struct {
    char type[32];
    char name[64];
} VarStruct;

char* ct_prepare(const char* str);

char* tb_keywords(char funcs[][MAX_FUNC_NAME], size_t func_count);

void parser(char** files, int file_count);

#endif
