local status, _ = pcall(vim.cmd, "colorscheme moonfly")
if not status then
  print("colorscheme not found")
  return
end

