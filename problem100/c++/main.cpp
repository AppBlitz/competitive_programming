#include <iostream>
using namespace std;

int calculateNumber(int number) {
    int count = 1;
    while (number != 1) {
        if (number % 2 == 1) {
            number = number * 3 + 1;
            count++;
        } else {
            number = number / 2;
            count++;
        }
    }
    return count;
}

int runNumber(int i, int j) {
    int result = 0;
    int start = i < j ? i : j;
    int end = i > j ? i : j;
    for (int v = start; v <= end; v++) {
        int auxliary = calculateNumber(v);
        if (auxliary > result) {
            result = auxliary;
        }
    }
    return result;
}

int main() {
    int i, j;
    // Lee pares hasta EOF
    while (cin >> i >> j) {
        cout << i << " " << j << " " << runNumber(i, j) << "\n";
    }
    return 0;
}
