#include <fstream>
#include <iostream>
#include <string>

#ifdef DEBUG_BUILD
#define D(x) x
#else
#define D(x)
#endif

using namespace std;

int main(int argc, char *argv[]) {
  if (argc != 2) {
    cerr << "Error: Missing input file name." << endl;
    exit(1);
  }

  int curPosition = 50;
  int countZero = 0;

  D(cout << "Input file: " << argv[1] << endl;)
  ifstream inputFile(argv[1]);

  if (!inputFile.is_open()) {
    cerr << "Error: Failed to open file: " << argv[1] << endl;
    exit(1);
  }

  string instruction;
  while (getline(inputFile, instruction)) {
    D(cout << instruction << endl;)

    char direction = instruction[0];
    int distance = stoi(instruction.substr(1));
    D(cout << direction << " " << distance << endl;)

    distance = distance % 100;

    if (direction == 'L') {
      distance = -distance;
    }
    D(cout << curPosition << " " << distance << endl;)

    int newPosition = (curPosition + distance + 100) % 100;

    if (newPosition == 0) {
      countZero++;
    }

    curPosition = newPosition;
  }

  inputFile.close();

  cout << "Password: " << countZero << endl;
  return 0;
}
