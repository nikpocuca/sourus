# Sourus

This is an all-in-one CLI monitor tool for memory, core load, and GPU utilization.
Why the name "Sourus"? It means lizard in greek, which is a play on "monitor lizard".

Theres a couple of releases for Intel-macos and Linux. 

## Colorful Themes! 

You can customize the theme of the cli app under `~/.sourus/settings.yml`. 

<img src="https://github.com/user-attachments/assets/00310bd6-8a58-44a0-8dfc-7c1e114c75b3" width="600" />

It will even show your GPU information! 

![sourus](https://github.com/user-attachments/assets/5a14ec6b-759e-4db1-917e-84d085b83e88)

You can also customize the size of the width of the app, and the arangement of the columnns. 

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


## Acknowledgments 

Many thanks to [Michael McCulloch](https://github.com/MichaelMcCulloch) for testing and providing feedback on v0.1.0. 
