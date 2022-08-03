#include <iostream>

using namespace std;

template<class T>
T Max(T a, T b) {
    return a > b ? a : b;
}

template<>
string Max<string>(string a, string b) {
    return a.length() > b.length() ? a : b;
}

int main() {
    cout << Max<int>(3, 2) << endl;
    cout << Max<string>(string("aaac"), string ("bbb"));
}