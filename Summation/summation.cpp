#include <iostream>
/*Write a program that finds the summation of every number from 1 to num (both inclusive). 
The number will always be a positive integer greater than 0. Your function only needs to return the result, 
what is shown between parentheses in the example below is how you reach that result and it's not part of it, 
see the sample tests.*/

using namespace std;

int summation(int num){
    int sum = 0;
    for (int i = 1; i <=num; i++) {
        sum += i;
    }
    return sum;
}

int main() {
    int test = 8;
    cout << "Resultado: " << summation(test) << endl;
    return 0;
}