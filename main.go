package main

import (
  "fmt"
  "log"
  "os/exec"
  "os/user"
  "regexp"
  "strings"
  "path/filepath"

  "github.com/bgentry/go-netrc/netrc"
  "github.com/google/go-github/github"
  "golang.org/x/oauth2"
)

var regexps = []*regexp.Regexp{
  regexp.MustCompile(`^git@[^:]+:(?P<owner>[^/]+)/(?P<repo>.+?)(?:\.git)?$`),
  regexp.MustCompile(`^ssh://git@[^/]+/(?P<owner>[^/]+)/(?P<repo>.+?)(?:\.git)?$`),
  regexp.MustCompile(`^https://[^/]+/(?P<owner>[^/]+)/(?P<repo>.+?)(?:\.git)?$`),
}

func parseURL(url string) (string, string) {
  for _, r := range regexps {
    m := r.FindStringSubmatch(url)
    if len(m) > 0 {
      return m[1], m[2]
    }
  }
  log.Fatal("not github repo")
  return "", ""
}

func getRemoteURL() (string) {
  out, err := exec.Command("git", "remote", "get-url", "origin").Output()
  if err != nil {
    log.Fatal(err)
  }
  return strings.TrimSpace(string(out))
}

func getAccessToken() (string) {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }
  m, err := netrc.FindMachine(filepath.Join(usr.HomeDir, ".netrc"), "api.github.com")
  if err != nil {
    log.Fatal(err)
  }
  return m.Password
}

func main() {
  accessToken := getAccessToken()

  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
  tc := oauth2.NewClient(oauth2.NoContext, ts)
  client := github.NewClient(tc)

  url := getRemoteURL()
  owner, repo := parseURL(url)
  pulls, _, err := client.PullRequests.List(owner, repo, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, p := range pulls {
    b := *p.Head
    fmt.Printf("#%v %v\t%v\n", *p.Number, *p.Title, *b.Ref)
  }
}
