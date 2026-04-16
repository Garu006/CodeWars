#include <stdio.h>
#include <iostream>
#include <vector>
#include <string>

using namespace std;

std::string smash(const std::vector<std::string>& words){
    string result = ""; // esta variable almacena el resultado final
    for (const std::string& word : words) { // iteramos sobre cada palabra con el vector de palabras
        result += word + " "; // se concatena cada palabra con un espacio entre ellas
    }
    return result; // retonarmos el resultado final
}

int main(){
    vector <string> prueba = {"hola", "mundo", "esto", "es", "una", "prueba"};
    cout << "Resultado: " << smash(prueba) << endl;
    return 0;
}
