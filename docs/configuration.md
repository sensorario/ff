## Configuration

 - change name of development branch
 - disable auto-tag whenever a support branch is merged
 - disable undo command
 - ask if user want to tag
 - stop asking for auto-tag
 - enable/disable git command output log
 - git push may be forced or not
 - git push with tags
 - remove merged branches already in origin

```json
{
  "features": {
    "applyFirstTag": false,
    "disableUndoCommand": false,
    "enableGitCommandLog": true,
    "forceOnPublish": false,
    "pushTagsOnPublish": false,
    "removeRemotelyMerged": false,
    "stopAskingForTags": false,
    "tagAfterMerge": true
  },
  "branches": {
    "historical": {
      "development": "master"
    }
  }
}
```
