## Configuration

### Instructions

To toggle a configuration type `ff config <variable>`.
To set a configuration type `ff config <variable> true`. (or false)
To set a language type `ff config lang "it"`. (only it and en are available)

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
- set language of current configuration

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
    "tagAfterMerge": true,
    "lang": "it"
  },
  "branches": {
    "historical": {
      "development": "master"
    }
    "support": {
      "feature": "feat"
    }
  }
}
```

### Add new configuration instruction

First of all is mandatory to add the new feature in jsonConf struct located in `src/cong.go`. Then in method `readConfiguration` add default value of the feature. Finally, update `src/step_config.go` to treat the `ff config <feature>`.
