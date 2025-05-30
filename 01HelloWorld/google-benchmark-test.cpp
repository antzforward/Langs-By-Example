//编译指令要变成下面，其他的设置我都设置好了，比如％ＩＮＣＬＵＤＥ％和％ＬＩＢ％路径。
//cl /EHsc /std:c++17 %1.cpp /link benchmark.lib shlwapi.lib
#include <benchmark/benchmark.h>

// The function to be benchmarked
int Fibonacci(int n) {
  if (n <= 1)
    return n;
  return Fibonacci(n - 1) + Fibonacci(n - 2);
}

// The benchmark function
static void BM_Fibonacci(benchmark::State& state) {
  // Perform setup here
  int n = state.range(0);

  for (auto _ : state) {
    // This code gets timed
    int result = Fibonacci(n);
    benchmark::DoNotOptimize(result);
  }
}

// Register the benchmark
BENCHMARK(BM_Fibonacci)->Arg(10)->Arg(20)->Arg(30);

// Run the benchmark
BENCHMARK_MAIN();