# cut

утилита является аналогом консольной команды cut (man cut).
Утилита принимает строки через STDIN, разбивает по
разделителю (TAB) на колонки и выводит запрошенные.

## Флаги:

* ```-f``` - ```"fields" ``` - выбрать поля (колонки) для вывода
* ```-d``` - ```"delimiter"``` - использовать другой разделитель(не ```\t```)
* ```-s``` - ```"separated"``` - разделитель на выводе

## Примеры:

Ввод:
```bash
go build # Сборка приложения
cat test.txt | ./cut  -f0,1,2,3 -d","
```
Вывод:

```bash
a,b,v,g,
asd,ewr,32,32,
231,32,131,21123,
12312,3211123,21,31,
```