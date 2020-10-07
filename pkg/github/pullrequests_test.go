package github

import (
	"bytes"
	"context"
	"github.com/google/go-github/v32/github"
	"log"
	"net/http"
	"reflect"
	"testing"
)

const webhookPayload = `{
  "action": "opened",
  "number": 1,
  "pull_request": {
    "url": "https://api.github.com/repos/Spazzy757/paul/pulls/1",
    "id": 1111111111,
    "html_url": "https://github.com/Spazzy757/paul/pull/1",
    "diff_url": "https://github.com/Spazzy757/paul/pull/1.diff",
    "patch_url": "https://github.com/Spazzy757/paul/pull/1.patch",
    "issue_url": "https://api.github.com/repos/Spazzy757/paul/issues/1",
    "number": 1,
    "state": "open",
    "locked": false,
    "title": "Added basic webserver",
    "user": {
      "login": "Spazzy757",
      "id": 111111111,
      "avatar_url": "https://avatars1.githubusercontent.com/u/19777480?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/Spazzy757",
      "html_url": "https://github.com/Spazzy757",
      "followers_url": "https://api.github.com/users/Spazzy757/followers",
      "following_url": "https://api.github.com/users/Spazzy757/following{/other_user}",
      "gists_url": "https://api.github.com/users/Spazzy757/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/Spazzy757/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/Spazzy757/subscriptions",
      "organizations_url": "https://api.github.com/users/Spazzy757/orgs",
      "repos_url": "https://api.github.com/users/Spazzy757/repos",
      "events_url": "https://api.github.com/users/Spazzy757/events{/privacy}",
      "received_events_url": "https://api.github.com/users/Spazzy757/received_events",
      "type": "User",
      "site_admin": false
    },
    "body": "# Changelog\r\n\r\nAdded a basic webserver that will handle Github requests",
    "created_at": "2020-10-06T09:46:47Z",
    "updated_at": "2020-10-06T09:46:47Z",
    "closed_at": null,
    "merged_at": null,
    "merge_commit_sha": null,
    "assignee": null,
    "assignees": [

    ],
    "requested_reviewers": [

    ],
    "requested_teams": [

    ],
    "labels": [

    ],
    "milestone": null,
    "draft": false,
    "commits_url": "https://api.github.com/repos/Spazzy757/paul/pulls/1/commits",
    "review_comments_url": "https://api.github.com/repos/Spazzy757/paul/pulls/1/comments",
    "review_comment_url": "https://api.github.com/repos/Spazzy757/paul/pulls/comments{/number}",
    "comments_url": "https://api.github.com/repos/Spazzy757/paul/issues/1/comments",
    "statuses_url": "https://api.github.com/repos/Spazzy757/paul/statuses/83e12d84247dcc85e05ea18d558be01ce6b0c128",
    "head": {
      "label": "Spazzy757:feature-added-webserver",
      "ref": "feature-added-webserver",
      "sha": "83e12d84247dcc85e05ea18d558be01ce6b0c128",
      "user": {
        "login": "Spazzy757",
        "id": 1111111111,
        "avatar_url": "https://avatars1.githubusercontent.com/u/19777480?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/Spazzy757",
        "html_url": "https://github.com/Spazzy757",
        "followers_url": "https://api.github.com/users/Spazzy757/followers",
        "following_url": "https://api.github.com/users/Spazzy757/following{/other_user}",
        "gists_url": "https://api.github.com/users/Spazzy757/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/Spazzy757/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/Spazzy757/subscriptions",
        "organizations_url": "https://api.github.com/users/Spazzy757/orgs",
        "repos_url": "https://api.github.com/users/Spazzy757/repos",
        "events_url": "https://api.github.com/users/Spazzy757/events{/privacy}",
        "received_events_url": "https://api.github.com/users/Spazzy757/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 111112111,
        "name": "paul",
        "full_name": "Spazzy757/paul",
        "private": false,
        "owner": {
          "login": "Spazzy757",
          "id": 111111111,
          "avatar_url": "https://avatars1.githubusercontent.com/u/19777480?v=4",
          "gravatar_id": "",
          "url": "https://api.github.com/users/Spazzy757",
          "html_url": "https://github.com/Spazzy757",
          "followers_url": "https://api.github.com/users/Spazzy757/followers",
          "following_url": "https://api.github.com/users/Spazzy757/following{/other_user}",
          "gists_url": "https://api.github.com/users/Spazzy757/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/Spazzy757/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/Spazzy757/subscriptions",
          "organizations_url": "https://api.github.com/users/Spazzy757/orgs",
          "repos_url": "https://api.github.com/users/Spazzy757/repos",
          "events_url": "https://api.github.com/users/Spazzy757/events{/privacy}",
          "received_events_url": "https://api.github.com/users/Spazzy757/received_events",
          "type": "User",
          "site_admin": false
        },
        "html_url": "https://github.com/Spazzy757/paul",
        "description": "A Github Bot that will help with reviewing API tests",
        "fork": false,
        "url": "https://api.github.com/repos/Spazzy757/paul",
        "forks_url": "https://api.github.com/repos/Spazzy757/paul/forks",
        "keys_url": "https://api.github.com/repos/Spazzy757/paul/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/Spazzy757/paul/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/Spazzy757/paul/teams",
        "hooks_url": "https://api.github.com/repos/Spazzy757/paul/hooks",
        "issue_events_url": "https://api.github.com/repos/Spazzy757/paul/issues/events{/number}",
        "events_url": "https://api.github.com/repos/Spazzy757/paul/events",
        "assignees_url": "https://api.github.com/repos/Spazzy757/paul/assignees{/user}",
        "branches_url": "https://api.github.com/repos/Spazzy757/paul/branches{/branch}",
        "tags_url": "https://api.github.com/repos/Spazzy757/paul/tags",
        "blobs_url": "https://api.github.com/repos/Spazzy757/paul/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/Spazzy757/paul/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/Spazzy757/paul/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/Spazzy757/paul/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/Spazzy757/paul/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/Spazzy757/paul/languages",
        "stargazers_url": "https://api.github.com/repos/Spazzy757/paul/stargazers",
        "contributors_url": "https://api.github.com/repos/Spazzy757/paul/contributors",
        "subscribers_url": "https://api.github.com/repos/Spazzy757/paul/subscribers",
        "subscription_url": "https://api.github.com/repos/Spazzy757/paul/subscription",
        "commits_url": "https://api.github.com/repos/Spazzy757/paul/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/Spazzy757/paul/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/Spazzy757/paul/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/Spazzy757/paul/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/Spazzy757/paul/contents/{+path}",
        "compare_url": "https://api.github.com/repos/Spazzy757/paul/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/Spazzy757/paul/merges",
        "archive_url": "https://api.github.com/repos/Spazzy757/paul/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/Spazzy757/paul/downloads",
        "issues_url": "https://api.github.com/repos/Spazzy757/paul/issues{/number}",
        "pulls_url": "https://api.github.com/repos/Spazzy757/paul/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/Spazzy757/paul/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/Spazzy757/paul/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/Spazzy757/paul/labels{/name}",
        "releases_url": "https://api.github.com/repos/Spazzy757/paul/releases{/id}",
        "deployments_url": "https://api.github.com/repos/Spazzy757/paul/deployments",
        "created_at": "2020-10-06T08:48:53Z",
        "updated_at": "2020-10-06T08:48:58Z",
        "pushed_at": "2020-10-06T09:46:19Z",
        "git_url": "git://github.com/Spazzy757/paul.git",
        "ssh_url": "git@github.com:Spazzy757/paul.git",
        "clone_url": "https://github.com/Spazzy757/paul.git",
        "svn_url": "https://github.com/Spazzy757/paul",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_projects": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": false,
        "forks_count": 0,
        "mirror_url": null,
        "archived": false,
        "disabled": false,
        "open_issues_count": 1,
        "license": {
          "key": "apache-2.0",
          "name": "Apache License 2.0",
          "spdx_id": "Apache-2.0",
          "url": "https://api.github.com/licenses/apache-2.0",
          "node_id": "MDc6TGljZW5zZTI="
        },
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "main",
        "allow_squash_merge": true,
        "allow_merge_commit": true,
        "allow_rebase_merge": true,
        "delete_branch_on_merge": false
      }
    },
    "base": {
      "label": "Spazzy757:main",
      "ref": "main",
      "user": {
        "login": "Spazzy757",
        "id": 11111111,
        "avatar_url": "https://avatars1.githubusercontent.com/u/19777480?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/Spazzy757",
        "html_url": "https://github.com/Spazzy757",
        "followers_url": "https://api.github.com/users/Spazzy757/followers",
        "following_url": "https://api.github.com/users/Spazzy757/following{/other_user}",
        "gists_url": "https://api.github.com/users/Spazzy757/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/Spazzy757/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/Spazzy757/subscriptions",
        "organizations_url": "https://api.github.com/users/Spazzy757/orgs",
        "repos_url": "https://api.github.com/users/Spazzy757/repos",
        "events_url": "https://api.github.com/users/Spazzy757/events{/privacy}",
        "received_events_url": "https://api.github.com/users/Spazzy757/received_events",
        "type": "User",
        "site_admin": false
      },
      "repo": {
        "id": 11111111,
        "name": "paul",
        "full_name": "Spazzy757/paul",
        "private": false,
        "owner": {
          "login": "Spazzy757",
          "id": 19777480,
          "avatar_url": "https://avatars1.githubusercontent.com/u/19777480?v=4",
          "gravatar_id": "",
          "url": "https://api.github.com/users/Spazzy757",
          "html_url": "https://github.com/Spazzy757",
          "followers_url": "https://api.github.com/users/Spazzy757/followers",
          "following_url": "https://api.github.com/users/Spazzy757/following{/other_user}",
          "gists_url": "https://api.github.com/users/Spazzy757/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/Spazzy757/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/Spazzy757/subscriptions",
          "organizations_url": "https://api.github.com/users/Spazzy757/orgs",
          "repos_url": "https://api.github.com/users/Spazzy757/repos",
          "events_url": "https://api.github.com/users/Spazzy757/events{/privacy}",
          "received_events_url": "https://api.github.com/users/Spazzy757/received_events",
          "type": "User",
          "site_admin": false
        },
        "html_url": "https://github.com/Spazzy757/paul",
        "description": "A Github Bot that will help with reviewing API tests",
        "fork": false,
        "url": "https://api.github.com/repos/Spazzy757/paul",
        "forks_url": "https://api.github.com/repos/Spazzy757/paul/forks",
        "keys_url": "https://api.github.com/repos/Spazzy757/paul/keys{/key_id}",
        "collaborators_url": "https://api.github.com/repos/Spazzy757/paul/collaborators{/collaborator}",
        "teams_url": "https://api.github.com/repos/Spazzy757/paul/teams",
        "hooks_url": "https://api.github.com/repos/Spazzy757/paul/hooks",
        "issue_events_url": "https://api.github.com/repos/Spazzy757/paul/issues/events{/number}",
        "events_url": "https://api.github.com/repos/Spazzy757/paul/events",
        "assignees_url": "https://api.github.com/repos/Spazzy757/paul/assignees{/user}",
        "branches_url": "https://api.github.com/repos/Spazzy757/paul/branches{/branch}",
        "tags_url": "https://api.github.com/repos/Spazzy757/paul/tags",
        "blobs_url": "https://api.github.com/repos/Spazzy757/paul/git/blobs{/sha}",
        "git_tags_url": "https://api.github.com/repos/Spazzy757/paul/git/tags{/sha}",
        "git_refs_url": "https://api.github.com/repos/Spazzy757/paul/git/refs{/sha}",
        "trees_url": "https://api.github.com/repos/Spazzy757/paul/git/trees{/sha}",
        "statuses_url": "https://api.github.com/repos/Spazzy757/paul/statuses/{sha}",
        "languages_url": "https://api.github.com/repos/Spazzy757/paul/languages",
        "stargazers_url": "https://api.github.com/repos/Spazzy757/paul/stargazers",
        "contributors_url": "https://api.github.com/repos/Spazzy757/paul/contributors",
        "subscribers_url": "https://api.github.com/repos/Spazzy757/paul/subscribers",
        "subscription_url": "https://api.github.com/repos/Spazzy757/paul/subscription",
        "commits_url": "https://api.github.com/repos/Spazzy757/paul/commits{/sha}",
        "git_commits_url": "https://api.github.com/repos/Spazzy757/paul/git/commits{/sha}",
        "comments_url": "https://api.github.com/repos/Spazzy757/paul/comments{/number}",
        "issue_comment_url": "https://api.github.com/repos/Spazzy757/paul/issues/comments{/number}",
        "contents_url": "https://api.github.com/repos/Spazzy757/paul/contents/{+path}",
        "compare_url": "https://api.github.com/repos/Spazzy757/paul/compare/{base}...{head}",
        "merges_url": "https://api.github.com/repos/Spazzy757/paul/merges",
        "archive_url": "https://api.github.com/repos/Spazzy757/paul/{archive_format}{/ref}",
        "downloads_url": "https://api.github.com/repos/Spazzy757/paul/downloads",
        "issues_url": "https://api.github.com/repos/Spazzy757/paul/issues{/number}",
        "pulls_url": "https://api.github.com/repos/Spazzy757/paul/pulls{/number}",
        "milestones_url": "https://api.github.com/repos/Spazzy757/paul/milestones{/number}",
        "notifications_url": "https://api.github.com/repos/Spazzy757/paul/notifications{?since,all,participating}",
        "labels_url": "https://api.github.com/repos/Spazzy757/paul/labels{/name}",
        "releases_url": "https://api.github.com/repos/Spazzy757/paul/releases{/id}",
        "deployments_url": "https://api.github.com/repos/Spazzy757/paul/deployments",
        "created_at": "2020-10-06T08:48:53Z",
        "updated_at": "2020-10-06T08:48:58Z",
        "pushed_at": "2020-10-06T09:46:19Z",
        "git_url": "git://github.com/Spazzy757/paul.git",
        "ssh_url": "git@github.com:Spazzy757/paul.git",
        "clone_url": "https://github.com/Spazzy757/paul.git",
        "svn_url": "https://github.com/Spazzy757/paul",
        "homepage": null,
        "size": 0,
        "stargazers_count": 0,
        "watchers_count": 0,
        "language": null,
        "has_issues": true,
        "has_projects": true,
        "has_downloads": true,
        "has_wiki": true,
        "has_pages": false,
        "forks_count": 0,
        "mirror_url": null,
        "archived": false,
        "disabled": false,
        "open_issues_count": 1,
        "license": {
          "key": "apache-2.0",
          "name": "Apache License 2.0",
          "spdx_id": "Apache-2.0",
          "url": "https://api.github.com/licenses/apache-2.0",
          "node_id": "MDc6TGljZW5zZTI="
        },
        "forks": 0,
        "open_issues": 1,
        "watchers": 0,
        "default_branch": "main",
        "allow_squash_merge": true,
        "allow_merge_commit": true,
        "allow_rebase_merge": true,
        "delete_branch_on_merge": false
      }
    },
    "_links": {
      "self": {
        "href": "https://api.github.com/repos/Spazzy757/paul/pulls/1"
      },
      "html": {
        "href": "https://github.com/Spazzy757/paul/pull/1"
      },
      "issue": {
        "href": "https://api.github.com/repos/Spazzy757/paul/issues/1"
      },
      "comments": {
        "href": "https://api.github.com/repos/Spazzy757/paul/issues/1/comments"
      },
      "review_comments": {
        "href": "https://api.github.com/repos/Spazzy757/paul/pulls/1/comments"
      },
      "review_comment": {
        "href": "https://api.github.com/repos/Spazzy757/paul/pulls/comments{/number}"
      },
      "commits": {
        "href": "https://api.github.com/repos/Spazzy757/paul/pulls/1/commits"
      },
      "statuses": {
        "href": "https://api.github.com/repos/Spazzy757/paul/statuses/83e12d84247dcc85e05ea18d558be01ce6b0c128"
      }
    },
    "author_association": "OWNER",
    "active_lock_reason": null,
    "merged": false,
    "mergeable": null,
    "rebaseable": null,
    "mergeable_state": "unknown",
    "merged_by": null,
    "comments": 0,
    "review_comments": 0,
    "maintainer_can_modify": false,
    "commits": 1,
    "additions": 137,
    "deletions": 0,
    "changed_files": 8
  },
  "repository": {
    "id": 301666609,
    "node_id": "dGVzdA==",
    "name": "paul",
    "full_name": "Spazzy757/paul",
    "private": false,
    "owner": {
      "login": "Spazzy757",
      "id": 19777480,
      "node_id": "dGVzdA==",
      "avatar_url": "https://avatars1.githubusercontent.com/u/19777480?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/Spazzy757",
      "html_url": "https://github.com/Spazzy757",
      "followers_url": "https://api.github.com/users/Spazzy757/followers",
      "following_url": "https://api.github.com/users/Spazzy757/following{/other_user}",
      "gists_url": "https://api.github.com/users/Spazzy757/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/Spazzy757/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/Spazzy757/subscriptions",
      "organizations_url": "https://api.github.com/users/Spazzy757/orgs",
      "repos_url": "https://api.github.com/users/Spazzy757/repos",
      "events_url": "https://api.github.com/users/Spazzy757/events{/privacy}",
      "received_events_url": "https://api.github.com/users/Spazzy757/received_events",
      "type": "User",
      "site_admin": false
    },
    "html_url": "https://github.com/Spazzy757/paul",
    "description": "A Github Bot that will help with reviewing API tests",
    "fork": false,
    "url": "https://api.github.com/repos/Spazzy757/paul",
    "forks_url": "https://api.github.com/repos/Spazzy757/paul/forks",
    "keys_url": "https://api.github.com/repos/Spazzy757/paul/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/Spazzy757/paul/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/Spazzy757/paul/teams",
    "hooks_url": "https://api.github.com/repos/Spazzy757/paul/hooks",
    "issue_events_url": "https://api.github.com/repos/Spazzy757/paul/issues/events{/number}",
    "events_url": "https://api.github.com/repos/Spazzy757/paul/events",
    "assignees_url": "https://api.github.com/repos/Spazzy757/paul/assignees{/user}",
    "branches_url": "https://api.github.com/repos/Spazzy757/paul/branches{/branch}",
    "tags_url": "https://api.github.com/repos/Spazzy757/paul/tags",
    "blobs_url": "https://api.github.com/repos/Spazzy757/paul/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/Spazzy757/paul/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/Spazzy757/paul/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/Spazzy757/paul/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/Spazzy757/paul/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/Spazzy757/paul/languages",
    "stargazers_url": "https://api.github.com/repos/Spazzy757/paul/stargazers",
    "contributors_url": "https://api.github.com/repos/Spazzy757/paul/contributors",
    "subscribers_url": "https://api.github.com/repos/Spazzy757/paul/subscribers",
    "subscription_url": "https://api.github.com/repos/Spazzy757/paul/subscription",
    "commits_url": "https://api.github.com/repos/Spazzy757/paul/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/Spazzy757/paul/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/Spazzy757/paul/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/Spazzy757/paul/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/Spazzy757/paul/contents/{+path}",
    "compare_url": "https://api.github.com/repos/Spazzy757/paul/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/Spazzy757/paul/merges",
    "archive_url": "https://api.github.com/repos/Spazzy757/paul/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/Spazzy757/paul/downloads",
    "issues_url": "https://api.github.com/repos/Spazzy757/paul/issues{/number}",
    "pulls_url": "https://api.github.com/repos/Spazzy757/paul/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/Spazzy757/paul/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/Spazzy757/paul/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/Spazzy757/paul/labels{/name}",
    "releases_url": "https://api.github.com/repos/Spazzy757/paul/releases{/id}",
    "deployments_url": "https://api.github.com/repos/Spazzy757/paul/deployments",
    "created_at": "2020-10-06T08:48:53Z",
    "updated_at": "2020-10-06T08:48:58Z",
    "pushed_at": "2020-10-06T09:46:19Z",
    "git_url": "git://github.com/Spazzy757/paul.git",
    "ssh_url": "git@github.com:Spazzy757/paul.git",
    "clone_url": "https://github.com/Spazzy757/paul.git",
    "svn_url": "https://github.com/Spazzy757/paul",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 1,
    "license": {
      "key": "apache-2.0",
      "name": "Apache License 2.0",
      "spdx_id": "Apache-2.0",
      "url": "https://api.github.com/licenses/apache-2.0",
      "node_id": "MDc6TGljZW5zZTI="
    },
    "forks": 0,
    "open_issues": 1,
    "watchers": 0,
    "default_branch": "main"
  },
  "sender": {
    "login": "Spazzy757",
    "id": 19777480,
    "node_id": "MDQ6VXNlcjE5Nzc3NDgw",
    "avatar_url": "https://avatars1.githubusercontent.com/u/19777480?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/Spazzy757",
    "html_url": "https://github.com/Spazzy757",
    "followers_url": "https://api.github.com/users/Spazzy757/followers",
    "following_url": "https://api.github.com/users/Spazzy757/following{/other_user}",
    "gists_url": "https://api.github.com/users/Spazzy757/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/Spazzy757/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/Spazzy757/subscriptions",
    "organizations_url": "https://api.github.com/users/Spazzy757/orgs",
    "repos_url": "https://api.github.com/users/Spazzy757/repos",
    "events_url": "https://api.github.com/users/Spazzy757/events{/privacy}",
    "received_events_url": "https://api.github.com/users/Spazzy757/received_events",
    "type": "User",
    "site_admin": false
  }
}`

