## Golang REST API 
### Сервис для хранения заметок пользователей
#### Создать пользователя
POST /users
#### Получить пользователя по id
GET /users/{userId}
#### Удалить пользователя по id
DELETE /users/{userId}
#### Создать заметку пользователя
POST /users/{userId}/notes
#### Получить список заметок пользователя
GET /users/{userId}/notes
#### Создать заметку пользователя по id
GET /users/{userId}/notes/{noteId}
#### Обновить заметку пользователя по id
PATCH /users/{userId}/notes/{noteId}
#### Удалить заметку пользователя по id
DELETE /users/{userId}/notes/{noteId}
