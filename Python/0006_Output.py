'''
----------------------------------------------- Output Formatting -----------------------------------------------

- we can do the output using the in-built given function to print the output.

Syntax:
    print(*objects, sep=' ', end='\n', file=sys.stdout, flush=False)

- This print function takes some objects which can be formatted and printed.
- sep(seperator) is separator used to seperate the objects to be printed, by default it's value is " "(space).
- end is used when the objects are all printed and at last what should be printed, by default it's value is "\n"(new line character).
- file is the file the file where the output has to printed, by default it is the standard output(terminal or something).
'''

print("first object is", 45, "seprator used here is '|'", "and the end of the line would be '>>'", sep="|", end=">>")