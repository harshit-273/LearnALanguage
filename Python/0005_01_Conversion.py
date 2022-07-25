'''
----------------------------------------------- Conversion -----------------------------------------------

- conversion is to be done when we need to use different type of datatype instead of which is available.
'''

int_str = 5
str_int = "400"
float_str = 6.4
str_float = "45.91"
int_float = 3
float_int = 4.5

print(int_str, "which is int, converted to string and has type -", type(str(int_str)))
print(str_int, "which is string, converted to int and has type-", type(int(str_int)))
print(float_str, "which is float, converted to string and has type -", type(str(float_str)))
print(str_float, "which is string, converted to float and has type -", type(float(str_float)))
print(int_float, "which is int, converted to float and has type -", type(float(int_float)))
print(float_int, "which is float, converted to int and has type -", type(int(float_int)))
