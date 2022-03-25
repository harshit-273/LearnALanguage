#include <bits/stdc++.h>

using namespace std;

int main() {
	int inputNumber;
	cout << "Please, enter a number to check if the number is prime or not: ";
	cin >> inputNumber;

	bool isPrime = 1;
	for(int i = 2; i <= sqrt(inputNumber); i++){
		if(inputNumber % i == 0){
			cout << inputNumber << " is not a prime number" << endl;
			isPrime = 0;
			break;
		}
	}
	if(isPrime){
		cout << inputNumber << " is a prime number" << endl;
	}

	return (0);
}