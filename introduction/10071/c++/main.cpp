#include <iostream>
using namespace std;

int main() {
    int v, t;
    while (cin >> v >> t) {
        if (v >= -100 && v <= 100 && t >= 0 && t <= 200) {
            cout << 2 * v * t << endl;
        }
    }
    return 0;
}
