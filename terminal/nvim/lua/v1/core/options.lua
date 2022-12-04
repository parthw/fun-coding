local opt = vim.opt

-- line numbers

opt.number = true

-- tabs & indentation

opt.tabstop = 2
opt.shiftwidth = 2
opt.expandtab = true
opt.autoindent = true

-- line wrapping

opt.wrap = false

-- search settings

opt.ignorecase = true
opt.smartcase = true

-- cursor line

opt.cursorline = true

-- appearance

opt.termguicolors = true
opt.background = "dark"
opt.signcolumn = "yes"

-- backspace key

opt.backspace = "indent,eol,start"

-- clipboard

opt.clipboard:append("unnamedplus")

-- split windows

opt.splitright = true
opt.splitbelow = true

opt.iskeyword:append("-")
