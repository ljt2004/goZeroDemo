# Go-Zero 自定义模板快速开始脚本

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  Go-Zero 自定义模板配置" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 设置 GOCTL_HOME 环境变量
$templatePath = Join-Path $PWD "deploy\goctl\1.10.1"
Write-Host "模板路径: $templatePath" -ForegroundColor Yellow

# 设置当前会话的环境变量
$env:GOCTL_HOME = $templatePath
Write-Host "已设置 GOCTL_HOME=$templatePath" -ForegroundColor Green
Write-Host ""

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  使用示例" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "1. 生成 API 代码:" -ForegroundColor White
Write-Host "   goctl api go -api user.api -dir ." -ForegroundColor Gray
Write-Host ""
Write-Host "2. 生成 Model 代码:" -ForegroundColor White
Write-Host "   goctl model mysql ddl -src user.sql -dir ." -ForegroundColor Gray
Write-Host ""
Write-Host "3. 生成 RPC 代码:" -ForegroundColor White
Write-Host "   goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=./" -ForegroundColor Gray
Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "  永久设置环境变量（可选）" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "若要永久设置，请运行以下命令:" -ForegroundColor White
Write-Host "[System.Environment]::SetEnvironmentVariable('GOCTL_HOME', '$templatePath', 'User')" -ForegroundColor Gray
Write-Host ""
Write-Host "设置完成后请重启终端使配置生效" -ForegroundColor Yellow
Write-Host ""
