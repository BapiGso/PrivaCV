package web

import "embed"

// 只嵌入运行时真正需要的资源，避免把源码或临时文件打进二进制。
//
//go:embed index.html favicon.jpg templates.json vendor content
var Assets embed.FS
