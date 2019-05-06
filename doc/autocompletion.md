## Autocompletion

Append the following lines in your .bash_profile file:

    _ff='patch pull conf authors undo tag commit complete feature help hotfix bugfix publish refactor reset status' && complete -W "${_ff}" 'ff'
