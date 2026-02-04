# GitHub User Activity

A simple command-line interface (CLI) application to fetch and display the recent activity of a GitHub user directly in the terminal.

**Project Source:** https://roadmap.sh/projects/github-user-activity

## Overview

This project demonstrates how to work with APIs, handle JSON data, and build a simple CLI application in Go. The application fetches a user's recent GitHub events and displays them in an organized, readable format in the terminal.

## Features

- ✅ Fetch recent GitHub user activity using the GitHub API
- ✅ Display activities in an organized format
- ✅ Support for custom activity count (defaults to 5)
- ✅ Error handling for invalid usernames and API failures
- ✅ Grouped activity display showing repeated events with counts
- ✅ Support for multiple GitHub event types:
  - CreateEvent (Repository creation)
  - PushEvent (Code pushes)
  - PullRequestEvent (Pull requests)
  - IssuesEvent (Issues opened)
  - WatchEvent (Repository stars)
  - ForkEvent (Repository forks)
  - DeleteEvent (Deletions)
  - ReleaseEvent (Releases)

## Installation

### Prerequisites

- Go 1.11 or higher
- Internet connection (to fetch GitHub API data)

### Build from Source

```bash
cd github_cli
go build -o github_cli main.go
```

## Usage

### Basic Usage

Display the last 5 recent activities for a user:

```bash
./github_cli <username>
```

**Example:**
```bash
./github_cli torvalds
```

### With Custom Activity Count

Display a custom number of recent activities:

```bash
./github_cli <username> <count>
```

**Example:**
```bash
./github_cli torvalds 10
```

## Example Output

```
Pushed to torvalds/linux (6 times)
Pushed to torvalds/subsurface-for-dirk (2 times)
Created pull request in torvalds/linux
Starred torvalds/linux-stable
```

## How It Works

1. **Command-line Parsing**: Accepts GitHub username and optional activity count
2. **API Request**: Fetches public events from `https://api.github.com/users/<username>/events/public`
3. **JSON Parsing**: Unmarshals the JSON response into Activity structs
4. **Activity Aggregation**: Groups consecutive identical event types and repositories
5. **Display**: Shows formatted activity with repetition counts

## Error Handling

The application handles the following error cases:

- Missing username argument
- Invalid input format
- Non-existent GitHub username (404 error)
- API request failures
- JSON parsing errors

## API Reference

This project uses the GitHub REST API's public events endpoint:

```
GET https://api.github.com/users/<username>/events/public
```

For more information, see the [GitHub API Documentation](https://docs.github.com/en/rest/activity/events?apiVersion=2022-11-28).

## Future Enhancements

Possible improvements for more advanced versions:

- Filter activities by event type
- Display activities in different formats (JSON, CSV, etc.)
- Cache fetched data to improve performance
- Fetch additional user information from other API endpoints
- Add configuration file support
- Display activity with timestamps
- Export activity to file

## Technology Stack

- **Language**: Go
- **External Dependencies**: None (uses only standard library)
- **APIs**: GitHub REST API

## License

This project is part of the roadmap.sh backend development learning path.
