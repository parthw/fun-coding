vim.g.mapleader = " "

local keymap = vim.keymap

-- clear search highlights in normal mode
keymap.set("n", "<leader>nh", ":nohl<CR>")
-- don't save delete from x to vim register
keymap.set("n", "x", '"_x')
-- split windows vertically
keymap.set("n", "<leader>sv", "<C-w>v")
-- split windows horizontally
keymap.set("n", "<leader>sh", "<C-w>s")
-- split windows equal width
keymap.set("n", "<leader>se", "<C-w>=")
-- close current split window
keymap.set("n", "<leader>sx", ":close<CR>")
-- open new tab
keymap.set("n", "<leader>to", ":tabnew<CR>")
-- close current tab
keymap.set("n", "<leader>tx", ":tabclose<CR>")
-- go to next tab
keymap.set("n", "<leader>tn", ":tabn<CR>")
-- go to prev tab
keymap.set("n", "<leader>tp", ":tabp<CR>")
-- vim-maximizer: toggle split window maximization
keymap.set("n", "<leader>sm", ":MaximizerToggle<CR>")
-- nvim-tree: toogle file explorer
keymap.set("n", "<leader>e", ":NvimTreeToggle<CR>")
-- telescope
keymap.set("n", "<leader>ff", "<cmd>Telescope find_files<cr>") -- find files within current working directory, respects .gitignore
keymap.set("n", "<leader>fs", "<cmd>Telescope live_grep<cr>") -- find string in current working directory as you type
keymap.set("n", "<leader>fc", "<cmd>Telescope grep_string<cr>") -- find string under cursor in current working directory
keymap.set("n", "<leader>fb", "<cmd>Telescope buffers<cr>") -- list open buffers in current neovim instance
keymap.set("n", "<leader>fh", "<cmd>Telescope help_tags<cr>") -- list available help tags
keymap.set("n", "<leader>gc", "<cmd>Telescope git_commits<cr>") -- list all git commits (use <cr> to checkout) ["gc" for git commits]
keymap.set("n", "<leader>gfc", "<cmd>Telescope git_bcommits<cr>") -- list git commits for current file/buffer (use <cr> to checkout) ["gfc" for git file commits]
keymap.set("n", "<leader>gb", "<cmd>Telescope git_branches<cr>") -- list git branches (use <cr> to checkout) ["gb" for git branch]
keymap.set("n", "<leader>gs", "<cmd>Telescope git_status<cr>") -- list current changes per file with diff preview ["gs" for git status]

