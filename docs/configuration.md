## Configuration

### Instructions

To toggle a configuration type `ff config <variable>`.
To set a configuration type `ff config <variable> true`. (or false)

### Available confs

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
