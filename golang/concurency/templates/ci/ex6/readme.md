1. Создаем канал `ch` типа `chan int`.
2. Запускаем функцию `sendItems(ch)` в отдельной горутине. В этой функции мы отправляем элементы в канал.
3. В функции sendItems выполняется цикл `for` от 1 до 5, и каждый элемент отправляется в канал с помощью оператора `<-`.
4. В функции `main` используем цикл `for range` для итерации по элементам канала `ch`. Когда канал закрывается (в данном случае, после отправки всех элементов), цикл автоматически завершается.
5. Внутри цикла выводим каждый элемент `item` в консоль.

После компиляции и запуска этого кода, вы увидите вывод чисел от 1 до 5 в консоль. Это гарантирует правильность итерации по элементам, отправленным в канал.