Необходимо разработать Go-приложение, которое итерирует числа Фибоначчи через канальный итератор. Серия чисел Фибоначчи определяется следующим образом:

- Первое число = 0
- Второе число = 1
- Каждое последующее число = сумма двух предыдущих чисел

Требуется реализовать две функции:

- `fibonacciGenerator(n int) <-chan int`: возвращает канал, который генерирует первые `n` чисел Фибоначчи.
- `main()`: вызывает `fibonacciGenerator(n)` и выводит сгенерированные числа.