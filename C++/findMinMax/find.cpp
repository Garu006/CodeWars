#include <stdio.h>
#include <iostream>
#include <vector>

using namespace std;

int min(std::vector<int> list) {
    int min = list[0]; // se asume que el primer elemento es el mínimo
    for (int i = 1; i < list.size(); i++ ) { // Se itera a traves del vector comenzando desde el segundo elemento
        if (list[i] < min) {
            min = list[i]; // si se encuentra un elemento menor, se actualiza el valor de min
        }
    }
    return min;
}

int max(std::vector<int> list){
    int max = list[0]; // se asume que el primer valor es el maximo
    for (int i = 1; i < list.size(); i++){
        if (list[i] > max) {
            max = list[i]; // si se encientra un elemento mayor, se actualiza el valor de max
        }
    }
    return max;
}

int main(){
    vector <int> prueba = {3, 4, 7, 1, 6, 2, 8, 5};
    cout << "El numero minimo es: " << min(prueba) << endl;
    cout << "El numero maximo es: " << max(prueba) << endl;
    return 0;
}