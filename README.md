github-issues
=============

[![GitHub release](https://img.shields.io/github/release/i2bskn/github-issues.svg?style=flat-square)](https://github.com/i2bskn/github-issues/releases)
[![GitHub license](https://img.shields.io/github/license/i2bskn/github-issues.svg?style=flat-square)](https://github.com/i2bskn/github-issues/blob/master/LICENSE.txt)

List of GitHub issues.

Installation
------------

Download from [releases](https://github.com/i2bskn/github-issues/releases) and stored in the `$PATH`.  

Or with homebrew:

```
brew update
brew tap i2bskn/i2bskn
brew install github-issues
```

Settings
--------

Required [Personal Access Token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) of GitHub.  
Personal Access Token must be set in in one of the following ways.

#### .gitconfig

Add the following settings to `.gitconfig`.

```
[github]
  token = <personal access token>
```

#### Environments

Set to the environment `GITHUB_TOKEN`.

```
GITHUB_TOKEN=<personal access token> github-issues
```

#### Command line argument

```
github-issues -token <personal access token>
```

Usage
-----

Just run the following command!

```
github-issues
```

Options
-------

### Repository

|Option|Type|Default|Description|
|------|----|-------|-----------|
|-current|Bool|false|Current repository only|
|-repo|String|empty|Specific repository only|
|-self|Bool|false|Your own repositories only|

### Pagination

|Option|Type|Default|Description|
|------|----|-------|-----------|
|-p|int|1|Specify further pages|
|-n|int|100|Specify a custom page size|

### Filter

|Option|Type|Default|Description|
|------|----|-------|-----------|
|-a|Bool|none|Refine issues assigned to you|
|-c|Bool|none|Refine issues created by you|
|-m|Bool|none|Refine issues mentioning you|

### State

Specify the state of the issues to display in `-state`.  
Can be either `open`, `closed`, `all`.

### Sort

Specify the sort of the issues to display in `-sort`.  
Can be either `created`, `updated`, `comments`.

### Format

You can specify the format in `-format`.

|Symbol|Description|
|------|-----------|
|%n|Issue number|
|%l|URL|
|%t|Title|
|%u|User|

Default format: `%n\t%l\t%t\t%u`

### Option Specified Example

```
github-issue -repo -n 5 -a -format "%t"
```
