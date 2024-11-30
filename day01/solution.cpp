#include <fstream>
#include <iostream>
#include <string>
#include <vector>

std::vector<std::string> readInput(const std::string &filename) {
  std::ifstream file(filename);
  std::vector<std::string> input;
  std::string line;

  while (std::getline(file, line)) {
    input.push_back(line);
  }

  return input;
}

int main() {
  std::vector<std::string> input = readInput("input.txt");

  for (const auto &str : input) {
    std::cout << str << " "; // Print each string followed by a space
  }
  std::cout << std::endl;
  return 0;
}
