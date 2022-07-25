'''
----------------------------------------------- Operators -----------------------------------------------

- operators are something which carry out some computation.

'''

'''
----------------------------------------------- Arithmetic Operators -----------------------------------------------

- the operators which are used for perfomring some sort of arthmetic operations.

	Operator								Meaning										Example
	----------------------------------------------------------------------------------------------
		+				Addition or unary plus 											a + b
		-				Subtract right operand from	left operand or unary minus			a - b
		*				Multiply two operands											a * b
		/				Divide left operand by right operand(results in float)			a / b
		//				Floor Division(results in int)									a // b
		%				Modulus - remainder of the left operand by right operand		a % b
		**				Exponent - left operand raised to the power of right operand	a ** b

'''

print(4 + 2)
print(5 - 3)
print(7 * 2)
print(9 / 3)
print(9 // 2)
print(5 % 3)
print(2 ** 4, end="\n\n-----------------------------------------------\n")

'''
----------------------------------------------- Comparison Operators -----------------------------------------------

- operators used to compare two operands.

	Operator            					Meaning													Example
	-----------------------------------------------------------------------------------------------------------------
		>			Returns true if left operand is greater than right operand						a > b
		<			Returns true if left operand is less than right operand							a < b
		==			Returns true if left operand is equal to right operand							a == b
		!=			Returns true if left operand is equal to not right operand						a != b
		>=			Returns true if left operand is greater than or equal to right operand			a >= b
		<=			Returns true if left operand is less than or equal to right operand				a <= b

'''

print(2 > 3)
print(2 < 3)
print(2 == 3)
print(2 != 3)
print(2 >= 3)
print(2 <= 3, end="\n\n-----------------------------------------------\n")

'''
----------------------------------------------- Logical Operators -----------------------------------------------

- operators used for logical operations

and		-			true if both the operands are true				-			a and b
or		-			false if both the operands are false			-			a or b
not		-			complements the operand							-			not a

'''

print(True and False)
print(False or False)
print(not False, end="\n\n-----------------------------------------------\n")

'''
----------------------------------------------- Bitwise Operators -----------------------------------------------

- operators used for operation on bits of the number

	Operator				Meaning										Example
	----------------------------------------------------------------------------
		|					Bitwise OR									a | b
		&					Bitwise AND									a & b
		~					Bitwise NOT									~a
		^					Bitwise XOR									a ^ b
		>>					Bitwise right shift							a >> b
		<<					Bitwise left shift							a << b

'''

print(10 | 4)
print(10 & 4)
print(~10)
print(10 ^ 4)
print(10 >> 1)
print(10 << 2)

