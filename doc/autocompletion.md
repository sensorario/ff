## Autocompletion

Append the following lines in your .bash_profile file:

```bash
_ff()
{
    local cur prev

    cur=${COMP_WORDS[COMP_CWORD]}
    prev=${COMP_WORDS[COMP_CWORD-1]}

    case ${COMP_CWORD} in
        1)
            COMPREPLY=($(compgen -W "patch pull conf config authors undo tag commit complete feature help hotfix bugfix publish refactor reset status" -- ${cur}))
            ;;
        2)
            case ${prev} in
                config)
                    COMPREPLY=($(compgen -W "tagAfterMerge disableUndoCommand stopAskingForTags applyFirstTag enableGitCommandLog forceOnPublish pushTagsOnPublish" -- ${cur}))
                    ;;
            esac
            ;;
        *)
            COMPREPLY=()
            ;;
    esac
}

complete -F _ff ff
```
