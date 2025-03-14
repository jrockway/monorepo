package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v32/github"
	"github.com/jessevdk/go-flags"
	"golang.org/x/oauth2"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type auth struct {
	AccessToken    string `long:"token" env:"GITHUB_TOKEN" description:"If authenticating as a user, the Personal Access Token to use to access Github."`
	AppID          int64  `long:"app-id" env:"GITHUB_APP_ID" description:"If authenticating as a Github app, the App ID provided by Github."`
	InstallationID int64  `long:"installation-id" env:"GITHUB_INSTALLATION_ID" description:"If authenticating as a Github app, the Installation ID provided by Github."`
	PrivateKey     string `long:"private-key" env:"GITHUB_PRIVATE_KEY" description:"If authenticating as a Github app, the full private key provided by Github."`
}

type config struct {
	Timeout       time.Duration `long:"timeout" description:"How long to wait for Github." default:"30s"`
	GithubOwner   string        `long:"owner" description:"The owner of the repository to edit."`
	GithubRepo    string        `long:"repo" description:"The repository to edit."`
	GithubBranch  string        `long:"branch" description:"The branch to edit."`
	File          string        `long:"file" description:"The file to edit."`
	Locations     []string      `long:"location" description:"The location in the YAML file to replace.  Repeatable."`
	Replacement   string        `long:"replacement" description:"The content to replace the text at the provided locations with."`
	DryRun        bool          `long:"dry-run" description:"Print the diff of the edit we would like to commit, rather than committing it."`
	AuthorName    string        `long:"author-name" description:"The full name of the user that will generate the commit."`
	AuthorEmail   string        `long:"author-email" description:"The email address of the user that will generate the commit."`
	CommitMessage string        `long:"message" description:"The desired text of the commit message."`
}

type fileInTree struct {
	Tree      *github.Tree
	CommitSHA string
	Content   string
}

func fetch(ctx context.Context, client *github.Client, owner, repo, branch, file string) (*fileInTree, error) {
	br, _, err := client.Repositories.GetBranch(ctx, owner, repo, branch)
	if err != nil {
		return nil, fmt.Errorf("get branch: %w", err)
	}
	commit := br.GetCommit()
	if commit == nil {
		return nil, errors.New("no commit on branch")
	}
	treeRef := commit.GetCommit().GetTree().GetSHA()
	if treeRef == "" {
		return nil, fmt.Errorf("no tree in commit %s", commit)
	}
	tree, _, err := client.Git.GetTree(ctx, owner, repo, treeRef, true)
	if err != nil {
		return nil, fmt.Errorf("fetch tree %s from commit %s: %w", treeRef, commit, err)
	}
	if tree.Truncated == nil || tree.GetTruncated() {
		return nil, fmt.Errorf("github truncated tree %s, aborting", treeRef)
	}
	var blobSHA string
	for _, e := range tree.Entries {
		if e.GetPath() == file {
			blobSHA = e.GetSHA()
			break
		}
	}
	if blobSHA == "" {
		return nil, fmt.Errorf("file not found in commit %s", commit)
	}

	blob, _, err := client.Git.GetBlob(ctx, owner, repo, blobSHA)
	if err != nil {
		return nil, fmt.Errorf("fetch blob %s: %w", blobSHA, err)
	}

	var content string
	if e, c := blob.GetEncoding(), blob.GetContent(); e == "utf-8" {
		content = c
	} else if e == "base64" {
		c, err := base64.StdEncoding.DecodeString(c)
		if err != nil {
			return nil, fmt.Errorf("decode blob %s: %w", blobSHA, err)
		}
		content = string(c)
	} else {
		return nil, fmt.Errorf("unknown content type %q", file)
	}
	return &fileInTree{
		Tree:      tree,
		CommitSHA: commit.GetSHA(),
		Content:   content,
	}, nil
}

func editYAML(input string, locations []string, replacement string) (string, error) {
	nodes, err := yaml.Parse(input)
	if err != nil {
		return "", fmt.Errorf("parse yaml: %w", err)
	}

	var filters []yaml.Filter
	for _, location := range locations {
		path := strings.Split(location, ".")
		filters = append(filters, yaml.Tee(yaml.Lookup(path...), yaml.Set(yaml.NewScalarRNode(replacement))))
	}
	if _, err := nodes.Pipe(filters...); err != nil {
		return "", fmt.Errorf("apply edits: %w", err)
	}
	out, err := nodes.String()
	if err != nil {
		return "", fmt.Errorf("format yaml: %w", err)
	}
	return out, nil
}

