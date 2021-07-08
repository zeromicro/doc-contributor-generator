package templates

var Contributors = `
# 社区贡献

## 作者

<div style="width: 80px;height: auto;">
<a href="https://github.com/kevwan">
<img src="https://avatars.githubusercontent.com/u/1918356?s=460&u=3c40d5f0fc2d3db824b477ab3785db812ce331e1&v=4" alt="kevwan">
<br>
<center><div style="max-width:80px; overflow: hidden;text-overflow: ellipsis;white-space: nowrap">kevwan</div></center>
</a>
</div>

## go-zero参与人员

<div style="display: flex;flex-wrap: wrap">
{{range $index,$item := .goZeroList}}

<div style="width: auto;height: auto;margin: 5px">
<a href="{{$item.HTMLURL}}">
<img src="{{$item.AvatarURL}}" width="80px" height="80px" alt="{{$item.Login}}"/>
<center><div style="max-width:80px; overflow: hidden;text-overflow: ellipsis;white-space: nowrap">{{$item.Login}}</div></center>
</a>
</div>{{end}}
</div>

## 文档贡献人员
<div style="display: flex;flex-wrap: wrap">
{{range $index,$item := .goZeroDocList}}

<div style="width: auto;height: auto;margin: 5px">
<a href="{{$item.HTMLURL}}">
<img src="{{$item.AvatarURL}}" width="80px" height="80px" alt="{{$item.Login}}"/>
<center><div style="max-width:80px; overflow: hidden;text-overflow: ellipsis;white-space: nowrap">{{$item.Login}}</div></center>
</a>
</div>{{end}}
</div>

`

var Contributors_EN = `
# Community contribution

## Author

<div style="width: 80px;height: auto;">
<a href="https://github.com/kevwan">
<img src="https://avatars.githubusercontent.com/u/1918356?s=460&u=3c40d5f0fc2d3db824b477ab3785db812ce331e1&v=4" alt="kevwan">
<br>
<center><div style="max-width:80px; overflow: hidden;text-overflow: ellipsis;white-space: nowrap">kevwan</div></center>
</a>
</div>

## go-zero contributors

<div style="display: flex;flex-wrap: wrap">
{{range $index,$item := .goZeroList}}

<div style="width: auto;height: auto;margin: 5px">
<a href="{{$item.HTMLURL}}">
<img src="{{$item.AvatarURL}}" width="80px" height="80px" alt="{{$item.Login}}"/>
<center><div style="max-width:80px; overflow: hidden;text-overflow: ellipsis;white-space: nowrap">{{$item.Login}}</div></center>
</a>
</div>{{end}}
</div>

## Document contributors
<div style="display: flex;flex-wrap: wrap">
{{range $index,$item := .goZeroDocList}}

<div style="width: auto;height: auto;margin: 5px">
<a href="{{$item.HTMLURL}}">
<img src="{{$item.AvatarURL}}" width="80px" height="80px" alt="{{$item.Login}}"/>
<center><div style="max-width:80px; overflow: hidden;text-overflow: ellipsis;white-space: nowrap">{{$item.Login}}</div></center>
</a>
</div>{{end}}
</div>

`
