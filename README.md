# atmos-cli

atmos-cli is a CLI to fetch divelogs from ATMOS Platform.

![](https://github.com/umatare5/atmos-cli/blob/images/promo.gif)

## Syntax

```bash
$ atmos
NAME:
   atmos - Client for ATMOS Platform

USAGE:
   atmos COMMAND [options...]

VERSION:
   0.1.0

COMMANDS:
   profile, p, pr, pro, prof  Fetch my profile.
   divelog, d, di, div, dive  Fetch my divelogs.
   help, h                    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --server value, -s value  URI to ATMOS API (default: "https://localhost/") [$ATMOS_API]
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

- Environment Variables

  | Environment Varible | Alternative to      | Description                          |
  | ------------------- | ------------------- | ------------------------------------ |
  | ATMOS_API           | --server (-s)       | URI of ATMOS Platform API.           |
  | ATMOS_USERID        | --user-id (-i)      | Unique ID in ATMOS Platform.         |
  | ATMOS_TOKEN         | --access-token (-t) | The token to use ATMOS Platform API. |

## Usage

### Tutorial

- Show the latest 10 dive-logs in pretty format.

  ```bash
  atmos divelog --pretty
  ```

- Show the latest 5 dive-logs in pretty format.

  ```bash
  atmos divelog --pretty --limit 5
  ```

- Show the cursor at end of latest 10 dive-logs.

  ```bash
  atmos divelog | jq -r '.response.page_info.end_cursor'
  ```

- Show 10 dive-logs after selected cursor in pretty format.

  ```bash
  atmos divelog --pretty --cursor $END_CURSOR_NAME
  ```

## Advanced

- `atmos divelog --divelog-id` shows the detail of a divelog.

  ```bash
  atmos divelog --divelog-id <randomChar> | jq .
  ```

  ```bash
  {
    "code": "0000000",
    "message": "",
    "response": {
      "divelog": {
        "air_in": 0,
        "air_in_text": "0",
        "air_out": 0,
        "air_out_text": "0",
        "air_temperature": "27.1",
        <snip>
      }
    }
  }
  ```

- `atmos profile` shows user profile.

  ```bash
  atmos profile | jq .
  ```

  ```bash
  {
    "code": "0000000",
    "message": "",
    "response": {
      "user": {
        "user_name": "umatare5",
        "status": "active",
        <snip>
      }
    }
  }
  ```

- `atmos profile --statistics` shows user statistics.

  ```bash
  atmos divelog --statistics | jq .
  ```

  ```bash
  {
    "code": "0000000",
    "message": "",
    "response": {
      "dive_time": 140437,
      "divelog_count": 42,
      "follower_count": 0,
      "following_count": 0,
      <snip>
    }
  }
  ```
