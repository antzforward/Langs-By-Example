package main

import (
    "math/rand"
    "testing"
    "time"
)

func generateData(n int) []int {
    rand.Seed(time.Now().UnixNano())
    data := make([]int, n)
    for i := range data {
        data[i] = rand.Intn(1000)
    }
    return data
}

func BenchmarkSort(b *testing.B) {
    data := generateData(100000)
    
    b.Run("QuickSort", func(b *testing.B) {
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            quickSort(data)
        }
    })
    
    b.Run("BuiltinSort", func(b *testing.B) {
        b.ResetTimer()
        for i := 0; i < b.N; i++ {
            sort.Ints(data)
        }
    })
}

func quickSort(arr []int) {
    // 实现快速排序算法
}