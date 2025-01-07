# notification-cli

`notification-cli` is a simple Command Line Interface (CLI) tool for test the notification service gRPC.

## Features
- Send SMS notifications
- Send Email notifications
- Customizable server address

## Prerequisites
- Go installed on your system ([Download Go](https://go.dev/dl/))
- A running gRPC server for the notification service

## Installation
1. Clone the repository containing the `notification-cli` code.
2. Build the CLI tool:
   ```bash
   go build -o notification-cli
   ```
3. The compiled `notification-cli` binary will be available in the project directory.

## Usage

### General Syntax
```bash
./notification-cli --type <notification_type> --username <username> [--phone <phone_number>] [--email <email_address>] [--server <server_address>]
```

### Options
| Option            | Description                                      | Required                     |
|-------------------|--------------------------------------------------|------------------------------|
| `--type`          | Type of notification (`sms` or `email`)         | Yes                          |
| `--username`      | Username of the recipient                       | Yes                          |
| `--phone`         | Phone number (required for SMS)                 | Required for `--type sms`    |
| `--email`         | Email address (required for Email)              | Required for `--type email`  |
| `--server`        | Address of the gRPC server (default: localhost:50051) | No                     |

### Examples

#### Send an SMS Notification
```bash
./notification-cli --type sms --username "John Doe" --phone "123456789"
```

#### Send an Email Notification
```bash
./notification-cli --type email --username "Jane Doe" --email "jane.doe@example.com"
```

#### Specify a Custom Server Address
```bash
./notification-cli --type sms --username "John Doe" --phone "123456789" --server "127.0.0.1:50051"
```

## Error Handling
- If required flags are missing or invalid, the tool will exit with an error message.
- Ensure the gRPC server is running and reachable at the specified address.