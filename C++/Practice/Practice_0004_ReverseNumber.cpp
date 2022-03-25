#include <bits/stdc++.h>

using namespace std;

int main() {
	int inputNumber, currentNumber, reverseNumber = 0;
	cout << "Please, enter a number to reverse it: ";
	cin >> inputNumber;

	for(currentNumber = inputNumber; currentNumber != 0; currentNumber = currentNumber / 10){
		reverseNumber = (reverseNumber * 10) + (currentNumber % 10);

	}
	cout << "Reverse of number " << inputNumber << " is " << reverseNumber << endl; 

	return (0);
}