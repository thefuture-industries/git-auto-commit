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
    if (strcmp(ext, ".h") == 0) return "c";
    if (strcmp(ext, ".java") == 0) return "java";
    if (strcmp(ext, ".cs") == 0) return "csharp";
    if (strcmp(ext, ".rs") == 0) return "rust";
    if (strcmp(ext, ".scala") == 0) return "scala";

    return "";
}
