#include <string.h>
#include <stdlib.h>
#include <stdio.h>

#include "define.h"

const char* detect_language(const char* filename) {
    const char* ext = strrchr(filename, '.');
    if (!ext) return "";

    if (strcmp(ext, ".go") == 0) return "go";
    if (strcmp(ext, ".py") == 0) return "python";
    if (strcmp(ext, ".js") == 0) return "javascript";
    if (strcmp(ext, ".ts") == 0) return "typescript";
    if (strcmp(ext, ".cpp") == 0) return "cpp";
    if (strcmp(ext, ".c") == 0) return "c";
    if (strcmp(ext, ".java") == 0) return "java";
    if (strcmp(ext, ".cs") == 0) return "csharp";
    if (strcmp(ext, ".rs") == 0) return "rust";
    if (strcmp(ext, ".scala") == 0) return "scala";

    return "";
}

void extract_functions(const char* diff, const char* lang, char funcs[][MAX_FUNC_NAME], int* func_count) {
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
                        strcpy(funcs[*func_count], fname);
                        (*func_count)++;
                    }
                }
            }

            // --- Golang ---
            else if (strcmp(lang, "golang") == 0) {
                if (strstr(buffer, "+func ")) {
                    if (sscanf(buffer, "+func %127[^ (]", fname) == 1) {
                        strcpy(funcs[*func_count], fname);
                        (*func_count)++;
                    }
                }
            }

            // --- Python / Scala ---
            else if (strcmp(lang, "python") == 0 || strcmp(lang, "scala") == 0) {
                if (strstr(buffer, "+def ")) {
                    if (sscanf(buffer, "+def %127[^ (]", fname) == 1) {
                        strcpy(funcs[*func_count], fname);
                        (*func_count)++;
                    }
                }
            }

            // --- Java / C# ---
            else if (strcmp(lang, "java") == 0 || strcmp(lang, "csharp") == 0) {
                if (strchr(buffer, '(') && strstr(buffer, "+public") || strstr(buffer, "+private") || strstr(buffer, "+protected")) {
                    if (sscanf(buffer, "+%*s %*s %127[^ (]", fname) == 1) {
                        strcpy(funcs[*func_count], fname);
                        (*func_count)++;
                    }
                }
            }

            // --- Rust ---
            else if (strcmp(lang, "rust") == 0) {
                if (strstr(buffer, "+fn ")) {
                    if (sscanf(buffer, "+fn %127[^ (]", fname) == 1) {
                        strcpy(funcs[*func_count], fname);
                        (*func_count)++;
                    }
                }
            }

            // --- JavaScript / TypeScript ---
            else if (strcmp(lang, "javascript") == 0 || strcmp(lang, "typescript") == 0) {
                if (strstr(buffer, "+function ")) {
                    if (sscanf(buffer, "+function %127[^ (]", fname) == 1) {
                        strcpy(funcs[*func_count], fname);
                        (*func_count)++;
                    }
                } else {
                    if (sscanf(buffer, "+%127[^ =:(]", fname) == 1 &&
                        strchr(buffer, '(') && strchr(buffer, ')')) {
                        strcpy(funcs[*func_count], fname);
                        (*func_count)++;
                    }
                }
            }
        }

        line += strlen(buffer);
        while (*line == '\n' || *line == '\r') line++;
    }
}
