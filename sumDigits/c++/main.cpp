#include <iostream>
#include <string>
#include <sstream>

using namespace std;

int sumDigits(int number);
int verificationLength(int number);

int main() {
    int number;
    while (true) {
        cin >> number;
        if (cin.fail() || number == 0) {
            break;
        }
        if (number <= 2000000000) {
            cout << verificationLength(number) << endl;
        }
    }
    return 0;
}

int verificationLength(int number) {
    int sum = sumDigits(number);
    string auxiliary = to_string(sum);
    if (auxiliary.length() > 1) {
        int n = stoi(auxiliary);
        return verificationLength(n);
    }
    return sum;
}

int sumDigits(int number) {
    int sum = 0;
    while (number != 0) {
        sum = (number % 10) + sum;
        number = number / 10;
    }
    return sum;
}
