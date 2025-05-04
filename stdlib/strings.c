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
