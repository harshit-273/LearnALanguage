#include <bits/stdc++.h>

using namespace std;

int main() {
	int inputNumber;
	cout << "Please, enter the number to check if it is armstrong number or not: ";
	cin >> inputNumber;

	int checkingNumber = 0, otherNumber = 0;
	for(int currentNumber = inputNumber; currentNumber != 0; currentNumber = currentNumber / 10){
		otherNumber = currentNumber % 10;
		checkingNumber = (otherNumber * otherNumber * otherNumber) + checkingNumber;
	}
	if(checkingNumber == inputNumber){
		cout << "Number " << inputNumber << " is an amrstrong number" << endl;
	}
	else{
		cout << "Number " << inputNumber << " is not an amrstrong number" << endl;
	}

	return (0);
}