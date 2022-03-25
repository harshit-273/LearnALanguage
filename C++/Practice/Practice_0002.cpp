#include<bits/stdc++.h>

using namespace std;

int main(){

	int n;
	cout << "Please enter no. of *'s or digits for pattern - ";
	cin >> n;

	/*
	Pattern - 1
	1 2 3 4 5
	1 2 3 4
	1 2 3
	1 2
	1
	*/
	/* Pattern - 1
	for(int i = n; i >= 1; i--){
		for(int j = 1; j <= i; j++){
			cout << j << " ";
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 2
	1
	0 1
	1 0 1
	0 1 0 1
	1 0 1 0 1
	*/
	/*
	for(int i = 1; i <= n; i++){
		for(int j = 1; j <= i; j++){
			if(((i + j) % 2) == 0){
				cout << "1 ";
			}
			else{
				cout << "0 ";
			}
		}
		cout << endl;
	}*/

	/*
	Pattern - 3
	    *****
	   *****
	  *****
	 *****
	*****
	*/
	/*
	for(int i = 1; i <= n; i++){
		for(int j = (n - i); j >= 1; j--){
			cout << " ";
		}
		for(int j = 1; j <= n; j++){
			cout << "*";
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 4
	    1
	   1 2
	  1 2 3
	 1 2 3 4
	1 2 3 4 5  
	*/
	/*
	for(int i = 1; i <= n; i++){
		for(int j = 1; j <= (n - i); j++){
			cout << " ";
		}
		for(int j = 1; j <= i; j++){
			cout << j << " ";
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 5
	    1
	   212
	  32123
	 4321234
	543212345
	*/
	/*
	for(int i = 1; i <= n; i++){
		for(int j = 1; j <= (n - i); j++){
			cout << " ";
		}
		for(int j = i; j >= 1; j--){
			cout << j;
		}
		for(int j = 2; j <= i; j++){
			cout << j;
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 6 - Diamond
	   *
	  ***
	 *****
	*******
	*******
	 *****
	  ***
	   *
	*/
	/*
	for(int i = 1; i <= n; i++){
		for(int j = (n - i); j >= 1; j--){
			cout << " ";
		}
		for(int j = 1; j <= ((2 * i) - 1); j++){
			cout << "*";
		}
		cout << endl;
	}
	for(int i = n; i >= 1; i--){
		for(int j = (n - i); j >= 1; j--){
			cout << " ";
		}
		for(int j = 1; j <= ((2 * i) - 1); j++){
			cout << "*";
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 7 - Zig-Zag(M)
	  *   *
	 * * * *
	*   *   *
	*/
	for(int i = 1; i <= 3; i++){
		for(int j = 1; j <= n; j++){
			if((((i + j) % 4) == 0) || (i == 2 && ((j % 4) == 0))){
				cout << "*";
			}
			else{
				cout << " ";
			}
		}
		cout << endl;
	}
}