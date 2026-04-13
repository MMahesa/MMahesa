# uptime-probe

Go CLI utility for checking HTTP and TCP targets concurrently. Built as a network/NOC portfolio project with automation-friendly output.

## Highlights

- Concurrent checks for multiple targets
- Supports HTTP and raw TCP endpoints
- Supports target lists from file
- Supports JSON config input
- Table and JSON output
- Timeout control through flags
- Retry support per target
- Concurrency control for worker execution
- Summary output for quick monitoring checks
- Exit code `2` when one or more targets are down
- Basic test coverage

## Run

```bash
go run ./cmd/uptime-probe --timeout=2s --retries=2 --concurrency=4 https://example.com 1.1.1.1:53
```

## File Input

```bash
go run ./cmd/uptime-probe --file=targets.txt
```

## JSON Config

```json
{
  "targets": [
    { "address": "https://example.com" },
    { "address": "1.1.1.1:53" }
  ]
}
```

```bash
go run ./cmd/uptime-probe --config=targets.json --format=json
```

## JSON Output

```bash
go run ./cmd/uptime-probe --format=json https://example.com 8.8.8.8:53
```
