#include <bits/stdc++.h>

using namespace std;

int main(){
	int arraySize;
	cout << "Please, enter the size of array: ";
	cin >> arraySize;

	int arrayOfInt[arraySize];
	cout << "Please, enter the values of the array: ";
	for(int i = 0; i < arraySize; i++){
		cin >> arrayOfInt[i];
	}

	cout << "You have entered: ";
	for(int i = 0; i < arraySize; i++){
		cout << arrayOfInt[i] << " ";
	}
}