## Configuration

 - change name of development branch
 - disable auto-tag whenever a support branch is merged
 - disable undo command
 - ask if user want to tag
 - stop asking for auto-tag
 - enable/disable git command output log

```json
{
  "features": {
    "tagAfterMerge": true,
    "disableUndoCommand": false,
    "stopAskingForTags": false,
    "applyFirstTag": false,
    "enableGitCommandLog": true,
    "forceOnPublish": false,
    "pushTagsOnPublish": false
  },
  "branches": {
    "historical": {
      "development": "master"
    }
  }
}
```
