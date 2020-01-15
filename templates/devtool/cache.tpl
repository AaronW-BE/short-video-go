{{ define "devtool/cache.tpl"}}
    {{template "header" .}}

    <h1>Devtool - Cache</h1>
    <div>
        共有 {{len .caches}} 个 cache
    </div>
    <div>
        <table border="1">
            <tr>
                <td>Key</td>
                <td>Value</td>
                <td>Expiration</td>
            </tr>
            {{range $k, $v := .caches}}
                <tr>
                    <td>{{$k}}</td>
                    <td>{{$v.Object}}</td>
                    <td>{{$v.Expiration}}</td>
                </tr>
            {{end}}
        </table>
    </div>
    {{template "footer"}}
{{end}}