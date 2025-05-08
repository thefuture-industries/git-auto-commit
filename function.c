#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "define.h"
#include "function.h"

void extract_added_functions(const char* diff, const char* lang, char a_funcs[][MAX_FUNC_NAME], int* a_func_count) {
    const char* line = diff;
    char buffer[1024];

    while (*line) {
        sscanf(line, "%[^\n]\n", buffer);
        if (strncmp(buffer, "+", 1) == 0) {
            char fname[128];

            // --- C / C++ ---
            if (strcmp(lang, "c") == 0 || strcmp(lang, "cpp") == 0) {
                if (strchr(buffer, '(') && strchr(buffer, ')') && strchr(buffer, '{')) {
                    if (sscanf(buffer, "+%*[^ ] %127[^ (]", fname) == 1) {
                        strcpy(a_funcs[*a_func_count], fname);
                        (*a_func_count)++;
                    }
                }
            }

            // --- Golang ---
            else if (strcmp(lang, "golang") == 0) {
                if (strstr(buffer, "+func ")) {
                    if (sscanf(buffer, "+func %127[^ (]", fname) == 1) {
                        strcpy(a_funcs[*a_func_count], fname);
                        (*a_func_count)++;
                    }
                }
            }

            // --- Python / Scala ---
            else if (strcmp(lang, "python") == 0 || strcmp(lang, "scala") == 0) {
                if (strstr(buffer, "+def ")) {
                    if (sscanf(buffer, "+def %127[^ (]", fname) == 1) {
                        strcpy(a_funcs[*a_func_count], fname);
                        (*a_func_count)++;
                    }
                }
            }

            // --- Java / C# ---
            else if (strcmp(lang, "java") == 0 || strcmp(lang, "csharp") == 0) {
                if (strchr(buffer, '(') && strstr(buffer, "+public") || strstr(buffer, "+private") || strstr(buffer, "+protected")) {
                    if (sscanf(buffer, "+%*s %*s %127[^ (]", fname) == 1) {
                        strcpy(a_funcs[*a_func_count], fname);
                        (*a_func_count)++;
                    }
                }
            }

            // --- Rust ---
            else if (strcmp(lang, "rust") == 0) {
                if (strstr(buffer, "+fn ")) {
                    if (sscanf(buffer, "+fn %127[^ (]", fname) == 1) {
                        strcpy(a_funcs[*a_func_count], fname);
                        (*a_func_count)++;
                    }
                }
            }

            // --- JavaScript / TypeScript ---
            else if (strcmp(lang, "javascript") == 0 || strcmp(lang, "typescript") == 0) {
                if (strstr(buffer, "+function ")) {
                    if (sscanf(buffer, "+function %127[^ (]", fname) == 1) {
                        strcpy(a_funcs[*a_func_count], fname);
                        (*a_func_count)++;
                    }
                } else {
                    if (sscanf(buffer, "+%127[^ =:(]", fname) == 1 &&
                        strchr(buffer, '(') && strchr(buffer, ')')) {
                        strcpy(a_funcs[*a_func_count], fname);
                        (*a_func_count)++;
                    }
                }
            }
        }

        line += strlen(buffer);
        while (*line == '\n' || *line == '\r') line++;
    }
}

void extract_deleted_functions(const char* diff, const char* lang, char d_funcs[][MAX_FUNC_NAME], int* d_func_count) {
    const char* line = diff;
    char buffer[1024];

    while (*line) {
        sscanf(line, "%[^\n]\n", buffer);
        if (strncmp(buffer, "-", 1) == 0) {
            char fname[128];

            // --- C / C++ ---
            if (strcmp(lang, "c") == 0 || strcmp(lang, "cpp") == 0) {
                if (strchr(buffer, '(') && strchr(buffer, ')') && strchr(buffer, '{')) {
                    if (sscanf(buffer, "+%*[^ ] %127[^ (]", fname) == 1) {
                        strcpy(d_funcs[*d_func_count], fname);
                        (*d_func_count)++;
                    }
                }
            }

            // --- Golang ---
            else if (strcmp(lang, "golang") == 0) {
                if (strstr(buffer, "+func ")) {
                    if (sscanf(buffer, "+func %127[^ (]", fname) == 1) {
                        strcpy(d_funcs[*d_func_count], fname);
                        (*d_func_count)++;
                    }
                }
            }

            // --- Python / Scala ---
            else if (strcmp(lang, "python") == 0 || strcmp(lang, "scala") == 0) {
                if (strstr(buffer, "+def ")) {
                    if (sscanf(buffer, "+def %127[^ (]", fname) == 1) {
                        strcpy(d_funcs[*d_func_count], fname);
                        (*d_func_count)++;
                    }
                }
            }

            // --- Java / C# ---
            else if (strcmp(lang, "java") == 0 || strcmp(lang, "csharp") == 0) {
                if (strchr(buffer, '(') && strstr(buffer, "+public") || strstr(buffer, "+private") || strstr(buffer, "+protected")) {
                    if (sscanf(buffer, "+%*s %*s %127[^ (]", fname) == 1) {
                        strcpy(d_funcs[*d_func_count], fname);
                        (*d_func_count)++;
                    }
                }
            }

            // --- Rust ---
            else if (strcmp(lang, "rust") == 0) {
                if (strstr(buffer, "+fn ")) {
                    if (sscanf(buffer, "+fn %127[^ (]", fname) == 1) {
                        strcpy(d_funcs[*d_func_count], fname);
                        (*d_func_count)++;
                    }
                }
            }

            // --- JavaScript / TypeScript ---
            else if (strcmp(lang, "javascript") == 0 || strcmp(lang, "typescript") == 0) {
                if (strstr(buffer, "+function ")) {
                    if (sscanf(buffer, "+function %127[^ (]", fname) == 1) {
                        strcpy(d_funcs[*d_func_count], fname);
                        (*d_func_count)++;
                    }
                } else {
                    if (sscanf(buffer, "+%127[^ =:(]", fname) == 1 &&
                        strchr(buffer, '(') && strchr(buffer, ')')) {
                        strcpy(d_funcs[*d_func_count], fname);
                        (*d_func_count)++;
                    }
                }
            }
        }

        line += strlen(buffer);
        while (*line == '\n' || *line == '\r') line++;
    }
}
