@echo off
setlocal enabledelayedexpansion

:: 编译单个Go文件脚本
:: 用法: go_build.bat filename.go 或拖动文件到脚本上

if "%~1"=="" (
    echo [错误] 请拖拽Go文件到脚本或指定参数
    echo 示例: go_build.bat main.go
    timeout /t 3 >nul
    exit /b 1
)

if not "%~x1"==".go" (
    echo [错误] 仅支持.go文件: %~nx1
    timeout /t 3 >nul
    exit /b 1
)

set "output=%~n1.exe"

echo.
echo 正在编译: %~nx1
echo 输出文件: %output%
echo 开始时间: %time%

go build -o "%output%" "%~1"

if errorlevel 1 (
    echo [失败] 编译错误: %~nx1
    timeout /t 5 >nul
    exit /b 1
) else (
    echo [成功] 编译完成: %output%
    echo 结束时间: %time%
)

endlocal