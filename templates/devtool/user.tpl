{{ define "devtool/user.tpl"}}
    {{template "header" .}}

    <h1>Devtool - User</h1>
    <div>
        共有 {{len .users}} 个 用户
    </div>
    <div class="">
        <table class=" cache-table table table-bordered">
            <thead>
                <tr class="table-header">
                    <td>Index</td>
                    <td>Username</td>
                    <td>Phone</td>
                    <td>State</td>
                    <td>Reg Date</td>
                    <td>Last Login</td>
                </tr>
            </thead>
            <tbody>
            {{range $k, $v := .users}}
                <tr>
                    <td>{{$k}}</td>
                    <td>{{$v.Username}}</td>
                    <td>{{$v.Phone}}</td>
                    <td>{{$v.State}}</td>
                    <td>{{$v.Register}}</td>
                    <td>{{$v.LastLogin}}</td>
                </tr>
            {{end}}
            </tbody>
        </table>
    </div>

    <style>
        .cache-table {
            table-layout: fixed;
        }
        .cache-table td {
            white-space: pre-wrap;
            /*overflow: hidden;*/
            word-break: break-all;
        }
    </style>
    {{template "footer"}}
{{end}}