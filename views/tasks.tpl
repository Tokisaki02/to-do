<!DOCTYPE html>
<html>
<head>
    <title>Мои задачи</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f2f2f2;
        }

        .container {
            background-color: white;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            max-width: 800px;
            margin: 50px auto;
        }

        h1, h2 {
            color: #333;
            margin-bottom: 10px;
        }

        ul {
            list-style: none;
            padding: 0;
            margin: 0;
        }

        li {
            padding: 10px;
            border-bottom: 1px solid #ccc;
        }

        li:last-child {
            border-bottom: none;
        }

        a {
            color: #337ab7;
            text-decoration: none;
        }

        a:hover {
            color: #23527c;
        }

        form {
            margin-top: 20px;
        }

        input[type=text] {
            width: 100%;
            padding: 12px 20px;
            margin: 8px 0;
            display: inline-block;
            border: 1px solid #ccc;
            border-radius: 4px;
            box-sizing: border-box;
        }

        button[type=submit] {
            background-color: #4CAF50;
            color: white;
            padding: 14px 20px;
            margin: 8px 0;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            width: 100%;
        }

        button[type=submit]:hover {
            background-color: #45a049;
        }

        .profile {
            margin-top: 20px;
        }

        .profile a {
            color: #337ab7;
            text-decoration: none;
        }

        .profile a:hover {
            color: #23527c;
        }
    </style>
</head>
<body>
    <div class="container">
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

        <div class="profile">
            <h2>Профиль</h2>
            <p>Привет, {{.User.Username}}!</p>
            <a href="/logout">Выйти</a>
        </div>
    </div>
</body>
</html>