type mockClient struct {
	resp *github.PullRequestReview
}

func (m *mockClient) CreateReview(ctx context.Context, owner string, repo string, number int, review *github.PullRequestReviewRequest) (*github.PullRequestReview, *github.Response, error) {
	return m.resp, nil, nil
}

func TestCreateReview(t *testing.T) {
	t.Run("Test Webhook is Handled correctly", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(webhookPayload)))
		req.Header.Set("X-GitHub-Event", "pull_request")
		ctx := context.Background()
		mc := &mockClient{
			resp: &github.PullRequestReview{
				ID: github.Int64(1),
			},
		}
		pr := &pullRequestClient{ctx: ctx, client: mc}
		event, _ := github.ParseWebHook(github.WebHookType(req), []byte(webhookPayload))
		switch e := event.(type) {
		case *github.PullRequestEvent:
			if err := comment(e.PullRequest, pr, "test"); err != nil {
				t.Fatalf("createReview: %v", err)
			}
		default:
			t.Fatalf("Event Type Not Pull Request")
		}
	})
}

func TestGetClient(t *testing.T) {
	t.Run("Validate Get Client Returns Type *github.Client", func(t *testing.T) {
		client := getClient()
		log.Println(reflect.TypeOf(client))
		expected := github.NewClient(nil)
		if reflect.TypeOf(client) != reflect.TypeOf(expected) {
			t.Fatalf("Github Client Not Returned")
		}
	})
}