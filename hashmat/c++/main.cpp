
#include <iostream>
#include <limits>

unsigned long long calculaDiference(unsigned long long soldierOne, unsigned long long soldierTwo) {
    unsigned long long diference;
    if (soldierOne > soldierTwo) {
        diference = soldierOne - soldierTwo;
    } else {
        diference = soldierTwo - soldierOne;
    }
    return diference;
}

int main() {
    unsigned long long soldierOne, soldierTwo;
    const unsigned long long MAX_VALUE = 4294967296ULL; 

    while (std::cin >> soldierOne >> soldierTwo) {
        if (soldierOne <= MAX_VALUE && soldierTwo <= MAX_VALUE) {
            std::cout << calculaDiference(soldierOne, soldierTwo) << std::endl;
        }
    }
    return 0;
}
