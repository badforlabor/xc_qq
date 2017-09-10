:: 设置项目名字

:: 设置编译32位程序，炫彩界面库的DLL是32位的，所以exe也只能是32位的。
set GOARCH=386
::set GOGCCFLAGS=-m32 -mthreads -fmessage-length=0

set exename=xc_qq

:: 依赖第三方库：go get github.com/akavel/rsrc
:: 需要手动创建一个manifest文件，名字叫：%exename%.exe.manifest
:: 需要手动创建一个icon文件，名字叫：%exename%.ico


rsrc -manifest %exename%.exe.manifest -ico %exename%.ico -o %exename%.syso
go build -o %exename%.exe -ldflags="-H windowsgui -linkmode internal"
del %exename%.syso

pause