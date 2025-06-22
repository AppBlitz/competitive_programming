#include <iostream>
#include <string>
#include <unordered_map>
#include <cctype> 
#include <algorithm> 

using namespace std;

string toUpper(string s) {
    transform(s.begin(), s.end(), s.begin(), 
              [](unsigned char c){ return toupper(c); });
    return s;
}

int main() {
    unordered_map<string, string> mp = {
        {"HELLO", "ENGLISH"},
        {"HOLA", "SPANISH"},
        {"HALLO", "GERMAN"},
        {"BONJOUR", "FRENCH"},
        {"CIAO", "ITALIAN"},
        {"ZDRAVSTVUJTE", "RUSSIAN"}
    };
    
    string s;
    int i = 1;
    
    while (cin >> s && s != "#" && i <= 2000) {
        if (s.length() <= 14) {
            string upper_s = toUpper(s);
            auto it = mp.find(upper_s);
            if (it != mp.end()) {
                cout << "Case " << i << ": " << it->second << endl;
            } else {
                cout << "Case " << i << ": UNKNOWN" << endl;
            }
        } else {
            break;
        }
        i++;
    }
    
    return 0;
}
