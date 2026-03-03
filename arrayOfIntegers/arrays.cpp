/*Given an array of integers your solution should find the smallest integer.

For example:

    Given [34, 15, 88, 2] your solution will return 2
    Given [34, -345, -1, 100] your solution will return -345

You can assume, for the purpose of this kata, that the supplied array will not be empty.*/

#include <stdio.h>
#include <iostream>
#include <vector>

using namespace std;
int findSmallest(vector <int> list){
    for (int number : list) {
        if (number < list[0]) {
            list[0] = number;
        }
    }   
    return list[0];
}
 
int main() {
    vector <int> prueba = {34, -35, -1, 100};
    cout << "Resultado: " << findSmallest(prueba) << endl;
    return 0;
}