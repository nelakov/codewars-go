# CodeWars Go

Solutions to [CodeWars](https://www.codewars.com/) katas in Go.

## Structure

```
kyu8/                           # 8 kyu (beginner)
  └── kata_name/
      ├── kata_name.go          # solution
      └── kata_name_test.go     # tests
kyu7/                           # 7 kyu
...
kyu1/                           # 1 kyu (hardest)
```

Each kata is a separate Go package containing the solution and table-driven tests.

## Run tests

```bash
# all katas
go test ./...

# specific kyu level
go test ./kyu8/...

# specific kata
go test ./kyu8/count_vowels/
```
