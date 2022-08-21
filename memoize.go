package main
  func Memoize(fn func(int) int) func(int) int {
      cache := make(map[int]int)
      return func(n int) int {
          if v, ok := cache[n]; ok {
              return v
          }
          cache[n] = fn(n)
          return cache[n]
      }
  }
