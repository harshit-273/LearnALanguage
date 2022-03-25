#include <bits/stdc++.h>

using namespace std;

int main() {
	// break
	cout << "-_-_-_-_-_-_-_-_-_-_-_- Break -_-_-_-_-_-_-_-_-_-_-_-" << endl;
	for(int counter = 1; counter <= 5; counter++) {
		if(counter > 3) {
			cout << "terminating loop at counter greater than 3" << endl;
			break;
		}
		cout << counter << endl;
	}

	// continue
	cout << "-_-_-_-_-_-_-_-_-_-_-_- Continue -_-_-_-_-_-_-_-_-_-_-_-" << endl;
	for(int other_counter = 1; other_counter <= 5; other_counter++) {
		if(other_counter % 2 == 0) {
			cout << "skipping even part" << endl;
			continue;
		}
		cout << other_counter << endl;
	}
}