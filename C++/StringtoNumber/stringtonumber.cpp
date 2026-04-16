// string to number
#include <iostream>
#include <string>

using namespace std;

int string_to_number(const std::string& str){
    int i = 0;
    int num = 0;

    for (i = 0; i < str.length(); i++) {
        num = num * 10 + (str[i] - '0');
    }
    return  num;
}

int main() {
    string prueba = "123456789";
    cout << "Resultado: " << string_to_number(prueba) << endl;
    return 0;
}