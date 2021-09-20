# goCounter

На стандандартный ввод подается список урлов. 

Например:

https://golang.org/
https://golang.org/doc/
https://golang.org/pkg/compress/
https://golang.org/pkg/compress/gzip/
https://golang.org/pkg/crypto/md5/
https://golang.org/pkg/debug/pe/
https://golang.org/pkg/log/syslog/
https://golang.org/pkg/sort/
https://golang.org/pkg/strconv/
https://golang.org/pkg/strings/
https://golang.org/pkg/sync/
https://golang.org/pkg/strings/
https://golang.org/pkg/time/
https://golang.org/pkg/unicode/
https://golang.org/pkg/unsafe/
https://godoc.org/golang.org/x/benchmarks
https://godoc.org/golang.org/x/net
https://godoc.org/golang.org/x/mobile

Необходимо написать утилиту, которая посчитает количество строк "Go",
найденных на каждом из ресурсов. И выведет результат в стандартный поток вывода.

Пример результата работы программы:

Count for https://golang.org/: 20
Count for https://golang.org/pkg/compress/: 9
Count for https://golang.org/doc/: 79
Count for https://golang.org/pkg/crypto/md5/: 10
Count for https://golang.org/pkg/compress/gzip/: 17
Count for https://golang.org/pkg/debug/pe/: 25
Count for https://golang.org/pkg/log/syslog/: 9
Count for https://golang.org/pkg/sync/: 22
Count for https://golang.org/pkg/sort/: 51
Count for https://golang.org/pkg/strconv/: 43
Count for https://golang.org/pkg/unsafe/: 15
Count for https://golang.org/pkg/strings/: 55
Count for https://golang.org/pkg/unicode/: 37
Count for https://golang.org/pkg/time/: 51
Count for https://godoc.org/golang.org/x/net: 27
Count for https://godoc.org/golang.org/x/mobile: 57
Count for https://godoc.org/golang.org/x/benchmarks: 24
Total: 551
