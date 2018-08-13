#include <string>
//基于std::string做一个最简单的缓存类MyBuffer。除了构造函数和析构函数之外，只有两个成员函数分别是返回底层的数据指针和缓存的大小。因为是二进制缓存，所以我们可以在里面中放置任意数据
struct MyBufferShiming {

    std::string* s_;

	MyBuffer(int size) {
		this->s_ = new std::string(size, char('\0'));
	}
	~MyBuffer() {
		delete this->s_;
	}

	int Size() const {
		return this->s_->size();
	}
	char* Data() {

};
int main() {
	auto pBuf = new MyBufferShiming(1024);

	auto data = pBuf->Data();
	auto size = pBuf->Size();

	delete pBuf;
}