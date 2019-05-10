## Features

The `ff` does not allow command if not allowed according to the git flow. For example, no hotfix/feature branches can be created if current branch is an hotfix/feature branch in turn.

An hotfix/feature branch can be created only from master.

In case of LTS (a minor version) after each new tag merge updates into master to keep development version updated with all hotfixes.

Logs are stored in .git/logger.log file

Tag directly from master.

Create git repository if not exists.

Undo last commit.

Any git command's output is logged in *.git/logger.log*. Logs may be disabled or not using configuration.

Fetch all branches and tags.
