github-issues
=============

List of GitHub issues. (by all repository)

Installation
------------

With go lang:

```
go get github.com/i2bskn/github-issues
```

Settings
--------

Required [Personal Access Token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/) of GitHub.  
Personal Access Token must be able to get in the `git-config` command.

```
[github]
  token = <personal access token>
```

Usage
-----

Just run the following command!

```
github-issues
```

Options
-------

### Pagination

|Long|Short|Default|Description|
|----|-----|-------|-----------|
|--page|-p|1|Specify further pages|
|--per-page|-n|100|Specify a custom page size|

### Filter

|Long|Short|Default|Description|
|----|-----|-------|-----------|
|--assigned|-a|none|Issues assigned to you|
|--created|-c|none|Issues created by you|
|--mentioned|-m|none|Issues mentioning you|
|--closed|none|none|Closed issue only|
|--all|none|none|With closed issue|

### Format

You can specify the format in `--format` or `-f`.

|Symbol|Description|
|------|-----------|
|%n|Issue number|
|%l|URL|
|%t|Title|
|%u|User|

Default format: `%n\t%l\t%t\t%u`

### Option Example

```
github-issue -n 5 -a --format="%t"
```
