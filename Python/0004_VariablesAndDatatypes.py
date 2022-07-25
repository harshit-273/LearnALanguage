'''
----------------------------------------------- Variable -----------------------------------------------

- variables are something which can be used to store some value.
- variables can be used to store many things like number, string, object, etc.,

Syntax:

    var_name = value        # this is how the variable is initialized and this value can be changed as and when needed

----------------------------------------------- Datatypes -----------------------------------------------

- datatypes is something which specifies what type of data is it.
- datatypes available in python are as follow:
    - int           >>          4
    - float         >>          5.3
    - complex       >>          2+5j
    - bool          >>          True/False
    - string        >>          "some string", """ multi line string"""
'''

int_var = 4
float_var = 5.3
complex_var = 2+5j
bool_var = True
string_var = """This is 
multi line
string"""

print(int_var, "is a type of", type(int_var))
print(float_var, "is a type of", type(float_var))
print(complex_var, "is a type of", type(complex_var))
print(bool_var, "is a type of", type(bool_var))
print(string_var, "is a type of", type(string_var))