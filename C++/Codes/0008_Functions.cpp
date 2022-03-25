#include <bits/stdc++.h>

using namespace std;

// defination of function
int sum(int number1, int number2){
	return (number1 + number2);
}

int main() {
	int num1, num2;
	cout << "Please, enter the numbers to have their addition: ";
	cin >> num1 >> num2;

	// calling a function
	cout << "Sum of the numbers " << num1 << " and " << num2 << ": " << sum(num1, num2) << endl;
	return (0);
}