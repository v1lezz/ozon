# Workshop #2



## Задача

Познакомиться с тестирование в Go на практике

## Этап 1
Познакомиться с проектом

## Этап 2
Сгенерировать моки и написать первый позитивный тест

1. Командой minimock сгенерить моки для интерфейсов
2. Написать тест на позитивный сценарий с AddItem
3. Использовать библиотеку testify

## Этап 3
Табличный тест на AddItem 

1. Написать табличный тест для AddItem с позитивным и негативным сценарием
2. Добавить в тест предподготовку моков
3. Использовать библиотеку testify

## Этап 4
Интеграционный тест для хэндлера AddItem

1. Покрываем позитивный сценарий
2. Должен быть поход в productService
3. Должен быть поход в in-memory-storage
4. Инициализация теста через suit
5. Использовать библиотеку testify


## Этап 5
Интеграционный тест для ListItem c testcontainer-ом

1. Покрываем позитивный сценарий
3. Должен быть поход в real-db-storage
4. Использовать postgres-testcontainer
4. Инициализация теста через suit
5. Использовать библиотеку testify

## Полезные ссылку
+ minimock - https://github.com/gojuno/minimock
+ testify - https://github.com/stretchr/testify
+ testcontainers for Go - https://golang.testcontainers.org