## Как запускать этот код
```
make up 
```
Эта команда поднимает в докере postgreSql и Selentium

```
make run
```
Запуск программы - для наглядности результаты выгружаются из бд и принтуются в консоль

```
make run
```
Остановить докер-контейнеры. При остановке они удаляются - настройка для удобства разработки


## Какие есть недоработки на данный момент

- не полная обработка ошибок;
- selenium не всегда выгружает всю страницу, причина бага не найдена;
- ссылки на товар полностью соответствуют ссылкам с сайта, но если скопировать их из бд, то они не работают;
