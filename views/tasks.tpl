<!-- views/tasks.tpl -->
<!DOCTYPE html>
<html>
<head>
    <title>Мои задачи</title>
</head>
<body>
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
</body>
</html>
