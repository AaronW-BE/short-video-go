{{define "header"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>

    <link href="https://cdn.bootcss.com/twitter-bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
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

        .links-container {
            margin-left: 20px;
        }
        .links-container .link {
            display: inline-block;
            color: white;
            margin-left: 20px;
            font-size: 20px;
        }
        a {
            color: white;
        }
    </style>
</head>
    <header>
        <div class="header-block">
            <img class="brand-logo" src="/favicon.ico" alt="">
            <div class="brand-title">开发控制台</div>
            <div class="links-container">
                <div class="link">
                    <a href="./main">Main</a>
                </div>
                <div class="link">
                    <a href="./user">User</a>
                </div>
                <div class="link">
                    <a href="./cache">Cache</a>
                </div>
            </div>
        </div>
    </header>
    <body>
    <div class="main-container">
{{end}}