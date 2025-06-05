#include <iostream>
#include <cmath>

using namespace std;

int CalculatingRows(int64_t warriors) {
    if (warriors > 0) {
        int auxiliary = sqrt((8 * warriors) + 1);
        return (auxiliary - 1) / 2;
    } else {
        return 0;
    }
}

int main() {
    int64_t n;
    cin >> n;
    
    if (n >= 0 && n < pow(10, 18)) {
        int64_t warriors;
        for (int64_t i = 0; i < n; i++) {
            cin >> warriors;
            cout << CalculatingRows(warriors) << endl;
        }
    }
    
    return 0;
}
