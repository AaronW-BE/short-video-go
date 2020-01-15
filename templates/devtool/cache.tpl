{{ define "devtool/cache.tpl"}}
    {{template "header" .}}

    <h1>Devtool - Cache</h1>
    <div>
        共有 {{len .caches}} 个 cache
    </div>
    <div class="">
        <table class=" cache-table table table-bordered">
            <thead>
                <tr class="table-header">
                    <td>Key</td>
                    <td>Value</td>
                    <td>Expiration</td>
                </tr>
            </thead>
            <tbody>
            {{range $k, $v := .caches}}
                <tr>
                    <td>{{$k}}</td>
                    <td>{{$v.Object}}</td>
                    <td>{{$v.Expiration}}</td>
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