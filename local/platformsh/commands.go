// Code generated by platformsh/generator/main.go
// DO NOT EDIT

/*
 * Copyright (c) 2021-present Fabien Potencier <fabien@symfony.com>
 *
 * This file is part of Symfony CLI project
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package platformsh

import (
	"github.com/symfony-cli/console"
)

var Commands = []*console.Command{
	{
		Category: "cloud",
		Name:     "_completion",
		Usage:    "BASH completion hook.",		Hidden:   console.Hide,
	},
	{
		Category: "cloud",
		Name:     "bot",
		Usage:    "The Platform.sh Bot",		Hidden:   console.Hide,
	},
	{
		Category: "cloud",
		Name:     "clear-cache",
		Aliases:  []*console.Alias{
			{Name: "cloud:clearcache"},
			{Name: "clearcache", Hidden: true},
			{Name: "cloud:cc"},
			{Name: "cc", Hidden: true},
		},
		Usage:    "Clear the CLI cache",
	},
	{
		Category: "cloud",
		Name:     "decode",
		Usage:    "Decode an encoded string such as PLATFORM_VARIABLES",
	},
	{
		Category: "cloud",
		Name:     "docs",
		Usage:    "Open the online documentation",
	},
	{
		Category: "cloud",
		Name:     "legacy-migrate",
		Usage:    "Migrate from the legacy file structure",		Hidden:   console.Hide,
	},
	{
		Category: "cloud",
		Name:     "multi",
		Usage:    "Execute a command on multiple projects",
	},
	{
		Category: "cloud",
		Name:     "web",
		Usage:    "Open the Web UI",
	},
	{
		Category: "cloud",
		Name:     "welcome",
		Usage:    "Welcome to Platform.sh",		Hidden:   console.Hide,
	},
	{
		Category: "cloud",
		Name:     "winky",
		Usage:    "",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:activity",
		Name:     "cancel",
		Aliases:  []*console.Alias{
			{Name: "activity:cancel", Hidden: true},
		},
		Usage:    "Cancel an activity",
	},
	{
		Category: "cloud:activity",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "activity:get", Hidden: true},
		},
		Usage:    "View detailed information on a single activity",
	},
	{
		Category: "cloud:activity",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "activity:list", Hidden: true},
			{Name: "cloud:activities"},
			{Name: "activities", Hidden: true},
			{Name: "cloud:act"},
			{Name: "act", Hidden: true},
		},
		Usage:    "Get a list of activities for an environment or project",
	},
	{
		Category: "cloud:activity",
		Name:     "log",
		Aliases:  []*console.Alias{
			{Name: "activity:log", Hidden: true},
		},
		Usage:    "Display the log for an activity",
	},
	{
		Category: "cloud:api",
		Name:     "curl",
		Aliases:  []*console.Alias{
			{Name: "api:curl", Hidden: true},
		},
		Usage:    "Run an authenticated cURL request on the Platform.sh API",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:app",
		Name:     "config-get",
		Aliases:  []*console.Alias{
			{Name: "app:config-get", Hidden: true},
		},
		Usage:    "View the configuration of an app",
	},
	{
		Category: "cloud:app",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "app:list", Hidden: true},
			{Name: "cloud:apps"},
			{Name: "apps", Hidden: true},
		},
		Usage:    "List apps in the project",
	},
	{
		Category: "cloud:auth",
		Name:     "api-token-login",
		Aliases:  []*console.Alias{
			{Name: "auth:api-token-login", Hidden: true},
		},
		Usage:    "Log in to Platform.sh using an API token",
	},
	{
		Category: "cloud:auth",
		Name:     "browser-login",
		Aliases:  []*console.Alias{
			{Name: "auth:browser-login", Hidden: true},
			{Name: "cloud:login"},
			{Name: "login", Hidden: true},
		},
		Usage:    "Log in to Platform.sh via a browser",
	},
	{
		Category: "cloud:auth",
		Name:     "info",
		Aliases:  []*console.Alias{
			{Name: "auth:info", Hidden: true},
		},
		Usage:    "Display your account information",
	},
	{
		Category: "cloud:auth",
		Name:     "logout",
		Aliases:  []*console.Alias{
			{Name: "auth:logout", Hidden: true},
			{Name: "cloud:logout"},
			{Name: "logout", Hidden: true},
		},
		Usage:    "Log out of Platform.sh",
	},
	{
		Category: "cloud:auth",
		Name:     "token",
		Aliases:  []*console.Alias{
			{Name: "auth:token", Hidden: true},
		},
		Usage:    "Obtain an OAuth 2 access token for requests to Platform.sh APIs",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:backup",
		Name:     "create",
		Aliases:  []*console.Alias{
			{Name: "backup:create", Hidden: true},
			{Name: "cloud:backup"},
			{Name: "backup", Hidden: true},
			{Name: "cloud:snapshot:create"},
			{Name: "snapshot:create", Hidden: true},
			{Name: "cloud:environment:backup"},
			{Name: "environment:backup", Hidden: true},
		},
		Usage:    "Make a backup of an environment",
	},
	{
		Category: "cloud:backup",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "backup:list", Hidden: true},
			{Name: "cloud:backups"},
			{Name: "backups", Hidden: true},
			{Name: "cloud:snapshots"},
			{Name: "snapshots", Hidden: true},
			{Name: "cloud:snapshot:list"},
			{Name: "snapshot:list", Hidden: true},
		},
		Usage:    "List available backups of an environment",
	},
	{
		Category: "cloud:backup",
		Name:     "restore",
		Aliases:  []*console.Alias{
			{Name: "backup:restore", Hidden: true},
			{Name: "cloud:environment:restore"},
			{Name: "environment:restore", Hidden: true},
			{Name: "cloud:snapshot:restore"},
			{Name: "snapshot:restore", Hidden: true},
		},
		Usage:    "Restore an environment backup",
	},
	{
		Category: "cloud:certificate",
		Name:     "add",
		Aliases:  []*console.Alias{
			{Name: "certificate:add", Hidden: true},
		},
		Usage:    "Add an SSL certificate to the project",
	},
	{
		Category: "cloud:certificate",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "certificate:delete", Hidden: true},
		},
		Usage:    "Delete a certificate from the project",
	},
	{
		Category: "cloud:certificate",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "certificate:get", Hidden: true},
		},
		Usage:    "View a certificate",
	},
	{
		Category: "cloud:certificate",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "certificate:list", Hidden: true},
			{Name: "cloud:certificates"},
			{Name: "certificates", Hidden: true},
			{Name: "cloud:certs"},
			{Name: "certs", Hidden: true},
		},
		Usage:    "List project certificates",
	},
	{
		Category: "cloud:commit",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "commit:get", Hidden: true},
		},
		Usage:    "Show commit details",
	},
	{
		Category: "cloud:commit",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "commit:list", Hidden: true},
			{Name: "cloud:commits"},
			{Name: "commits", Hidden: true},
		},
		Usage:    "List commits",
	},
	{
		Category: "cloud:db",
		Name:     "dump",
		Aliases:  []*console.Alias{
			{Name: "db:dump", Hidden: true},
			{Name: "cloud:sql-dump"},
			{Name: "sql-dump", Hidden: true},
			{Name: "cloud:environment:sql-dump"},
			{Name: "environment:sql-dump", Hidden: true},
		},
		Usage:    "Create a local dump of the remote database",
	},
	{
		Category: "cloud:db",
		Name:     "size",
		Aliases:  []*console.Alias{
			{Name: "db:size", Hidden: true},
		},
		Usage:    "Estimate the disk usage of a database",
	},
	{
		Category: "cloud:db",
		Name:     "sql",
		Aliases:  []*console.Alias{
			{Name: "db:sql", Hidden: true},
			{Name: "cloud:sql"},
			{Name: "sql", Hidden: true},
			{Name: "cloud:environment:sql"},
			{Name: "environment:sql", Hidden: true},
		},
		Usage:    "Run SQL on the remote database",
	},
	{
		Category: "cloud:domain",
		Name:     "add",
		Aliases:  []*console.Alias{
			{Name: "domain:add", Hidden: true},
		},
		Usage:    "Add a new domain to the project",
	},
	{
		Category: "cloud:domain",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "domain:delete", Hidden: true},
		},
		Usage:    "Delete a domain from the project",
	},
	{
		Category: "cloud:domain",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "domain:get", Hidden: true},
		},
		Usage:    "Show detailed information for a domain",
	},
	{
		Category: "cloud:domain",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "domain:list", Hidden: true},
			{Name: "cloud:domains"},
			{Name: "domains", Hidden: true},
		},
		Usage:    "Get a list of all domains",
	},
	{
		Category: "cloud:domain",
		Name:     "update",
		Aliases:  []*console.Alias{
			{Name: "domain:update", Hidden: true},
		},
		Usage:    "Update a domain",
	},
	{
		Category: "cloud:environment",
		Name:     "activate",
		Aliases:  []*console.Alias{
			{Name: "environment:activate", Hidden: true},
		},
		Usage:    "Activate an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "branch",
		Aliases:  []*console.Alias{
			{Name: "environment:branch", Hidden: true},
			{Name: "cloud:branch"},
			{Name: "branch", Hidden: true},
		},
		Usage:    "Branch an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "checkout",
		Aliases:  []*console.Alias{
			{Name: "environment:checkout", Hidden: true},
			{Name: "cloud:checkout"},
			{Name: "checkout", Hidden: true},
		},
		Usage:    "Check out an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "curl",
		Aliases:  []*console.Alias{
			{Name: "environment:curl", Hidden: true},
		},
		Usage:    "Run an authenticated cURL request on an environment's API",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:environment",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "environment:delete", Hidden: true},
			{Name: "cloud:environment:deactivate"},
			{Name: "environment:deactivate", Hidden: true},
		},
		Usage:    "Delete an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "drush",
		Aliases:  []*console.Alias{
			{Name: "environment:drush", Hidden: true},
			{Name: "cloud:drush"},
			{Name: "drush", Hidden: true},
		},
		Usage:    "Run a drush command on the remote environment",
	},
	{
		Category: "cloud:environment",
		Name:     "http-access",
		Aliases:  []*console.Alias{
			{Name: "environment:http-access", Hidden: true},
			{Name: "cloud:httpaccess"},
			{Name: "httpaccess", Hidden: true},
		},
		Usage:    "Update HTTP access settings for an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "info",
		Aliases:  []*console.Alias{
			{Name: "environment:info", Hidden: true},
			{Name: "cloud:environment:metadata"},
			{Name: "environment:metadata", Hidden: true},
		},
		Usage:    "Read or set properties for an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "init",
		Aliases:  []*console.Alias{
			{Name: "environment:init", Hidden: true},
		},
		Usage:    "Initialize an environment from a public Git repository",
	},
	{
		Category: "cloud:environment",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "environment:list", Hidden: true},
			{Name: "cloud:environments"},
			{Name: "environments", Hidden: true},
			{Name: "cloud:env"},
			{Name: "env", Hidden: true},
		},
		Usage:    "Get a list of environments",
	},
	{
		Category: "cloud:environment",
		Name:     "logs",
		Aliases:  []*console.Alias{
			{Name: "environment:logs", Hidden: true},
			{Name: "cloud:log"},
			{Name: "log", Hidden: true},
			{Name: "cloud:logs"},
			{Name: "logs", Hidden: true},
		},
		Usage:    "Read an environment's logs",
	},
	{
		Category: "cloud:environment",
		Name:     "merge",
		Aliases:  []*console.Alias{
			{Name: "environment:merge", Hidden: true},
			{Name: "cloud:merge"},
			{Name: "merge", Hidden: true},
		},
		Usage:    "Merge an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "push",
		Aliases:  []*console.Alias{
			{Name: "environment:push", Hidden: true},
			{Name: "cloud:push"},
			{Name: "push", Hidden: true},
			{Name: "deploy"},
			{Name: "cloud:deploy"},
		},
		Usage:    "Push code to an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "redeploy",
		Aliases:  []*console.Alias{
			{Name: "environment:redeploy", Hidden: true},
			{Name: "cloud:redeploy"},
			{Name: "redeploy", Hidden: true},
		},
		Usage:    "Redeploy an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "relationships",
		Aliases:  []*console.Alias{
			{Name: "environment:relationships", Hidden: true},
			{Name: "cloud:relationships"},
			{Name: "relationships", Hidden: true},
		},
		Usage:    "Show an environment's relationships",
	},
	{
		Category: "cloud:environment",
		Name:     "scp",
		Aliases:  []*console.Alias{
			{Name: "environment:scp", Hidden: true},
			{Name: "cloud:scp"},
			{Name: "scp", Hidden: true},
		},
		Usage:    "Copy files to and from current environment using scp",
	},
	{
		Category: "cloud:environment",
		Name:     "set-remote",
		Aliases:  []*console.Alias{
			{Name: "environment:set-remote", Hidden: true},
		},
		Usage:    "Set the remote environment to map to a branch",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:environment",
		Name:     "ssh",
		Aliases:  []*console.Alias{
			{Name: "environment:ssh", Hidden: true},
			{Name: "cloud:ssh"},
			{Name: "ssh", Hidden: true},
		},
		Usage:    "SSH to the current environment",
	},
	{
		Category: "cloud:environment",
		Name:     "synchronize",
		Aliases:  []*console.Alias{
			{Name: "environment:synchronize", Hidden: true},
			{Name: "cloud:sync"},
			{Name: "sync", Hidden: true},
		},
		Usage:    "Synchronize an environment's code and/or data from its parent",
	},
	{
		Category: "cloud:environment",
		Name:     "url",
		Aliases:  []*console.Alias{
			{Name: "environment:url", Hidden: true},
			{Name: "cloud:url"},
			{Name: "url", Hidden: true},
		},
		Usage:    "Get the public URLs of an environment",
	},
	{
		Category: "cloud:environment",
		Name:     "xdebug",
		Aliases:  []*console.Alias{
			{Name: "environment:xdebug", Hidden: true},
			{Name: "cloud:xdebug"},
			{Name: "xdebug", Hidden: true},
		},
		Usage:    "Open a tunnel to Xdebug on the environment",
	},
	{
		Category: "cloud:integration",
		Name:     "activity:get",
		Aliases:  []*console.Alias{
			{Name: "integration:activity:get", Hidden: true},
		},
		Usage:    "View detailed information on a single integration activity",
	},
	{
		Category: "cloud:integration",
		Name:     "activity:list",
		Aliases:  []*console.Alias{
			{Name: "integration:activity:list", Hidden: true},
			{Name: "cloud:i:act"},
			{Name: "i:act", Hidden: true},
			{Name: "cloud:integration:activities"},
			{Name: "integration:activities", Hidden: true},
		},
		Usage:    "Get a list of activities for an integration",
	},
	{
		Category: "cloud:integration",
		Name:     "activity:log",
		Aliases:  []*console.Alias{
			{Name: "integration:activity:log", Hidden: true},
		},
		Usage:    "Display the log for an integration activity",
	},
	{
		Category: "cloud:integration",
		Name:     "add",
		Aliases:  []*console.Alias{
			{Name: "integration:add", Hidden: true},
		},
		Usage:    "Add an integration to the project",
	},
	{
		Category: "cloud:integration",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "integration:delete", Hidden: true},
		},
		Usage:    "Delete an integration from a project",
	},
	{
		Category: "cloud:integration",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "integration:get", Hidden: true},
		},
		Usage:    "View details of an integration",
	},
	{
		Category: "cloud:integration",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "integration:list", Hidden: true},
			{Name: "cloud:integrations"},
			{Name: "integrations", Hidden: true},
		},
		Usage:    "View a list of project integration(s)",
	},
	{
		Category: "cloud:integration",
		Name:     "update",
		Aliases:  []*console.Alias{
			{Name: "integration:update", Hidden: true},
		},
		Usage:    "Update an integration",
	},
	{
		Category: "cloud:integration",
		Name:     "validate",
		Aliases:  []*console.Alias{
			{Name: "integration:validate", Hidden: true},
		},
		Usage:    "Validate an existing integration",
	},
	{
		Category: "cloud:local",
		Name:     "build",
		Aliases:  []*console.Alias{
			{Name: "local:build", Hidden: true},
			{Name: "cloud:build"},
			{Name: "build", Hidden: true},
		},
		Usage:    "Build the current project locally",
	},
	{
		Category: "cloud:local",
		Name:     "clean",
		Aliases:  []*console.Alias{
			{Name: "local:clean", Hidden: true},
			{Name: "cloud:clean"},
			{Name: "clean", Hidden: true},
		},
		Usage:    "Remove old project builds",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:local",
		Name:     "dir",
		Aliases:  []*console.Alias{
			{Name: "local:dir", Hidden: true},
			{Name: "cloud:dir"},
			{Name: "dir", Hidden: true},
		},
		Usage:    "Find the local project root",
	},
	{
		Category: "cloud:local",
		Name:     "drush-aliases",
		Aliases:  []*console.Alias{
			{Name: "local:drush-aliases", Hidden: true},
			{Name: "cloud:drush-aliases"},
			{Name: "drush-aliases", Hidden: true},
		},
		Usage:    "Find the project's Drush aliases",
	},
	{
		Category: "cloud:mount",
		Name:     "download",
		Aliases:  []*console.Alias{
			{Name: "mount:download", Hidden: true},
		},
		Usage:    "Download files from a mount, using rsync",
	},
	{
		Category: "cloud:mount",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "mount:list", Hidden: true},
			{Name: "cloud:mounts"},
			{Name: "mounts", Hidden: true},
		},
		Usage:    "Get a list of mounts",
	},
	{
		Category: "cloud:mount",
		Name:     "size",
		Aliases:  []*console.Alias{
			{Name: "mount:size", Hidden: true},
		},
		Usage:    "Check the disk usage of mounts",
	},
	{
		Category: "cloud:mount",
		Name:     "upload",
		Aliases:  []*console.Alias{
			{Name: "mount:upload", Hidden: true},
		},
		Usage:    "Upload files to a mount, using rsync",
	},
	{
		Category: "cloud:organization",
		Name:     "billing:address",
		Aliases:  []*console.Alias{
			{Name: "organization:billing:address", Hidden: true},
		},
		Usage:    "View or change an organization's billing address",
	},
	{
		Category: "cloud:organization",
		Name:     "billing:profile",
		Aliases:  []*console.Alias{
			{Name: "organization:billing:profile", Hidden: true},
		},
		Usage:    "View or change an organization's billing profile",
	},
	{
		Category: "cloud:organization",
		Name:     "curl",
		Aliases:  []*console.Alias{
			{Name: "organization:curl", Hidden: true},
		},
		Usage:    "Run an authenticated cURL request on an organization's API",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:organization",
		Name:     "info",
		Aliases:  []*console.Alias{
			{Name: "organization:info", Hidden: true},
		},
		Usage:    "View or change organization details",
	},
	{
		Category: "cloud:organization",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "organization:list", Hidden: true},
			{Name: "cloud:orgs"},
			{Name: "orgs", Hidden: true},
			{Name: "cloud:organizations"},
			{Name: "organizations", Hidden: true},
		},
		Usage:    "List organizations",
	},
	{
		Category: "cloud:organization",
		Name:     "subscription:list",
		Aliases:  []*console.Alias{
			{Name: "organization:subscription:list", Hidden: true},
			{Name: "cloud:organization:subscriptions"},
			{Name: "organization:subscriptions", Hidden: true},
		},
		Usage:    "List subscriptions within an organization",
	},
	{
		Category: "cloud:organization",
		Name:     "user:add",
		Aliases:  []*console.Alias{
			{Name: "organization:user:add", Hidden: true},
		},
		Usage:    "Invite a user to an organization",
	},
	{
		Category: "cloud:organization",
		Name:     "user:delete",
		Aliases:  []*console.Alias{
			{Name: "organization:user:delete", Hidden: true},
		},
		Usage:    "Remove a user from an organization",
	},
	{
		Category: "cloud:organization",
		Name:     "user:get",
		Aliases:  []*console.Alias{
			{Name: "organization:user:get", Hidden: true},
		},
		Usage:    "View an organization user",
	},
	{
		Category: "cloud:organization",
		Name:     "user:list",
		Aliases:  []*console.Alias{
			{Name: "organization:user:list", Hidden: true},
			{Name: "cloud:organization:users"},
			{Name: "organization:users", Hidden: true},
		},
		Usage:    "List organization users",
	},
	{
		Category: "cloud:organization",
		Name:     "user:update",
		Aliases:  []*console.Alias{
			{Name: "organization:user:update", Hidden: true},
		},
		Usage:    "Update an organization user",
	},
	{
		Category: "cloud:project",
		Name:     "clear-build-cache",
		Aliases:  []*console.Alias{
			{Name: "project:clear-build-cache", Hidden: true},
		},
		Usage:    "Clear a project's build cache",
	},
	{
		Category: "cloud:project",
		Name:     "create",
		Aliases:  []*console.Alias{
			{Name: "project:create", Hidden: true},
			{Name: "cloud:create"},
			{Name: "create", Hidden: true},
		},
		Usage:    "Create a new project",
	},
	{
		Category: "cloud:project",
		Name:     "curl",
		Aliases:  []*console.Alias{
			{Name: "project:curl", Hidden: true},
		},
		Usage:    "Run an authenticated cURL request on a project's API",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:project",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "project:delete", Hidden: true},
		},
		Usage:    "Delete a project",
	},
	{
		Category: "cloud:project",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "project:get", Hidden: true},
			{Name: "cloud:get"},
			{Name: "get", Hidden: true},
		},
		Usage:    "Clone a project locally",
	},
	{
		Category: "cloud:project",
		Name:     "info",
		Aliases:  []*console.Alias{
			{Name: "project:info", Hidden: true},
			{Name: "cloud:project:metadata"},
			{Name: "project:metadata", Hidden: true},
		},
		Usage:    "Read or set properties for a project",
	},
	{
		Category: "cloud:project",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "project:list", Hidden: true},
			{Name: "cloud:projects"},
			{Name: "projects", Hidden: true},
			{Name: "cloud:pro"},
			{Name: "pro", Hidden: true},
		},
		Usage:    "Get a list of all active projects",
	},
	{
		Category: "cloud:project",
		Name:     "set-remote",
		Aliases:  []*console.Alias{
			{Name: "project:set-remote", Hidden: true},
		},
		Usage:    "Set the remote project for the current Git repository",
	},
	{
		Category: "cloud:repo",
		Name:     "cat",
		Aliases:  []*console.Alias{
			{Name: "repo:cat", Hidden: true},
		},
		Usage:    "Read a file in the project repository",
	},
	{
		Category: "cloud:repo",
		Name:     "ls",
		Aliases:  []*console.Alias{
			{Name: "repo:ls", Hidden: true},
		},
		Usage:    "List files in the project repository",
	},
	{
		Category: "cloud:route",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "route:get", Hidden: true},
		},
		Usage:    "View detailed information about a route",
	},
	{
		Category: "cloud:route",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "route:list", Hidden: true},
			{Name: "cloud:routes"},
			{Name: "routes", Hidden: true},
			{Name: "cloud:environment:routes"},
			{Name: "environment:routes", Hidden: true},
		},
		Usage:    "List all routes for an environment",
	},
	{
		Category: "cloud:self",
		Name:     "install",
		Aliases:  []*console.Alias{
			{Name: "self:install", Hidden: true},
			{Name: "cloud:local:install"},
			{Name: "local:install", Hidden: true},
		},
		Usage:    "Install or update CLI configuration files",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:self",
		Name:     "stats",
		Aliases:  []*console.Alias{
			{Name: "self:stats", Hidden: true},
		},
		Usage:    "View stats on GitHub package downloads",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:self",
		Name:     "update",
		Aliases:  []*console.Alias{
			{Name: "self:update", Hidden: true},
			{Name: "cloud:self-update"},
			{Name: "self-update", Hidden: true},
			{Name: "cloud:update"},
			{Name: "update", Hidden: true},
		},
		Usage:    "Update the CLI to the latest version",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:service",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "service:list", Hidden: true},
			{Name: "cloud:services"},
			{Name: "services", Hidden: true},
		},
		Usage:    "List services in the project",
	},
	{
		Category: "cloud:service",
		Name:     "mongo:dump",
		Aliases:  []*console.Alias{
			{Name: "service:mongo:dump", Hidden: true},
			{Name: "cloud:mongodump"},
			{Name: "mongodump", Hidden: true},
		},
		Usage:    "Create a binary archive dump of data from MongoDB",
	},
	{
		Category: "cloud:service",
		Name:     "mongo:export",
		Aliases:  []*console.Alias{
			{Name: "service:mongo:export", Hidden: true},
			{Name: "cloud:mongoexport"},
			{Name: "mongoexport", Hidden: true},
		},
		Usage:    "Export data from MongoDB",
	},
	{
		Category: "cloud:service",
		Name:     "mongo:restore",
		Aliases:  []*console.Alias{
			{Name: "service:mongo:restore", Hidden: true},
			{Name: "cloud:mongorestore"},
			{Name: "mongorestore", Hidden: true},
		},
		Usage:    "Restore a binary archive dump of data into MongoDB",
	},
	{
		Category: "cloud:service",
		Name:     "mongo:shell",
		Aliases:  []*console.Alias{
			{Name: "service:mongo:shell", Hidden: true},
			{Name: "cloud:mongo"},
			{Name: "mongo", Hidden: true},
		},
		Usage:    "Use the MongoDB shell",
	},
	{
		Category: "cloud:service",
		Name:     "redis-cli",
		Aliases:  []*console.Alias{
			{Name: "service:redis-cli", Hidden: true},
			{Name: "cloud:redis"},
			{Name: "redis", Hidden: true},
		},
		Usage:    "Access the Redis CLI",
	},
	{
		Category: "cloud:session",
		Name:     "switch",
		Aliases:  []*console.Alias{
			{Name: "session:switch", Hidden: true},
		},
		Usage:    "<fg=white;bg=red>[ BETA ]</> Switch between sessions",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:source-operation",
		Name:     "run",
		Aliases:  []*console.Alias{
			{Name: "source-operation:run", Hidden: true},
		},
		Usage:    "<fg=white;bg=red>[ BETA ]</> Run a source operation",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:ssh-cert",
		Name:     "info",
		Aliases:  []*console.Alias{
			{Name: "ssh-cert:info", Hidden: true},
		},
		Usage:    "Display information about the current SSH certificate",		Hidden:   console.Hide,
	},
	{
		Category: "cloud:ssh-cert",
		Name:     "load",
		Aliases:  []*console.Alias{
			{Name: "ssh-cert:load", Hidden: true},
		},
		Usage:    "Generate an SSH certificate",
	},
	{
		Category: "cloud:ssh-key",
		Name:     "add",
		Aliases:  []*console.Alias{
			{Name: "ssh-key:add", Hidden: true},
		},
		Usage:    "Add a new SSH key",
	},
	{
		Category: "cloud:ssh-key",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "ssh-key:delete", Hidden: true},
		},
		Usage:    "Delete an SSH key",
	},
	{
		Category: "cloud:ssh-key",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "ssh-key:list", Hidden: true},
			{Name: "cloud:ssh-keys"},
			{Name: "ssh-keys", Hidden: true},
		},
		Usage:    "Get a list of SSH keys in your account",
	},
	{
		Category: "cloud:subscription",
		Name:     "info",
		Aliases:  []*console.Alias{
			{Name: "subscription:info", Hidden: true},
		},
		Usage:    "Read or modify subscription properties",
	},
	{
		Category: "cloud:tunnel",
		Name:     "close",
		Aliases:  []*console.Alias{
			{Name: "tunnel:close", Hidden: true},
		},
		Usage:    "Close SSH tunnels",
	},
	{
		Category: "cloud:tunnel",
		Name:     "info",
		Aliases:  []*console.Alias{
			{Name: "tunnel:info", Hidden: true},
		},
		Usage:    "View relationship info for SSH tunnels",
	},
	{
		Category: "cloud:tunnel",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "tunnel:list", Hidden: true},
			{Name: "cloud:tunnels"},
			{Name: "tunnels", Hidden: true},
		},
		Usage:    "List SSH tunnels",
	},
	{
		Category: "cloud:tunnel",
		Name:     "open",
		Aliases:  []*console.Alias{
			{Name: "tunnel:open", Hidden: true},
		},
		Usage:    "Open SSH tunnels to an app's relationships",
	},
	{
		Category: "cloud:tunnel",
		Name:     "single",
		Aliases:  []*console.Alias{
			{Name: "tunnel:single", Hidden: true},
		},
		Usage:    "Open a single SSH tunnel to an app relationship",
	},
	{
		Category: "cloud:user",
		Name:     "add",
		Aliases:  []*console.Alias{
			{Name: "user:add", Hidden: true},
		},
		Usage:    "Add a user to the project",
	},
	{
		Category: "cloud:user",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "user:delete", Hidden: true},
		},
		Usage:    "Delete a user from the project",
	},
	{
		Category: "cloud:user",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "user:get", Hidden: true},
			{Name: "cloud:user:role"},
			{Name: "user:role", Hidden: true},
		},
		Usage:    "View a user's role(s)",
	},
	{
		Category: "cloud:user",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "user:list", Hidden: true},
			{Name: "cloud:users"},
			{Name: "users", Hidden: true},
		},
		Usage:    "List project users",
	},
	{
		Category: "cloud:user",
		Name:     "update",
		Aliases:  []*console.Alias{
			{Name: "user:update", Hidden: true},
		},
		Usage:    "Update user role(s) on a project",
	},
	{
		Category: "cloud:variable",
		Name:     "create",
		Aliases:  []*console.Alias{
			{Name: "variable:create", Hidden: true},
		},
		Usage:    "Create a variable",
	},
	{
		Category: "cloud:variable",
		Name:     "delete",
		Aliases:  []*console.Alias{
			{Name: "variable:delete", Hidden: true},
		},
		Usage:    "Delete a variable",
	},
	{
		Category: "cloud:variable",
		Name:     "get",
		Aliases:  []*console.Alias{
			{Name: "variable:get", Hidden: true},
			{Name: "cloud:vget"},
			{Name: "vget", Hidden: true},
		},
		Usage:    "View a variable",
	},
	{
		Category: "cloud:variable",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "variable:list", Hidden: true},
			{Name: "cloud:variables"},
			{Name: "variables", Hidden: true},
			{Name: "cloud:var"},
			{Name: "var", Hidden: true},
		},
		Usage:    "List variables",
	},
	{
		Category: "cloud:variable",
		Name:     "update",
		Aliases:  []*console.Alias{
			{Name: "variable:update", Hidden: true},
		},
		Usage:    "Update a variable",
	},
	{
		Category: "cloud:worker",
		Name:     "list",
		Aliases:  []*console.Alias{
			{Name: "worker:list", Hidden: true},
			{Name: "cloud:workers"},
			{Name: "workers", Hidden: true},
		},
		Usage:    "Get a list of all deployed workers",
	},
}
