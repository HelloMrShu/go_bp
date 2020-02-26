<!-- <!DOCTYPE html> -->
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <script type="text/javascript" src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    <link rel="stylesheet" type="text/css" href="https://maxcdn.bootstrapcdn.com/bootswatch/3.3.5/cerulean/bootstrap.min.css">
    <title>Document</title>
</head>

<body>

    <div class="container">
        <div class="row">
            <div class="col-md-1">ID</div>
            <div class="col-md-3">名称</div>
        </div>
        {{range .}}
        <div class="row">
            <div class="col-md-1">{{.Id}}</div>
            <div class="col-md-3">{{.Name}}</div>
        </div>
        {{end}}
    </div>
</body>
</html>