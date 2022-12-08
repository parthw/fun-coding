-- Ensure packer is installed
-- copied from https://github.com/wbthomason/packer.nvim#bootstrapping

local ensure_packer = function()
  local fn = vim.fn
  local install_path = fn.stdpath('data')..'/site/pack/packer/start/packer.nvim'
  if fn.empty(fn.glob(install_path)) > 0 then
    fn.system({'git', 'clone', '--depth', '1', 'https://github.com/wbthomason/packer.nvim', install_path})
    vim.cmd [[packadd packer.nvim]]
    return true
  end
  return false
end
local packer_bootstrap = ensure_packer()

-- Autocommand that reloads neovim when this file is saved 
vim.cmd([[
  augroup packer_user_config
    autocmd!
    autocmd BufWritePost plugins-setup.lua source <afile> | PackerSync
  augroup end
]])

return require('packer').startup(function(use)
  use 'wbthomason/packer.nvim'
  -- My plugins here
  use 'nvim-lua/plenary.nvim' -- lua functions that many plugins use
  use 'bluz71/vim-moonfly-colors' -- prefered coloscheme
  use 'christoomey/vim-tmux-navigator' -- keymap navigation b/w splits
  use 'szw/vim-maximizer' -- maximizes and restores current window
  use 'numToStr/Comment.nvim' -- commenting with gc
  use 'nvim-tree/nvim-tree.lua' -- file explorer
  use 'nvim-tree/nvim-web-devicons' -- nvim-tree icons
  use 'nvim-lualine/lualine.nvim' -- statusline plugin
  use { 'nvim-telescope/telescope-fzf-native.nvim', run = 'make' } -- dependency for better sorting performance
  use { 'nvim-telescope/telescope.nvim', branch = '0.1.x' } -- fuzzy finder
  use 'lewis6991/gitsigns.nvim' -- show git line modifications on left hand side

   -- configuring code-completion and lsp servers
  use 'hrsh7th/nvim-cmp' -- completion plugin
  use 'hrsh7th/cmp-buffer' -- source for text in buffer
  use 'hrsh7th/cmp-path' -- source for file system paths
  use 'L3MON4D3/LuaSnip' -- snippet engine
  use 'saadparwaiz1/cmp_luasnip' -- for autocompletion
  use 'rafamadriz/friendly-snippets' -- useful snippets
  use 'williamboman/mason.nvim' -- in charge of managing lsp servers, linters & formatters
  use 'williamboman/mason-lspconfig.nvim' -- bridges gap b/w mason & lspconfig
  use 'neovim/nvim-lspconfig' -- easily configure language servers
  use 'hrsh7th/cmp-nvim-lsp' -- for autocompletion
  use { 'glepnir/lspsaga.nvim', branch = 'main' } -- enhanced lsp uis
  use 'onsails/lspkind.nvim' -- vs-code like icons for autocompletion
  use 'jose-elias-alvarez/null-ls.nvim' -- configure formatters & linters
  use 'jayp0521/mason-null-ls.nvim' -- bridges gap b/w mason & null-ls
  use {
    'nvim-treesitter/nvim-treesitter',
    run = function()
      local ts_update = require('nvim-treesitter.install').update({ with_sync = true })
      ts_update()
    end,
  }
  use 'windwp/nvim-autopairs' -- autoclose parens, brackets, quotes, etc...
  use { 'windwp/nvim-ts-autotag', after = 'nvim-treesitter' } -- autoclose tags


  -- Automatically set up your configuration after cloning packer.nvim
  -- Put this at the end after all plugins
  if packer_bootstrap then
    require('packer').sync()
  end
end)

