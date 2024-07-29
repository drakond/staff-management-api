
# Staff Management API

## Описание

Этот проект представляет собой API для управления сотрудниками и департаментами с использованием Go и MongoDB. 

## Требования

- Go 1.16+
- MongoDB

## Установка

1. Клонируйте репозиторий:

    ```bash
    git clone https://github.com/drakond/staff-management-api.git
    cd staff-management-api
    ```

2. Установите зависимости:

    ```bash
    go mod tidy
    ```

3. Запустите MongoDB сервер.

## Конфигурация

Измените строку подключения к MongoDB в файле `main.go`:

```go
mongoStorage, err := NewMongoStorage("mongodb://localhost:27017", "your_database_name")
```

Замените `mongodb://localhost:27017` на адрес вашего MongoDB сервера и `your_database_name` на имя вашей базы данных.

## Запуск

Запустите приложение:

```bash
go run main.go
```

## API Эндпоинты

### Сотрудники

- **Создать сотрудника**

  `POST /employee`

  **Тело запроса**:

  ```json
  {
    "id": 1,
    "name": "John Doe",
    "position": "Software Engineer"
  }
  ```

- **Получить сотрудника по ID**

  `GET /employee/:id`

- **Получить всех сотрудников**

  `GET /employee`

- **Обновить сотрудника**

  `PUT /employee/:id`

  **Тело запроса**:

  ```json
  {
    "name": "John Doe",
    "position": "Senior Software Engineer"
  }
  ```

- **Удалить сотрудника**

  `DELETE /employee/:id`

### Департаменты

- **Создать департамент**

  `POST /department`

  **Тело запроса**:

  ```json
  {
    "id": 1,
    "name": "IT",
    "employees": []
  }
  ```

- **Получить департамент по ID**

  `GET /department/:id`

- **Удалить департамент**

  `DELETE /department/:id`

- **Добавить сотрудника в департамент**

  `PUT /department/:department_id/employee/:employee_id`

## Лицензия

Этот проект лицензирован под [MIT License](LICENSE).
