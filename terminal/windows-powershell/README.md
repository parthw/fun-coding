### Steps

1. put file user_profile.ps1 in ~\.config\powershell\user_profile.ps1
2. execute vim $PROFILE.CurrentUserCurrentHost and put the text - 
```
. $env:USERPROFILE\.config\powershell\user_profile.ps1
```
3. copy text from powershell-settings.json and put in your settings.json of powershell