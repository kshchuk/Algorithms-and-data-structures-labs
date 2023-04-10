#if __has_include(<benchmark/benchmark.h>)
#pragma comment ( lib, "Shlwapi.lib" )
#include <benchmark/benchmark.h>     // https://github.com/google/benchmark 
#endif

#if __has_include(<gtest/gtest.h>)
#include <gtest/gtest.h>        // https://google.github.io/googletest
#endif

#include <random>

#include "Matrix.h"

constexpr int kTestSize = 100;  // розмір тестових матриць/рядків
constexpr int kTestsCount = 100; // кількість викликів кожного тесту
constexpr double kMinMaxElementsRange = 0.1; // елемегти матриці від -5 до 5


#ifdef GOOGLETEST_INCLUDE_GTEST_GTEST_H_


TEST(Row, IndexationOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        Row r(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r[i] = i;
        }

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r[i], i);
        }
    }
}

TEST(Row, CopyOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        Row r(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r[i] = i;
        }

        Row r1;
        r1 = r;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r[i], r1[i]);
        }
    }
}


TEST(Row, CompareOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize), r2(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng); r2[i] = distr(eng);
        }

        EXPECT_FALSE(r1 == r2);

        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = r2[i] = distr(eng);
        }

        EXPECT_TRUE(r1 == r2);
    }
}


TEST(Row, AdditionOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize), r2(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng); r2[i] = distr(eng);
        }

        Row r3 = r1 + r2;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r3[i], r1[i] + r2[i]);
        }
    }
}

TEST(Row, MultiplicationOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng);
        }
        double mult = distr(eng);
        Row r3 = r1 * mult;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r3[i], r1[i] * mult);
        }
    }
}

TEST(Row, DivisionOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng);
        }
        double div = distr(eng);
        Row r3 = r1 / div;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r3[i], r1[i] / div);
        }
    }
}


TEST(Row, SubtractionOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize), r2(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng); r2[i] = distr(eng);
        }

        Row r3 = r1 - r2;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r3[i], r1[i] - r2[i]);
        }
    }
}

TEST(Row, AdditionOperatorWithAssignment)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize), r2(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng); r2[i] = distr(eng);
        }

        Row r3 = r1;

        r1 += r2;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_TRUE(std::abs(r3[i] - (r1[i] - r2[i])) < eps);
        }
    }
}

TEST(Row, SubstractionOperatorWithAssignment)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize), r2(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng); r2[i] = distr(eng);
        }

        Row r3 = r1;
        
        r1 -= r2;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_TRUE(std::abs(r3[i] - (r1[i] + r2[i])) < eps);
        }
    }
}

TEST(Row, MultiplicatioOperatorWithAssignment)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng);
        }

        Row r3 = r1;

        double mult = distr(eng);

        r1 *= mult;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r3[i], r1[i] / mult);
        }
    }
}

TEST(Row, DivisionOperatorWithAssignment)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Row r1(kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
        {
            r1[i] = distr(eng);
        }

        Row r3 = r1;

        double div = distr(eng);

        r1 /= div;

        for (size_t i = 0; i < kTestSize; i++)
        {
            EXPECT_DOUBLE_EQ(r3[i], r1[i] * div);
        }
    }
}

TEST(Matrix, IndexationOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        Matrix m1(kTestSize, kTestSize);
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                m1[i][j] = i * j;

        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                EXPECT_DOUBLE_EQ(m1[i][j], i * j);
    }
}

TEST(Matrix, CopyOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        Matrix m1(kTestSize, kTestSize);
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                m1[i][j] = i * j;

        Matrix m2;
        m2 = m1;
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                EXPECT_DOUBLE_EQ(m2[i][j], i * j);
    }
}

TEST(Matrix, CompareOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Matrix m1(kTestSize, kTestSize), m2(kTestSize, kTestSize);
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                m1[i][j] = m2[i][j] = distr(eng);

        EXPECT_TRUE(m1 == m2);

        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
            {
                m1[i][j] = distr(eng); m2[i][j] = distr(eng);
            }

        EXPECT_FALSE(m1 == m2);
    }
}

