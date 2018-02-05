// Copyright © 2018 ChaiShushan <chaishushan{AT}gmail.com>.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

#include <stdio.h>

struct Int {
	int Twice() {
		const int* p = (int*)(this);
		return (*p) * 2;
	}
};

int main() {
	int x = 42;
	int v = ((Int*)(&x))->Twice();
	printf("%d\n", v);
	return 0;
}
