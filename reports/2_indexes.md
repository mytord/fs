Индексы
====

Страница: http://fs.xy8.ru/people

Цель: провести нагрузочное тестирование страницы поиска человека по имени и фамилии, найти варианты для оптимизации.

Кол-во анкет пользователей - 1 000 000.

Нагрузочные тесты проводились с помощью инструмента JMeter.

Измеряем Latency и throughput до индекса
---

Исходные данные:
- фильтр по префиксам имени `my` и фамилии `moo`
- первые 30 записей (limit 30, offset 0)
- connect/response таймаут = 15000ms
- максимальная производительность 1 потока ограничена 1 rps
- нагрузка подавалась в течение 180 секунд

Результат:

|  Threads count  |    1 |    3 |    5 |    8 |   10 |    13 |
|----------------:|-----:|-----:|-----:|-----:|-----:|------:|
|       Errors, % |    0 |    0 |    0 |    0 |    0 |   100 |
| Throughput, RPS | 0.90 | 1.10 | 1.10 | 1.20 | 1.10 |  1.10 |
| Latency 99%, ms | 1737 | 3610 | 5019 | 7678 | 9870 | 15000 |
| Latency Avg, ms | 1107 | 2829 | 4540 | 6847 | 9151 | 12391 |

Текущая пропускная способность достаточно низкая. На 13 потоках - начинаем упираться в connect/response.

Добавляем индекс
---

Создаем композитный индекс на first_name + last_name, с учётом селективности: количество уникальных имён в бд больше, чем фамилий.

Запрос:
```
CREATE INDEX idx_profiles_first_name_last_name ON profiles (first_name, last_name);
```

EXPLAIN:

```
EXPLAIN ANALYZE
SELECT * FROM profiles
WHERE first_name like 'my%' and last_name like 'moo%'
LIMIT 30 OFFSET 0
```
```
-> Limit: 30 row(s)  (cost=3517.32 rows=30) (actual time=0.413..2.332 rows=9 loops=1)
    -> Index range scan on profiles using idx_profiles_first_name_last_name over ('my' <= first_name <= 'my...' AND 'moo' <= last_name <= 'moo...'), with index condition: ((`profiles`.first_name like 'my%') and (`profiles`.last_name like 'moo%'))  (cost=3517.32 rows=4253) (actual time=0.411..2.327 rows=9 loops=1)
```


Измеряем Latency и throughput после добавления индекса
---

Исходные данные:
- фильтр по префиксам имени `my` и фамилии `moo`
- первые 30 записей (limit 30, offset 0)
- connect/response таймаут = 5000ms
- максимальная производительность 1 потока ограничена 1 rps
- нагрузка подавалась в течение 180 секунд по схеме (15+180+15)

Результат:

|   Threads count |  10 |   25 |   50 |   75 |  100 |  125 |   200 |
|----------------:|----:|-----:|-----:|-----:|-----:|-----:|------:|
|       Errors, % |   0 |    0 |    0 | 1.36 | 4.81 | 8.45 | 22.24 |
| Throughput, RPS | 9.4 | 23.4 | 46.9 | 64.6 | 75.6 | 83.5 |  96.7 |
| Latency 99%, ms | 375 |  347 |  344 | 5000 | 5000 | 5000 |  5000 |
| Latency Avg, ms | 109 |  106 |  108 |  251 |  426 |  634 |  1578 |

Оптимизация оказалось достаточно эффективной. Система хорошо справляется с нагрузкой до 50 rps. При большей нагрузке - часть запросов начинает отваливаться по connect/response таймауту, а при нагрузке в 200 RPS - на сервере заканчиваются доступные БД-коннекты.