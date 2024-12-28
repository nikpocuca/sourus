# Sourus

This is an all-in-one CLI monitor tool for memory, core load, and GPU utilization.
Why the name "Sourus"? It means lizard in greek, which is a play on "monitor lizard".

Theres a couple of releases for Intel-macos and Linux. 

## Build 

Requires `go`-lang version 1.23+, and you need to pull in a few dependencies. The entire binary is only 4MB in size.

I have tested it out on MacOS and Linux with both GPU and non-GPU devices visible.

```bash
go build .
```

To run `sourus`, just call it via CLI:

```bash
./sourus 
╭────────────────────────────────────────────────────────────╮
│                       SOURUS MONITOR                       │
╰────────────────────────────────────────────────────────────╯
┌────────────────────────────────────────────────────────────┐
│                  Host Memory: 3.38G/23.7G                  │
│          █████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  14%          │
│                                                            │
│                                                            │
│                         Core Loads                         │
│             0.0% ░░░░░░░░░░      2.0% ░░░░░░░░░░           │
│             4.0% ░░░░░░░░░░      0.0% ░░░░░░░░░░           │
│             4.1% ░░░░░░░░░░      2.0% ░░░░░░░░░░           │
│             2.0% ░░░░░░░░░░      2.0% ░░░░░░░░░░           │
│             2.0% ░░░░░░░░░░      0.0% ░░░░░░░░░░           │
│             4.1% ░░░░░░░░░░      2.0% ░░░░░░░░░░           │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

You can add it to your `PATH` or place it in  `/usr/bin/`.
