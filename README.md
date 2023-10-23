# Online Hardware Store Helper

## Description

Imagine an online store that sells various products: laptops, monitors, phones, system units, watches, microphones.

Anyone can go to the store's website and place an order for several products of the store. For example, you can make an order in which there will be 2 laptops and 1 phone.

Also imagine that you have a warehouse on which there are racks. And there are goods on these shelves. For example, there are laptops and monitors on rack "A". There are mobile phones on the shelf "Б". The names of the racks may not be unique, that is, it is quite possible that there will be 2 different racks "A".

Each product can have several racks on which this product lies. One of these shelves is the main one, and the rest are secondary. For example, mobile phones are on racks "Б", "З", "В", and "Б" is the main rack for phones.

## Usage
At the entrance to the program, the numbers of orders are submitted, for which it is necessary to issue goods on the main shelves.
```shell
go run main.go 10,11,14,15
```

## Example output
```
Страница сборки заказов 10,11,14,15

===Стеллаж А
Ноутбук (id=1)
заказ 10, 2 шт

Телевизор (id=2)
заказ 11, 3 шт

Ноутбук (id=1)
заказ 14, 3 шт

===Стеллаж Б
Телефон (id=3)
заказ 10, 1 шт
доп стеллаж: З,В

===Стеллаж Ж
Системный блок (id=4)
заказ 14, 4 шт

Часы
заказ 15, 1 шт (id=5)
доп стеллаж: А

Микрофон (id=6)
заказ 10, 1 шт
```
