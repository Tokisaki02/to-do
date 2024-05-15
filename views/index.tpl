<!DOCTYPE html>
<html>
<head>
    <title>To-do App</title>
</head>
<body>
    {{if not .User}}
        <h1>Регистрация</h1>
        <form action="/register" method="post">
            <input type="text" name="username" placeholder="Имя пользователя">
            <input type="password" name="password" placeholder="Пароль">
            <button type="submit">Зарегистрироваться</button>
        </form>
        <p>Уже есть аккаунт? <a href="/login">Войти</a></p>
    {{else}}
        <h1>Мои задачи</h1>
        <ul>
            {{range .Tasks}}
                <li>
                    {{if .Done}}
                        <del>{{.Content}}</del>
                    {{else}}
                        {{.Content}}
                    {{end}}
                    <a href="/toggletaskdone?id={{.Id}}">{{if .Done}}Отметить как не выполненную{{else}}Отметить как выполненную{{end}}</a>
                    <a href="/deletetask?id={{.Id}}">Удалить</a>
                </li>
            {{end}}
        </ul>

        <h2>Добавить новую задачу</h2>
        <form action="/addtask" method="post">
            <input type="text" name="content" placeholder="Новая задача">
            <button type="submit">Добавить</button>
        </form>

        <h2>Профиль</h2>
        <p>Привет, {{.User.Username}}!</p>
        <a href="/logout">Выйти</a>
    {{end}}
</body>
</html>
