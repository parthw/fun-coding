# Alias
Set-Alias ll ls
Set-Alias grep findstr

# fzf setup
$env:FZF_DEFAULT_OPTS='--height 70% --layout=reverse --border'
function ff {
  vim (fzf --preview "bat --style=numbers --color=always --line-range :500 {}")
}

# prompt
function Write-BranchName () {
  try {
    $branch = git rev-parse --abbrev-ref HEAD

    if ($branch -eq "HEAD") {
      # we're probably in detached HEAD state, so print the SHA
      $branch = git rev-parse --short HEAD
      Write-Host " ($branch)" -ForegroundColor "red"
    }
    else {
      # we're on an actual branch, so print it
      Write-Host " ($branch)" -ForegroundColor "green"
    }
  } catch {
    # we'll end up here if we're in a newly initiated git repo
    Write-Host " (no branches yet)" -ForegroundColor "yellow"
  }
}

function prompt {
  $base = "PS "
  $path = "$($executionContext.SessionState.Path.CurrentLocation)"
  $userPrompt = "$('>' * ($nestedPromptLevel + 1)) "

  Write-Host "`n$base" -NoNewline

  if (Test-Path .git) {
    Write-Host $path -NoNewline -ForegroundColor "blue"
    Write-BranchName
  }
  else {
    # we're not in a repo so don't bother displaying branch name/sha
    Write-Host $path -ForegroundColor "blue"
  }

  return $userPrompt
}