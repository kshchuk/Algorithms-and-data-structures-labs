#include "Matrix.h"

#include <stdexcept>
#include <cassert>


Matrix Matrix::Gauss_JordanElimination(Matrix to_get_inversed)
{
	assert(to_get_inversed.get_columns_count() == 
		to_get_inversed.get_rows_count());

	size_t matrix_size = to_get_inversed.get_rows_count();

	Matrix inversed_matrix(matrix_size, matrix_size);
	inversed_matrix.MakeIndentityMatrix();

	double multiplier;

	// Проходимося по кожному стовпчику, перетворюємо матрицю в діагональну матрицю
	for (size_t i = 0; i < matrix_size; i++)
	{

	    // По кожному рядку
		for (size_t j = 0; j < matrix_size; j++)	
		{
		if (fabs(to_get_inversed[i][i]) < eps)
		{
			// std::cout << "Determinant = 0\n";
			throw std::overflow_error("Determinant = 0");
		}
			if (j != i)
			{
				multiplier = to_get_inversed[j][i] / to_get_inversed[i][i];

				to_get_inversed[j] -= to_get_inversed[i] * multiplier;
				inversed_matrix[j] -= inversed_matrix[i] * multiplier;
			}
		}	
	}
	// Переходимо від діагональної матриці в одиничну
	for (size_t i = 0; i < matrix_size; i++)
	{
		multiplier = 1 / to_get_inversed[i][i];
		
		inversed_matrix[i] *= multiplier;
	}

	return inversed_matrix;
}