//Write a function "greet" that returns "hello world" of type std::string.
#include <string>
#include <iostream>

std::string greet() {
    return "hello world!";
}

int main() {
    std::cout << greet() << std::endl;
    return 0;
}
