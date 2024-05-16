<!DOCTYPE html>
<html>
<head>
    <title>Авторизация</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f2f2f2;
            text-align: center;
            margin-top: 50px;
        }

        .login-container {
            background-color: white;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            max-width: 400px;
            margin: 0 auto;
        }

        input[type="text"], input[type="password"] {
            width: 100%;
            padding: 12px;
            margin: 8px 0;
            display: inline-block;
            border: 1px solid #ccc;
            border-radius: 4px;
            box-sizing: border-box;
        }

        button {
            background-color: #4CAF50;
            color: white;
            padding: 14px 20px;
            margin: 8px 0;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            width: 100%;
        }

        button:hover {
            background-color: #45a049;
        }

        .error-message {
            color: red;
            font-weight: bold;
            margin-bottom: 10px;
        }
    </style>
</head>
<body>
    <div class="login-container">
        <h2>Авторизация</h2>
        {{if .ErrorMessage}}
            <p class="error-message">{{.ErrorMessage}}</p>
        {{end}}
        <form action="/login" method="post">
            <input type="text" name="username" placeholder="Имя пользователя" required>
            <br>
            <input type="password" name="password" placeholder="Пароль" required>
            <br>
            <button type="submit">Войти</button>
        </form>
        <p>Нет аккаунта? <a href="/register">Зарегистрироваться</a></p>
    </div>
</body>
</html>
