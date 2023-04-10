#include "Matrix.h"
#include <cassert>
#include <stdexcept>

Matrix Matrix::InverseMatrixLU(Matrix original)
{
    assert(original.get_columns_count() ==
        original.get_rows_count());

    size_t dim = original.get_rows_count();

    Matrix L(dim, dim), U(dim, dim), L_inverse(dim, dim), U_inverse(dim, dim);
    L.setToZero();
    U.setToZero();
    L_inverse.setToZero();
    U_inverse.setToZero();
    //////////////////////Crout method///////////////////////////
    int i, row, column;
    double ftemp;
    for (row = 0; row < dim; row++)
    {
        L[row][0] = original[row][0];
    }
    for (row = 0; row < dim; row++)
    {
        U[row][row] = 1.0;
    }
    for (column = 1; column < dim; column++)
    {
        U[0][column] = original[0][column] / original[0][0];
    }
    for (column = 1; column < dim; column++)
    {
        for (row = column; row < dim; row++) 
        {
            ftemp = 0.0;
            for (i = 0; i < column; i++)
                ftemp += L[row][i] * U[i][column];
            L[row][column] = original[row][column] - ftemp;
        }
        for (row = column + 1; row < dim; row++) 
        {
            ftemp = 0;
            for (i = 0; i < column; i++)
                ftemp += L[column][i] * U[i][row];
            ftemp = original[column][row] - ftemp;
            U[column][row] = ftemp / L[column][column];
        }
    }
    /////////////////////////////////////////////////////////

    for (row = 0; row < dim; row++) 
    {
        L_inverse[row][row] = 1.0 / L[row][row];
    }
    for (column = 0; column < dim - 1; column++) 
    {
        for (row = column + 1; row < dim; row++) 
        {
            ftemp = 0;
            for (i = column; i < row; i++)
                ftemp += L[row][i] * L_inverse[i][column];
            L_inverse[row][column] = -1.0 * ftemp / L[row][row];
        }
    }

    for (row = 0; row < dim; row++)
    {
        U_inverse[row][row] = 1.0;
    }
    for (row = 0; row < dim - 1; row++) 
    {
        for (column = row + 1; column < dim; column++) 
        {
            ftemp = 0;
            for (i = row; i < column; i++)
            {
                ftemp += U_inverse[row][i] * U[i][column];
            }
            U_inverse[row][column] = -1.0 * ftemp;
            //U_inverse[row][column] = -1.0 * ftemp / U[row][row];
        }
    }
    return U * L;
}