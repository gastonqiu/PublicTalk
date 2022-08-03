#include <iostream>

typedef int UINT4;

using namespace std;

class Hack
{
};

Hack& operator< (Hack &a , Hack &b)
{
    cout << "less than operator" << endl;
    return a;
}

Hack& operator> (Hack &a, Hack &b)
{
    cout <<  "greater than operator" << endl;
    return a;
}

int main(int argc, char ** argv)
{
    Hack vector;
    Hack UINT4;
    Hack foo;

    vector<UINT4> foo;

    return(0);
}
