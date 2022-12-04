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