TEST(Matrix, MultiplicationOperator)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Matrix m1(kTestSize, kTestSize), m2(kTestSize, kTestSize);
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
            {
                m1[i][j] = distr(eng); m2[i][j] = distr(eng);
            }

        Matrix m4(kTestSize, kTestSize);
        for (size_t i = 0; i < kTestSize; i++)
            for (size_t j = 0; j < kTestSize; j++)
            {
                m4[i][j] = 0;
                for (size_t k = 0; k < kTestSize; k++)
                    m4[i][j] += m1[i][k] * m2[k][j];
            }

        Matrix m3 = m1 * m2;
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                EXPECT_DOUBLE_EQ(m3[i][j], m4[i][j]);
    }
}

TEST(Matrix, setToZero)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-1000, 1000);

        Matrix m1(kTestSize, kTestSize);
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
            {
                m1[i][j] = distr(eng);
            }

        m1.setToZero();

        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                EXPECT_DOUBLE_EQ(m1[i][j], 0);
    }
}

TEST(Matrix, MakeIndentityMatrix)
{
    for (size_t i = 0; i < kTestsCount; i++)
    {
        std::random_device rd;
        std::mt19937 eng(rd());
        std::uniform_real_distribution<double> distr(-0.001, 0.001);

        Matrix m1(kTestSize, kTestSize);
        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
            {
                m1[i][j] = distr(eng);
            }

        m1.MakeIndentityMatrix();

        for (size_t i = 0; i < kTestSize; ++i)
            for (size_t j = 0; j < kTestSize; ++j)
                EXPECT_DOUBLE_EQ(m1[i][j], i == j ? 1 : 0);
    }
}

TEST(MatrixInversion, Gauss_JordanMethod)
{
    Matrix identity_matrix(kTestSize, kTestSize); identity_matrix.MakeIndentityMatrix();
    Matrix matrix(kTestSize, kTestSize), inversed(kTestSize, kTestSize);

    ASSERT_EQ(identity_matrix * matrix, matrix);

    for (size_t j = 0; j < kTestsCount; ++j)
    {
        matrix.SetRandom(kMinMaxElementsRange * -1, kMinMaxElementsRange);
        try
        {
            inversed = Matrix::Gauss_JordanElimination(matrix);
            EXPECT_EQ(inversed * matrix, identity_matrix);
        }
        catch (const std::overflow_error& e) {
            std::cout << e.what() << std::endl;
            EXPECT_STREQ("Determinant = 0", e.what());
        }
        catch (...)
        {
            std::cout << "Неочікувана помилка\n";
            EXPECT_EQ(0, 1);
        }

        // std::cout << inversed << "\n";
    }
}

TEST(InverseMatrix, TestName)
{
    Matrix identity_matrix(kTestSize, kTestSize);
    identity_matrix.MakeIndentityMatrix();

    Matrix matrix(kTestSize, kTestSize);
    Matrix inversed(kTestSize, kTestSize);
    inversed = Matrix::InverseMatrixLU(matrix);
    matrix.SetRandom(kMinMaxElementsRange * -1, kMinMaxElementsRange);

    EXPECT_EQ(inversed * matrix, identity_matrix);
    //EXPECT_EQ(1, 1);
    //EXPECT_TRUE(true);
}

#endif // GOOGLETEST_INCLUDE_GTEST_GTEST_H_


#ifdef BENCHMARK_BENCHMARK_H_

static void Gauss_JordanElimination_BM(benchmark::State& state)
{
    Matrix matrix(state.range(0), state.range(0));

    // Міряємо час роботи алгоритму 
    for (auto _ : state)
    {
        state.PauseTiming();
        matrix.SetRandom(kMinMaxElementsRange * -1, kMinMaxElementsRange);
        state.ResumeTiming();

        try
        {
            benchmark::DoNotOptimize(Matrix::Gauss_JordanElimination(matrix));
        }
        catch (const std::overflow_error& e) {
            std::cout << e.what() << "\n";
        }
    }
    state.SetComplexityN(state.range(0));
}



BENCHMARK(Gauss_JordanElimination_BM)->Unit(benchmark::kMillisecond)->
    RangeMultiplier(2)->Range(2, 2 << 9)->Complexity();

#endif // BENCHMARK_BENCHMARK_H_