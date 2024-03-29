if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi
source /etc/zprofile
export ZSH="$HOME/.oh-my-zsh"
ZSH_THEME="powerlevel10k/powerlevel10k"
plugins=(git zsh-z zsh-syntax-highlighting zsh-autosuggestions aws terraform poetry kubectl helm docker-compose web-search)
source $ZSH/oh-my-zsh.sh
export TERM=xterm-256color
export GREP_OPTIONS='--color=auto' GREP_COLOR='1;32'
export CLICOLOR=1
export LSCOLORS=exfxcxdxbxegedabagacad
setopt interactivecomments
export OBJC_DISABLE_INITIALIZE_FORK_SAFETY=YES
export GPG_TTY=/dev/ttys000
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh

export GOPATH="$HOME/go"
export GO111MODULE="on"
export SYSTEM_PYTHON_BIN="/Users/parthwadhwa/Library/Python/3.9/bin"
export PATH="$HOME/bin:$GOPATH/bin:/usr/local/bin:$SYSTEM_PYTHON_BIN:$PATH"

export PYENV_ROOT="$HOME/.pyenv"
command -v pyenv >/dev/null || export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init -)"
export PATH="$HOME/.poetry/bin:$PATH"
export WORKON_HOME=$HOME/.virtualenvs
export POETRY_VIRTUALENVS_PATH=$HOME/.virtualenvs
export CPPFLAGS="-I/usr/local/opt/openjdk@11/include"
export LC_CTYPE=en_US.UTF-8
export LC_ALL=en_US.UTF-8

alias docker_clean_containers='docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q)'
alias busybox='kubectl run -i --tty busybox --image=busybox --restart=Never --rm -- sh'
alias cat='bat'
alias nv='nvim'
alias sf='fd --full-path /Users/parthwadhwa/Documents/ --type f --max-depth 8 --exclude .git | fzf --height 30% --layout=reverse --border | xargs nvim'
alias sd='fd --full-path /Users/parthwadhwa/Documents/ --type d --max-depth 8 --exclude .git | fzf --height 30% --layout=reverse --border | xargs nvim'
command -v lsd > /dev/null && alias ls='lsd --group-dirs first'
command -v lsd > /dev/null && alias tree='lsd --tree'
complete -C '/usr/local/bin/aws_completer' aws
complete -o nospace -C /usr/local/bin/terraform terraform
source <(kubectl completion zsh)
source <(helm completion zsh)
export TF_PLUGIN_CACHE_DIR="$HOME/.terraform.d/plugin-cache"

export KUBECTL_EXTERNAL_DIFF='delta --side-by-side --navigate'
source ~/.zshrc_personal
