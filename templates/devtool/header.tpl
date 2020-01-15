{{define "header"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>
    <style>
        html,body {
            margin: 0;
            padding: 0;
        }
        .header-block {
            padding: 20px;
            background-color: #3b9eff;
            display: flex;
            align-items: center;
        }
        .brand-title {
            color: #ffffff;
            font-size: 1.5rem;
        }
        .brand-logo {
            width: 40px;
        }
        .main-container {
            padding: 10px;
        }
    </style>
</head>
    <header>
        <div class="header-block">
            <img class="brand-logo" src="/favicon.ico" alt="">
            <div class="brand-title">开发控制台</div>
        </div>
    </header>
    <body>
    <div class="main-container">
{{end}}