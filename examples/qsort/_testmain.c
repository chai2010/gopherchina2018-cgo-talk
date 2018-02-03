// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdio.h>
#include <stdlib.h>

#define DIM(x) (sizeof(x)/sizeof((x)[0]))

static int compare(const void* a, const void* b) {
	return ( *(int*)a - *(int*)b );
}

static void printArray(int arr[], int len) {
	int i;
	for(i = 0; i < len; i++) {
		printf ("%d ", arr[i]);
	}
	printf("\n");
}

int main() {
	int values[] = { 42, 9, 101, 95, 27, 25 };
	printArray(values, DIM(values));

	qsort(values, DIM(values), sizeof(values[0]), compare);
	printArray(values, DIM(values));

	return 0;
}
