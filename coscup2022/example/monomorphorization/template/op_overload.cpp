#include <iostream>

using namespace std;

template <class T>
T myMax(T x, T y)
{
    return (x > y) ? x : y;
}


class Rectangle {
public:
    Rectangle(int w, int l) : w_(w), l_(l) {}

    int GetSize() { return w_ * l_; }

    int GetWid() { return w_; }

    int GetLen() { return l_; }

private:
    int w_;
    int l_;
};


ostream& operator<<(std::ostream& out, Rectangle a) {
    return out << "size: " << a.GetSize();
}

bool operator> (Rectangle& ra, Rectangle& rb) {
    return ra.GetSize() > rb.GetSize();
}

bool operator< (Rectangle& ra, Rectangle& rb) {
    return ra.GetSize() < rb.GetSize();
}

Rectangle operator+ (Rectangle& ra, Rectangle &rb) {
    int w = ra.GetWid() + rb.GetWid();
    int l = ra.GetLen() + rb.GetLen();
    Rectangle rect(w, l);

    return rect;
}


int main()
{
    Rectangle c_a(3, 4);
    Rectangle c_b(6, 7);

    if (c_a > c_b) {
        cout << "c_a > c_b" << endl;
    } else {
        cout << "c_b > c_a" << endl;
    }

    cout << c_a + c_b << endl;

    Rectangle c_d = myMax<Rectangle>(c_a, c_b);
    cout << c_d << endl;

    return 0;
}


