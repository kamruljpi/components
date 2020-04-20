package theme1

var List = map[string]string{"login/theme1": `{{define "login_theme1"}}
    <!DOCTYPE html>
    <html>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">

        <title>{{.Title}}</title>
        <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">

        <link rel="stylesheet" href="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.css"}}">

        <style>
            body.login-page {
                background-color: {% .BackgroundColor %};
                height: 80%;
            }

            div.login-logo a {
                color: white;
            }

            .text-center.text-muted {
                color: white;
            }

            .text-center.text-muted a {
                color: #92adce;
            }

            button.btn.btn-flat {
                background-color: {% .LoginBtnColor %};
            }
        </style>

    </head>
    <body class="hold-transition login-page" data-gr-c-s-loaded="true">
    <div class="login-box">
        <div class="login-logo">
            <a href="/"><b>{{.Title}}</b></a>
        </div>

        <div class="login-box-body">
            <form action="{{.UrlPrefix}}/signin" id="sign-in-form">
                <div class="form-group has-feedback 1">
                    <input type="text" class="form-control" placeholder="{{lang "username"}}" id="username">
                    <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
                </div>
                <div class="form-group has-feedback 1">
                    <input type="password" class="form-control" placeholder="{{lang "password"}}" id="password">
                    <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                </div>
                <div class="row">
                    <div class="col-xs-8">
                    </div>
                    <div class="col-xs-4">
                        <button type="submit" class="btn btn-primary btn-block btn-flat">{{lang "login"}}</button>
                    </div>
                </div>
            </form>

        </div>
    </div>

    <div class="text-center text-muted">
        <small>
            <strong>Powered by <a href="https://github.com/z-song/laravel-admin"
                                  target="_blank">GoAdmin</a></strong>
        </small>
    </div>

    </body>

    <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.js"}}"></script>
    <script>

        {% if .TencentWaterProofWallID  %}

        let captcha = new TencentCaptcha("{% .TencentWaterProofWallID %}", function (res) {
            console.log(res);
            // res（用户主动关闭验证码）= {ret: 2, ticket: null}
            // res（验证成功） = {ret: 0, ticket: "String", randstr: "String"}
            if (res.ret === 0) {
                $.ajax({
                    dataType: 'json',
                    type: 'POST',
                    url: '{{.UrlPrefix}}/signin',
                    async: 'true',
                    data: {
                        'username': $("#username").val(),
                        'password': $("#password").val(),
                        'token': res.ticket
                    },
                    success: function (data) {
                        location.href = data.data.url
                    },
                    error: function (data) {
                        alert('{{lang "login fail"}}');
                    }
                });
            } else {
                alert("验证失败")
            }
        }, {});

        {% end %}

        $("#sign-in-form").submit(function (e) {
            e.preventDefault();
            {% if .TencentWaterProofWallID  %}
            captcha.show();
            {% else  %}
            $.ajax({
                dataType: 'json',
                type: 'POST',
                url: '{{.UrlPrefix}}/signin',
                async: 'true',
                data: {
                    'username': $("#username").val(),
                    'password': $("#password").val()
                },
                success: function (data) {
                    location.href = data.data.url
                },
                error: function (data) {
                    alert("Login failed");
                }
            });
            {% end %}
        });
    </script>

    </html>
{{end}}`}
