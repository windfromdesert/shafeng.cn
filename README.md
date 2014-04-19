## My blog created code

使用语言：Golang

### 使用说明

+   /archive-----存档文件夹
+   /css---------css文件夹
+   /go----------golang 代码文件夹
+   /htm---------htm网页文件夹
+   /image-------图片文件夹
+   /mo----------模板文件夹
+   /txt---------网页 Markdown 格式备份文件夹
+   about--------关于网页
+   link---------友情链接网页
+   post.txt-----发表文章模板[使用markdown少量格式]

### /go 文件夹代码使用说明

archive.go

+   自动创建存档页代码，存档页创建并保存于 '/archive' 文件夹中
+   按照第 17, 18 行代码格式创建新增标签

index.go

+   自动创建网站主页代码，即根目录中的 index.html
+   在主页上将显示最后发表的10篇文章

post.go

+   自动发表文章代码，即在 '/htm' 文件夹中将生成新网页文件
