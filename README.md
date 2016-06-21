# prf

**prf** - *PullRequest Fetch* - lists open-pullrequests to stdout

## SYNOPSIS

```bash
prf
```

## DESCRIPTION

Is checking out a git branch for code review bothering you?
In general, you may open the PR's page then copy the branch name then back to terminal then checkout the branch.

Using `prf` with interactive filtering tools such as [peco](https://githuc.com/peco/peco) or [fzf](https://github.com/junegunn/fzf) provides you a solution for the problem (see [EXAMPLES](#examples) below).

## EXAMPLES

Checkout the git branch corresponding to a open-pullrequest interactivly using peco:

```bash
prf | peco | cut -f3 | xargs git checkout
```

## INSTALLATION

```bash
go get github.com/yuku-t/prf
```

Or clone the [repository](https://github.com/yuku-t/prf) and run:

```bash
make install
```

Then issue [a personal access token](https://github.com/settings/tokens/new) with "repo" scope and register it as "api.github.com" machine's password to netrc file:

```bash
cat <<-EOS >>~/.netrc
machine api.github.com
  password PERSONAL_ACCESS_TOKEN
EOS
chmod 600 ~/.netrc
```
