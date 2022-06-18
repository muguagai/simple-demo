# simple-demo

# 使用说明
1. 安装mumu安卓模拟器，https://mumu.163.com/index.html
2. 下载抖声apk，https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7
3. 进入抖声APP，双击右下角“我”配置服务器地址：http://本地IP:8080 
4. 按照该教程安装mysql数据库，并创建douyin数据库
5. 安装redis，启动redis持久化功能
6. 将respository/db_init.go中的"root:123@tcp"中的root和111111分别替换成你的mysql账户和密码
8. 运行命令  go test -v db_test.go db_init.go common.go user.go，否则会报undefined的错，错误原因参考：https://blog.csdn.net/love666666shen/article/details/119929929
10. 将controller/publish.go中的ip地址替换成自己本地的
11. 需要下载ffmpeg，https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-essentials.zip；之后解压，将bin目录配置到系统环境变量Path中
12. 进入项目根目录，运行go build && ./simple-demo（如果go build需要下载但是速度慢，可以运行“go env -w GOPROXY=https://goproxy.cn,direct”）
