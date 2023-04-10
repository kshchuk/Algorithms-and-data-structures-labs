#include <iostream>
#include <stdexcept>

#if __has_include(<benchmark/benchmark.h>)
#pragma comment ( lib, "Shlwapi.lib" )
#include <benchmark/benchmark.h>     // https://github.com/google/benchmark 
#endif

#if __has_include(<gtest/gtest.h>)
#include <gtest/gtest.h>        // https://google.github.io/googletest
#endif
#include "Matrix.h"


int main(int argc, char** argv)
{
	// Старт юніт тестів
#ifdef GOOGLETEST_INCLUDE_GTEST_GTEST_H_
	::testing::InitGoogleTest(&argc, argv);
	RUN_ALL_TESTS();
#endif // GOOGLETEST_INCLUDE_GTEST_GTEST_H_
	std::cout << "\n";

//	// Старт бенчмарків
//#ifdef BENCHMARK_BENCHMARK_H_
//	::benchmark::RunSpecifiedBenchmarks();
//#endif // BENCHMARK_BENCHMARK_H_

	std::cout << "Type size:" << std::endl;
	int x, y;
	std::cin >> x >> y;

	Matrix matr(x, y);
	std::cin >> matr;

	try
	{
		matr = Matrix::Gauss_JordanElimination(matr);
		std::cout << matr << std::endl;

	}
	catch (const std::overflow_error& e)
	{
		std::cout << "Overflow exception: " << e.what() << std::endl;
	}


    Row row1(2), row2(2), row3;
    row1[0] = 12; row1[1] = 10;
    row2[0] = 22; row2[1] = 20;

    row3 = row1 + row2;

    // std::cout << row3[0] << ' ' << row3[1] << '\n';

    Matrix m(2, 2); m[0][0] = 5; m[1][1] = 442;

    // std::cout << m;

    Matrix k(m);
    Matrix g = m;

   //while(1)
   //    Matrix d(m);

    std::cout << k << g;

	srand(time(NULL));
	setlocale(LC_ALL, "Ukr");
	int n = 10;
	Matrix A(2, 2);
	Matrix B(2, 2);
	//for (int i = 0; i < n; ++i)
	//	A[i] = new int[n];
	//int** B = new int* [n];
	//for (int i = 0; i < n; ++i)
	//	B[i] = new int[n];
	//int** RES = new int* [n];
	//for (int i = 0; i < n; ++i)
	//	RES[i] = new int[n];

	//SetRandom(A, 2);
	//SetRandom(B, 2);
	A.SetRandom(0, 3);
	B.SetRandom(0, 3);
	std::cout << "\nМатриця А: " << std::endl;
	//Print_Matrix(A, 2);
	std::cout << A;
	std:: cout << std::endl;
	std::cout << "Матриця B: " << std::endl;
	//Print_Matrix(B, 2);
	std::cout << B;
	std::cout << std::endl;
	std::cout << "Результат множення матриць AxB методом Штрассена: " << std::endl;
	Matrix RES = A*B;
	//Print_Matrix(RES, n);
	std::cout << RES;
}

