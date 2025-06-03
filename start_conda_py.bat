@echo off
REM 激活conda环境，这里environment 是我自己设置好的
call conda activate p13_learn

REM 检查是否成功激活环境
if errorlevel 1 (
    echo 错误：无法激活conda环境 p13_learn
    pause
    exit /b 1
)

REM 执行Python脚本
@echo on
python "%~1"
@echo off
