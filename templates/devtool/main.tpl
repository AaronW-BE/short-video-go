{{ define "devtool/main.tpl" }}
    {{template "header"}}
        <div class="sys-info-container">
            <div class="info-item">
                <div class="data">{{.sysInfo.CpuAmount}}</div>
                <div class="caption">核心数</div>
            </div>
            <div class="info-item">
                <div class="data">{{.sysInfo.Mem.Total}}</div>
                <div class="caption">内存(字节)</div>
            </div>
        </div>
    {{template "footer" }}

    <style>
        .sys-info-container {
            display: flex;
        }
        .info-item {
            padding: 20px;
            text-align: center;
            margin-left: 5rem;
        }
        .info-item .data {
            font-size: 1.3rem;
            font-weight: bold;
        }

    </style>
{{ end }}