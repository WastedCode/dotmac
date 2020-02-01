set nocompatible              " be iMproved, required
filetype off                  " required

set rtp+=~/.dotmac/vim/bundle/Vundle.vim
call vundle#begin()

" let Vundle manage Vundle, required
Plugin 'VundleVim/Vundle.vim'

" Go Plugins
Plugin 'fatih/vim-go'

" Git
Plugin 'tpope/vim-fugitive'
Plugin 'airblade/vim-gitgutter'

" General usability
Plugin 'ctrlpvim/ctrlp.vim'
Plugin 'lifepillar/vim-solarized8'
Plugin 'vim-airline/vim-airline'
Plugin 'dense-analysis/ale'
Plugin 'vim-syntastic/syntastic'

call vundle#end()
filetype plugin indent on  

set number
set showcmd                                 
set showmode
set visualbell
set tabstop=2 shiftwidth=2 expandtab
set autoindent
set smartindent
set smarttab

let g:go_def_mode='gopls'
let g:go_info_mode='gopls'

set background=dark
syntax on
colorscheme solarized8

let g:ale_linters = {
\   'go': ['golangci-lint'],
\}
