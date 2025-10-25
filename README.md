[![Unix Grep Lite CI](https://github.com/alexey0b/unix_grep_lite/actions/workflows/ci.yaml/badge.svg)](https://github.com/alexey0b/unix_grep_lite/actions/workflows/ci.yaml)
[![Coverage Status](https://coveralls.io/repos/github/alexey0b/unix_grep_lite/badge.svg?branch=main)](https://coveralls.io/github/alexey0b/unix_grep_lite?branch=main)

# 🧘🏼‍♀️ Unix Grep Lite

Легковесная реализация утилиты `grep` на Go с поддержкой основных флагов Unix grep.

---

## ✅ Поддерживаемые флаги

| Флаг                    | Описание                      | Пример                                              |
| ----------------------- | ----------------------------- | --------------------------------------------------- |
| `-n, --line-number`     | Показать номера строк         | `echo -e "hello\nworld" \| ./unix_grep_lite -n "world"` |
| `-v, --invert-match`    | Инвертировать совпадения      | `echo -e "hello\nworld" \| ./unix_grep_lite -v "hello"` |
| `-c, --count`           | Подсчитать совпадения         | `echo -e "test\ntest\nother" \| ./unix_grep_lite -c "test"` |
| `-i, --ignore-case`     | Игнорировать регистр          | `echo -e "Hello\nWORLD" \| ./unix_grep_lite -i "hello"` |
| `-F, --fixed-strings`   | Фиксированные строки          | `echo -e "test.txt\ntest" \| ./unix_grep_lite -F "test."` |
| `-A, --after-context N` | N строк после совпадения      | `echo -e "a\nb\nc" \| ./unix_grep_lite -A 1 "b"` |
| `-B, --before-context N`| N строк до совпадения         | `echo -e "a\nb\nc" \| ./unix_grep_lite -B 1 "b"` |
| `-C, --context N`       | N строк до и после совпадения | `echo -e "a\nb\nc" \| ./unix_grep_lite -C 1 "b"` |

---

## ▶️ Использование

- **Клонируйте репозиторий:**

```bash
git clone https://github.com/alexey0b/unix_grep_lite
```

- **Соберите проект:**

```sh
make build
```

- **Посмотрите все доступные команды**

```bash
make help
```

---

## Примеры использования утилиты на текстовых файлов из директории `/example`

### Базовый поиск

```bash
make grep PATTERN="world" INPUT_FILE="example/text.txt"
# Output:
# world
```

### Поиск с номерами строк

```bash
make grep PATTERN="world" FLAGS="-n" INPUT_FILE="example/text.txt"
# Output:
# 2:world
```

### Подсчет совпадений

```bash
make grep PATTERN="test" FLAGS="-c" INPUT_FILE="example/text.txt"
# Output:
# 2
```

### Контекст вокруг совпадения

```bash
make grep PATTERN="pattern" FLAGS="-C 1" INPUT_FILE="example/text.txt"
# Output:
# test
# pattern
# line1
```

### Инвертированный поиск

```bash
make grep PATTERN="world" FLAGS="-v" INPUT_FILE="example/text.txt"
# Output:
# hello
# test
# pattern
...
```

### Игнорирование регистра

```bash
make grep PATTERN="hello" FLAGS="-i" INPUT_FILE="example/text.txt"
# Output:
# hello
# Hello
```

### Фиксированные строки (без regex)

```bash
make grep PATTERN="test." FLAGS="-F" INPUT_FILE="example/text.txt"
# Output:
# test.txt
```

### Комбинированные флаги

```bash
make grep PATTERN="test" FLAGS="-niv" INPUT_FILE="example/text.txt"
# Output:
# 1:hello
# 2:world
# 4:pattern
...
```

---

## 🛠️ Технические ресурсы

### Требования

- Go 1.18+
- Unix/Linux/macOS

--- 

### Зависимости

- **[spf13/pflag](https://github.com/spf13/pflag)** - POSIX/GNU-style флаги
- **[stretchr/testify](https://github.com/stretchr/testify)** - Тестирование

---

## 📚 Полезные команды

```bash
# Посмотреть все доступные команды
make help

# Запустить тесты
make test

# Запустить линтер (должен быть установлен golangci-lint)
make lint 

```

**Справочная информация:**

- [Официальная документация GNU cut](https://www.gnu.org/software/grep/manual/grep.html)

---