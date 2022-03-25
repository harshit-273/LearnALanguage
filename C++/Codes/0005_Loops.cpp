#include <bits/stdc++.h>

using namespace std;

int main() {
	// for loop
	cout << "-_-_-_-_-_-_-_-_-_-_-_- For Loop -_-_-_-_-_-_-_-_-_-_-_-" << endl;
	for(int counter = 1; counter <= 10; counter++) {
		cout << counter << endl;
	}

	// while loop
	cout << "-_-_-_-_-_-_-_-_-_-_-_- While Loop -_-_-_-_-_-_-_-_-_-_-_-" << endl;
	int other_counter = 1;
	while(other_counter <= 10) {
		cout << other_counter << endl;
		other_counter++;
	}

	// do while loop
	cout << "-_-_-_-_-_-_-_-_-_-_-_- Do While Loop -_-_-_-_-_-_-_-_-_-_-_-" << endl;
	int some_other_counter = 1;
	do {
		cout << some_other_counter << endl;
		some_other_counter++;
	}while(some_other_counter <= 10);
}