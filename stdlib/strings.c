#include <stdlib.h>
#include <string.h>

char* concat_strings(const char* str1, const char* str2) {
    char* result = malloc(strlen(str1) + strlen(str2) + 1);
    strcpy(result, str1);
    strcat(result, str2);

    return result;
}

char* join_strings(char* arr[], int len, const char* separator) {
    char* result = malloc(1);
    result[0] = '\0';

    for (int i = 0; i < len; i++) {
        result = concat_strings(result, arr[i]);
        if (i < len - 1) {
            result = concat_strings(result, separator);
        }
    }

    return result;
}

void remove_all_spaces(char* str) {
    char* dst = str;
    int in_space = 0;

    while (*str) {
        if (*str != ' ') {
            *dst++ = *str;
            in_space = 0;
        } else if (!in_space) {
            *dst++ = ' ';
            in_space = 1;
        }
        str++;
    }
    *dst = '\0';

    if (*dst == ' ' && dst != str) {
        memmove(str, str + 1, strlen(str));
    }
}
