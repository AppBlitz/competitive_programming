#include <iostream>
#include <string>
#include <algorithm>
#include <cstdlib>

using namespace std;

pair<int, long long> computePalindrome(long long num);
long long reverseNumber(long long num);
bool isPalindrome(long long num);

int main() {
    int n;
    cin >> n;
    cin.ignore(); 

    for (int i = 0; i < n; i++) {
        string line;
        getline(cin, line);
        
        try {
            long long number = stoll(line);
            if (number < 0) continue;
            
            auto result = computePalindrome(number);
            cout << result.first << " " << result.second << endl;
        } catch (...) {
            continue;
        }
    }
    return 0;
}

pair<int, long long> computePalindrome(long long num) {
    int iterations = 0;
    while (iterations < 1000) {
        if (isPalindrome(num)) {
            return make_pair(iterations, num);
        }
        num += reverseNumber(num);
        iterations++;
    }
    return make_pair(iterations, num); 
}

long long reverseNumber(long long num) {
    long long reversed = 0;
    while (num > 0) {
        reversed = reversed * 10 + num % 10;
        num /= 10;
    }
    return reversed;
}

bool isPalindrome(long long num) {
    string s = to_string(num);
    for (int i = 0; i < s.length() / 2; i++) {
        if (s[i] != s[s.length() - 1 - i]) {
            return false;
        }
    }
    return true;
}
