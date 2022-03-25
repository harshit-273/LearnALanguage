#include <bits/stdc++.h>

using namespace std;

int main() {
	/* 
	// For Pattern 1,2
	int row, col;
	cout << "Please enter no. of rows and columns - ";
	cin >> row >> col;
	*/
	/* 
	// For Pattern 3, 4, 5, 6, 7
	*/
	int n;
	cout << "Please enter no. of *'s or digits for pattern - ";
	cin >> n;

	/* 
	Pattern - 1

	*****
	*****
	*****
	***** 
	*/

	/*
	for(int i = 0; i < row; i++){
		for(int j = 0; j < col; j++){
			cout << "*";
		}
		cout << endl;
	}
	*/

	/* 
	Pattern - 2
	*****
	*   *
	*   *
	***** 
	*/

	/*
	for(int i = 0; i < row; i++){
		for(int j = 0; j < col; j++){
			if(((i != 0) && (i != (row - 1))) && ((j != 0) && (j != (col - 1)))){
				cout << " ";
			}
			else {
				cout << "*";
			}
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 3
	*****
	****
	***
	**
	*
	*/
	/*
	for(int i = 0; i < n; i++){
		for(int j = (n - i); j > 0; j--){
			cout << "*";
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 4
	    *
	   **
	  ***
	 ****
	*****
	*/
	/*
	for(int i = 1; i <= n; i++){
		for(int j = 1; j <= (n - i); j++){
			cout << " ";
		}
		for(int k = 1; k <= i; k++){
			cout << "*";
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 5
	1
	2 2
	3 3 3
	4 4 4 4
	5 5 5 5 5
	*/
	/*
	for(int i = 1; i <= n; i++){
		for(int j = 1; j <= i; j++){
			cout << j << " ";
		}
		cout << endl;
	}
	*/

	/*
	Pattern - 6 - Floyd's Triangle
	1
	2 3
	4 5 6
	7 8 9 10
	11 12 13 14 15
	*/
	/*
	int init_value = 1;
	for(int i = 1; i <= n ; i++){
		for(int j = 1; j <= i; j++){
			cout << init_value << " ";
			init_value++;
		}
		cout << endl;
	}
	*/
	/*
	Pattern - 7
	*      *
	**    **
	***  ***
	********
	********
	***  ***
	**    **
	*      *
	*/

	for(int i = 1; i <= n; i++){
		for(int j = 1; j <= (2 * n); j++){
			if((j <= i) || (((2 * n) - (j - 1)) <= i)){
				cout << "*";
			}
			else{
				cout << " ";
			}
		}
		cout << endl;
	}
	for(int i = n; i >= 1; i--){
		for(int j = 1; j <= (2 * n); j++){
			if((j <= i) || (((2 * n) - (j - 1)) <= i)){
				cout << "*";
			}
			else{
				cout << " ";
			}
		}
		cout << endl;
	}
}
