# Проект "Управление счетами"
Этот проект представляет собой простой сервер для управления банковскими счетами через HTTP API с использованием Go и фреймворка Echo.

# Установка и запуск
### Установка зависимостей:

```bash
go get -u github.com/labstack/echo/v4
```
### Запуск сервера:

```
bash
go run main.go
```
Сервер будет доступен по адресу http://localhost:1323.

# API Методы
- GET /account
  - Получить информацию о счете по имени.
- POST /account/create
  - Создать новый счет с указанным именем и начальным балансом.
POST /account/delete 
  - Удалить счет по имени.
- POST /account/amount 
  - Изменить баланс счета (пополнение или снятие средств).
- POST /account/name 
  - Изменить имя счета.
- POST /account/transaction 
  - Перевести средства с одного счета на другой.

# Использование командной строки
Вы также можете использовать командную строку для выполнения операций над счетами. Для этого используются параметры командной строки:

```bash
go run cli.go -cmd create -name alice -amount 100
```
Поддерживаемые команды: create, get, delete, changeName, replenishment, withdrawal, transaction.

# Примеры использования
### Создание счета:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "alice", "amount": 50}' http://localhost:1323/account/create
```

### Получение информации о счете:
```bash
curl -X GET http://localhost:1323/account?name=alice
```

### Перевод средств:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "alice", "amount": 20, "to": "bob"}' http://localhost:1323/account/transaction
```
