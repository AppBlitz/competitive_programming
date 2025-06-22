
#include <iostream>
#include <string>
#include <regex>

int countWords(const std::string& line) {
    std::regex word_regex("[A-Za-z]+");
    std::sregex_iterator words_begin(line.begin(), line.end(), word_regex);
    std::sregex_iterator words_end;
    return std::distance(words_begin, words_end);
}

int main() {
    std::string line;
    while (std::getline(std::cin, line)) {
        std::cout << countWords(line) << std::endl;
    }
    return 0;
}
