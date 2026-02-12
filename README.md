# timetable

A lightweight web app that displays your [Toggl Track](https://toggl.com/track/) time entries as a two-week timetable, broken down by project and day. Built with Go and HTMX.

## Running

### Locally

```sh
go run .
```

### Docker

```sh
docker build -t timetable .
docker run -p 8080:8080 timetable
```

The app will be available at `http://localhost:8080`.

## Usage

1. Open the app in your browser.
2. Enter your Toggl API token and click the button.
3. Your time entries from the last 14 days are displayed in a table with projects as rows and dates as columns.

You can find your API token in your [Toggl profile settings](https://track.toggl.com/profile).
