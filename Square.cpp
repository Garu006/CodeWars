/*Complete the square sum function so that it squares each 
number passed into it and then sums the results together*/

// for example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.

#include <vector>
#include <stdio.h>
#include <iostream>

int square_sum(const std::vector<int>& numbers) {
    int sum = 0;
    for (int number : numbers) {
        sum += number * number;
    }
    return sum;
}

int main() {
    std::vector<int> prueba = {1, 2, 2};
    std::cout << "Resultado: " << square_sum(prueba) << std::endl;
    return 0;
}