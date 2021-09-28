@REM go build -ldflags="-H windowsgui" -o WINUI.exe
@REM 那么具体里面这段 -ldflags="-s -w -H windowsgui" 是什么意思呢？

@REM -ldflags=“参数”： 表示将引号里面的参数传给编译器
@REM -s：去掉符号信息  （这样panic时，stack trace就没有任何文件名/行号信息了，这等价于普通C/C+=程序被strip的效果）
@REM -w：去掉DWARF调试信息 （得到的程序就不能用gdb调试了）

@REM -H windowsgui : 以windows gui形式打包，不带dos窗口。其中注意H是大写的

@REM go build -ldflags="-s -w -H windowsgui" -o WINUI.exe

go build -ldflags="" -o WINUI.exe