# Project creation

Должен быть установлен node.js

`npm init`

`npm install --save sequelize`

`npm install --save pg pg-hstore`

Добавить в корень проекта .gitignore со стройкой `node_modules`

Добавить файл .env, заполнить его данными как в .env.example

Добавляем `.env` в gitignore

Установить глобально

`npm install -g sequelize-cli`

Или подгружать каждый раз

`npx sequelize [command]`

Инициализировать:

`sequelize init`

Создание БД:

`sequelize db:create`

Создание первой модели:

`sequelize model:generate --name guardians --attributes name:string --underscored`
