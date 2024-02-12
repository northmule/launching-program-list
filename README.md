# launching-program-list
Запуск программ указанных в json файле конфигурации

### Сборка для своей ОС
```shell
GOOS=windows go build -o /home/user/launching-app_json
```
```shell
GOOS=linux go build -o /home/user/launching-app_json
```

### Пример фйала конфигурации
**Наименование**: apps.json

**Расположение**: в папке с исполняемым файлом

**Содержимое:**
 - Linux
```json 
[
  {
    "sortIndex": 2,
    "path": "/usr/bin/kate",
    "days": [
      1,
      7
    ]
  },
  {
    "sortIndex": 3,
    "path": "/snap/bin/telegram-desktop",
    "days": [
      1,
      7
    ]
  },
  {
    "sortIndex": 4,
    "path": "/usr/bin/skypeforlinux",
    "days": [
      1,
      2,
      3,
      4,
      5,
      6,
      7
    ]
  }
]
```
 - Windows
```json
[
  {
    "sortIndex": 1,
    "path": "C:\\Program Files\\Notepad++\\notepad++.exe",
    "days": [
      1,
      2,
      3,
      4,
      5,
      6,
      7
    ]
  },
  {
    "sortIndex": 2,
    "path": "C:\\Program Files\\Mozilla Firefox\\firefox.exe",
    "days": [
      1,
      2,
      3,
      4,
      5,
      6,
      7
    ]
  }
]

```
* sortIndex - порядок запуска
* path - путь к запускаемому приложению
* days - дни недели в которые запуск возможен (1-7)
