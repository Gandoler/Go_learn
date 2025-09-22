#  Go + database/sql — шпаргалка по работе с БД

##  Подключение

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3" // или другой драйвер
)

db, err := sql.Open("sqlite3", "data.db")
if err != nil {
    log.Fatal(err)
}
defer db.Close()
```

---

##  Выполнение запросов

### 1. `Exec` — для `INSERT`, `UPDATE`, `DELETE`, `DDL`

```go
res, err := db.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Alice", 30)
if err != nil { log.Fatal(err) }

id, _ := res.LastInsertId()
rows, _ := res.RowsAffected()
fmt.Println("id:", id, "rows:", rows)
```

---

### 2. `QueryRow` — получить одну строку

```go
var name string
var age int

err := db.QueryRow("SELECT name, age FROM users WHERE id = ?", 1).
    Scan(&name, &age)

if err == sql.ErrNoRows {
    fmt.Println("нет такой записи")
} else if err != nil {
    log.Fatal(err)
}
```

---

### 3. `Query` — получить несколько строк

```go
rows, err := db.Query("SELECT id, name FROM users")
if err != nil { log.Fatal(err) }
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err != nil {
        log.Fatal(err)
    }
    fmt.Println(id, name)
}

if err := rows.Err(); err != nil {
    log.Fatal(err)
}
```

---

### 4. `Prepare` + `Exec/Query`

```go
stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
if err != nil { log.Fatal(err) }
defer stmt.Close()

_, err = stmt.Exec("Bob", 25)
_, err = stmt.Exec("Charlie", 40)
```

>  Хорошо, если один запрос надо выполнять много раз.

---

### 5. Транзакции (`Begin / Commit / Rollback`)

```go
tx, err := db.Begin()
if err != nil { log.Fatal(err) }

_, err = tx.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Dave", 22)
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

_, err = tx.Exec("UPDATE users SET age = ? WHERE name = ?", 23, "Dave")
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

if err = tx.Commit(); err != nil {
    log.Fatal(err)
}
```

---

### 6. `QueryRowContext` / `QueryContext` (с `context.Context`)

> Используются для таймаутов, отмены запросов.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

var count int
err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
```

---

##  Частые приёмы

### Проверка на `sql.ErrNoRows`

```go
err := db.QueryRow("SELECT name FROM users WHERE id = ?", 99).Scan(&name)
if errors.Is(err, sql.ErrNoRows) {
    fmt.Println("Нет записи")
}
```

---

### Получить число (`COUNT`, `SUM`, ...)

```go
var count int
err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
```

---

### Массовая вставка (batch insert)

```go
tx, _ := db.Begin()
stmt, _ := tx.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
defer stmt.Close()

for _, u := range []struct{Name string; Age int}{
    {"Eve", 20}, {"Frank", 35}, {"Grace", 28},
} {
    stmt.Exec(u.Name, u.Age)
}
tx.Commit()
```

---

##  Когда что использовать?

| Метод          | Для чего                                 |
| -------------- | ---------------------------------------- |
| `Exec`         | `INSERT`, `UPDATE`, `DELETE`, `DDL`      |
| `QueryRow`     | Один результат (например, `SELECT ...`)  |
| `Query`        | Несколько строк (например, `SELECT ...`) |
| `Prepare`      | Повторное выполнение одного запроса      |
| `Begin/Commit` | Транзакции                               |
| `*_Context`    | С таймаутами и отменой                   |

---

##  Советы

* Всегда закрывай `rows` → `defer rows.Close()`.
* У `sql.DB` нет `Close` для каждой операции — это пул соединений, закрывается только в конце программы.
* Используй `errors.Is(err, sql.ErrNoRows)` для проверки на «нет результата».
* Для часто используемых запросов → `Prepare`.
* Для сложных операций → транзакции (`Begin`/`Commit`).
