
#include <iostream>
using namespace std;

int multiplicationNumber(int number) {
    int result = number * 567;
    result /= 9;
    result += 7492;
    result *= 235;
    result /= 47;
    result -= 498;

    int auxiliary = 0;
    for(int v = 1; v <= 2; v++) {
        auxiliary = result % 10;
        result /= 10;
    }
    if (auxiliary < 0) {
        auxiliary *= -1;
    }
    return auxiliary;
}

int main()
{
    int test;
    cin >> test;
    while (test-- >= 0) { 
        int number;
        if (cin >> number) {
            cout << multiplicationNumber(number) << '\n';
        }
    }
    return 0;
}