func commit(ctx context.Context, client *github.Client, baseTreeSHA, baseCommit, owner, repo, branch, filename, content, commitMsg string, author *github.CommitAuthor) (string, error) {
	contentType := "base64"
	base64Content := base64.StdEncoding.EncodeToString([]byte(content))
	blob, _, err := client.Git.CreateBlob(ctx, owner, repo, &github.Blob{
		Encoding: &contentType,
		Content:  &base64Content,
	})
	if err != nil {
		return "", fmt.Errorf("create blob: %w", err)
	}
	mode := "100644"
	entryType := "blob"
	tree, _, err := client.Git.CreateTree(ctx, owner, repo, baseTreeSHA, []*github.TreeEntry{{
		Path: &filename,
		Mode: &mode,
		Type: &entryType,
		SHA:  blob.SHA,
	}})
	if err != nil {
		return "", fmt.Errorf("create tree with blob %s: %w", blob.GetSHA(), err)
	}

	now := time.Now()
	author.Date = &now
	commit, _, err := client.Git.CreateCommit(ctx, owner, repo, &github.Commit{
		Author:    author,
		Committer: author,
		Message:   &commitMsg,
		Parents:   []*github.Commit{{SHA: &baseCommit}},
		Tree:      tree,
	})
	if err != nil {
		return "", fmt.Errorf("create commit from tree %s and parent %s: %w", tree.GetSHA(), baseCommit, err)
	}
	head := fmt.Sprintf("heads/%s", branch)
	_, _, err = client.Git.UpdateRef(ctx, owner, repo, &github.Reference{Ref: &head, Object: &github.GitObject{SHA: commit.SHA}}, false)
	if err != nil {
		return "", fmt.Errorf("move %s to commit %s: %w", head, commit.GetSHA(), err)
	}
	return commit.GetSHA(), nil
}

func main() {
	var cfg config
	var auth auth

	fp := flags.NewParser(nil, flags.HelpFlag|flags.PassDoubleDash)
	if _, err := fp.AddGroup("Authentication", "", &auth); err != nil {
		panic(err)
	}
	if _, err := fp.AddGroup("Configuration", "", &cfg); err != nil {
		panic(err)
	}
	if _, err := fp.Parse(); err != nil {
		if ferr, ok := err.(*flags.Error); ok && ferr.Type == flags.ErrHelp {
			fmt.Fprintf(os.Stderr, "%s", ferr.Message)
			os.Exit(2)
		}
		fmt.Fprintf(os.Stderr, "flag parsing: %v\n", err)
		os.Exit(3)
	}

	ctx, c := context.WithTimeout(context.Background(), cfg.Timeout)
	defer c()

	var client *github.Client
	if auth.AppID != 0 && auth.InstallationID != 0 && len(auth.PrivateKey) > 0 {
		log.Println("Authenticating to Github as an app installation")
		tr := http.DefaultTransport
		itr, err := ghinstallation.New(tr, auth.AppID, auth.InstallationID, []byte(auth.PrivateKey))
		if err != nil {
			log.Fatalf("new github apps key: %v", err)
		}
		client = github.NewClient(&http.Client{Transport: itr})
	} else if auth.AccessToken != "" {
		log.Println("Authenticating to Github with a token")
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: auth.AccessToken})
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		log.Fatal("no authentication credentials provided")
	}

	orig, err := fetch(ctx, client, cfg.GithubOwner, cfg.GithubRepo, cfg.GithubBranch, cfg.File)
	if err != nil {
		log.Fatalf("fetch %s from github.com/%s/%s@%s: %v", cfg.File, cfg.GithubOwner, cfg.GithubRepo, cfg.GithubBranch, err)
	}

	new, err := editYAML(orig.Content, cfg.Locations, cfg.Replacement)
	if err != nil {
		log.Fatalf("replace content at locations %#v with %q in file %s: %v", cfg.Locations, cfg.Replacement, cfg.File, err)
	}
	if cfg.DryRun {
		fmt.Fprintf(os.Stderr, "Using content from commit %s\n", orig.CommitSHA)
		fmt.Print(new)
		os.Exit(0)
	}

	author := &github.CommitAuthor{
		Email: &cfg.AuthorEmail,
		Name:  &cfg.AuthorName,
	}
	sha, err := commit(ctx, client, orig.Tree.GetSHA(), orig.CommitSHA, cfg.GithubOwner, cfg.GithubRepo, cfg.GithubBranch, cfg.File, new, cfg.CommitMessage, author)
	if err != nil {
		log.Fatalf("commit new yaml: %v", err)
	}

	log.Printf("created commit %s", sha)
}
