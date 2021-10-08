<div align="center">
    <img src="resources/images/logobig.png" alt="Local Farm " width="200">
    <h1>The Farmer Journal</h1>
    <a href="https://t.me/mirlocalfarm"><img src="https://img.shields.io/badge/Telegram-blue.svg?logo=telegram&style=flat&label=chat%20on" alt="telegram"></a>
    <img src="https://img.shields.io/badge/semver-1.7.2-green.svg?maxAge=2592000" alt="semver">
    <a href="https://opensource.org/licenses/Apache-2.0" target="_blank"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="License"></a>
</div>

# Внимание

Это ветка разработки проекта Локальная Ферма. Изменения могут происходить ежедневно. Если вам нужна стабильная версия, вы можете открыть [эту ветку](https://github.com/mir-one/localfarm/tree/master).

---

**Local Farm** - программное обеспечение для прогрессивного растениеводства с открытым исходным кодом. Управляйте агробизнесом, контролируйте уровень воды в резервуарах, составляйте списки задач, отслеживайте запасы и прогресс выращивания растений. Локальная ферма Подходит для закрытых и открытых типов ферм.

Скачайте Локальную Ферму для Windows x64 и Linux x64 на [странице релиза](https://github.com/mir-one/localfarm/releases/tag/1.7.1).

![Скриншот](screenshot.PNG)

## Оглавление

* [Начало](#Начало)
    * [Требования](#Требования)
    * [Инструкции по сборке](#Инструкции-по-сборке)
    * [Ядро СУБД](#Ядро-СУБД)
    * [Запуск тестирования](#Запуск-тестирования)
* [REST APIs](#rest-apis)
* [Дорожная карта](#Дорожная-карта)
* [Вклад в проект](#Вклад-в-проект)
    * [Локализация](#localisation)
* [Поддержка](#support-us)
    * [Backers](#backers)
    * [Sponsors](#sponsors)
* [Authors](#authors)
* [Copyright and License](#copyright-and-license)

## Начало

Это программное обеспечение создано на языке программирования [Go](https://golang.org). Это означает, что вы получите исполняемый двоичный файл для запуска на вашем компьютере. Вам **не нужно** дополнительное программное обеспечение, такое как MAMP, XAMPP или WAMP для запуска проекта **Локальная Ферма**, но вам может потребоваться база данных MySQL, если вы решите использовать ее вместо SQLite *(база данных по умолчанию).*

Если ваша ОС не указана на странице релизов, вам придется самостоятельно собрать проект Локальная Ферма для своей ОС. Вы можете следовать нашим инструкциям по сборке **Локальной Фермы**. 

### Требования
- [Go](https://golang.org) >= 1.11
- [NodeJS](https://nodejs.org/en/) 8 or 10

### Инструкции по сборке
1. Clone the repo using `git clone https://github.com/mir-one/localfarm.git`
2. Checkout the current stable version by using `git checkout tags/1.7.2 -b v1.7.2`
3. From the project root, call `go get` to install the Go dependencies.
4. Create a new file `conf.json` using the values from the `conf.json.example` and set it with your own values.
5. Issue `npm install` to install Vue.js dependencies.
6. To build the Vue.js, just run `npm run dev` for development purpose or `npm run prod` for production purpose.
7. Compile the source code with `go build`. It will produces `localfarm.exe` (on Windows) or `localfarm` (on Linux and OSX.)
8. Run the program from Terminal by issuing `./localfarm`, or from Windows Command Prompt by issuing `.\localfarm.exe`.
9. The default username and password are `localfarm / localfarm`.

### Ядро СУБД

Локальная Ферма uses SQLite as the default database engine. You may use MySQL as your database engine by replacing `sqlite` with `mysql` at `localfarm_persistence_engine` field in your `conf.json`.

```
{
  "app_port": "8080",
  "localfarm_persistence_engine": "sqlite",
  "demo_mode": true,
  "upload_path_area": "uploads/areas",
  "upload_path_crop": "uploads/crops",
  "sqlite_path": "db/sqlite/localfarm.db",
  "mysql_host": "127.0.0.1",
  "mysql_port": "3306",
  "mysql_dbname": "localfarm",
  "mysql_user": "root",
  "mysql_password": "root",
  "redirect_uri": [
      "http://localhost:8080",
      "http://127.0.0.1:8080"
  ],
  "client_id": "f0ece679-3f53-463e-b624-73e83049d6ac"
}
```

### Запуск тестирования
- Use `go test ./...` to run all the Go tests.
- Use `npm run cypress:run` to run the end-to-end test

## REST APIs
**Локальная Ферма** have REST APIs to easily integrate with any softwares, even you can build a mobile app client for it. You can import the JSON file inside Postman directory to [Postman app](https://www.getpostman.com).

## Дорожная карта
We want to share our high-level details of our roadmap, so that others can see our priorities in LocalFarm development. You can read our roadmap on [the wiki](https://github.com/mir-one/localfarm/wiki/Roadmap).

## Вклад в проект
We welcome contributions, but request you to follow these [guidelines](contributing.md).

### Локализация

You can help us to localise LocalFarm into your language by following these steps:

1. Copy `languages/template.pot` and paste it to `languages/locale` directory.
2. Rename it with your language locale code e.g: `en_AU.po`, `de_DE.po`, etc.
3. Fill `msgstr` key with your translation. You can edit the `.po` file by using text editor or PO Edit software.
4. Pull request your translation to the `master` branch.

### Build LocalFarm localisation by yourself

**Note:** You will need to install GNU Gettext for your OS. Get it [here](https://www.gnu.org/software/gettext/).

You can build LocalFarm in your language by changing the default language inside `resources/js/app.js`.

```
Vue.use(GetTextPlugin, {
  availableLanguages: { // add your language here
    en_GB: 'British English',
    id_ID: 'Bahasa Indonesia',
    hu_HU: 'Magyar Nyelv'
  },
  defaultLanguage: 'en_GB', // change this to your language
  translations: translations,
  silent: false
})
```

Then follow the instruction to [build LocalFarm](#building-instructions).

## Copyright and License

Copyright to LocalFarm and other contributors under [Apache 2.0](https://github.com/mir-one/localfarm/blob/master/LICENSE) open source license.
