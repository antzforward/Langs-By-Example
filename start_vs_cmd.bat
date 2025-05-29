@echo off
REM 设置 vs 编译环境
call "C:\Program Files\Microsoft Visual Studio\2022\Community\VC\Auxiliary\Build\vcvarsall.bat" x64

REM 编译当前文件
cl /nologo /EHsc /std:c++17 /Fe:"%~n1.exe" "%~1" /link benchmark.lib shlwapi.lib

REM 运行可执行文件（可选）
if %errorlevel% equ 0 (
    echo 运行程序:
    "%~n1.exe"
)
