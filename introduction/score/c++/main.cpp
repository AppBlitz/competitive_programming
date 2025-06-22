#include <iostream>
#include <string>

using namespace std;

int main() {
    int testCase;
    string stringCase;
    
    if (cin >> testCase) {
        for (int i = 0; i < testCase; i++) {
            int sumString = 0;
            int auxiliary = 1;
            
            cin >> ws; 
            getline(cin, stringCase);
            
            for (char c : stringCase) {
                if (c == 79) {  
                    sumString += auxiliary;
                    auxiliary += 1;
                } else {
                    auxiliary = 1;
                }
            }
            
            cout << sumString << endl;
        }
    }
    
    return 0;
}
