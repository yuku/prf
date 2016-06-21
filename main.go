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

func main() {
  out, err := exec.Command("git", "remote", "get-url", "origin").Output()
  if err != nil {
    log.Fatal(err)
  }
  url := strings.TrimSpace(string(out))
  var owner string
  var repo string
  r1 := regexp.MustCompile(`^git@[^:]+:(?P<owner>[^/]+)/(?P<repo>.+?)(?:\.git)?$`)
  m1 := r1.FindStringSubmatch(url)
  if len(m1) > 0 {
    owner = m1[1]
    repo = m1[2]
  } else {
    r2 := regexp.MustCompile(`^ssh://git@[^/]+/(?P<owner>[^/]+)/(?P<repo>.+?)(?:\.git)?$`)
    m2 := r2.FindStringSubmatch(url)
    if len(m2) > 0 {
      owner = m2[1]
      repo = m2[2]
    } else {
      r3 := regexp.MustCompile(`^https://[^/]+/(?P<owner>[^/]+)/(?P<repo>.+?)(?:\.git)?$`)
      m3 := r3.FindStringSubmatch(url)
      if len(m3) > 0 {
        owner = m3[1]
        repo = m3[2]
      } else {
        log.Fatal("not github repo")
      }
    }
  }

  usr, err := user.Current()
  if err != nil {
    log.Fatal( err )
  }
  m, _ := netrc.FindMachine(filepath.Join(usr.HomeDir, ".netrc"), "api.github.com")

  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: m.Password},
  )
  tc := oauth2.NewClient(oauth2.NoContext, ts)
  client := github.NewClient(tc)

  pulls, _, err := client.PullRequests.List(owner, repo, nil)
  if err != nil {
    log.Fatal(err)
  }
  for _, p := range pulls {
    b := *p.Head
    fmt.Printf("#%v %v\t%v\n", *p.Number, *p.Title, *b.Ref)
  }
}
