# wsl-ssh-amb

This little program invokes `wsl ssh` or `ssh.exe` by detect ssh-target has special prefix (or not).

## Build

Please build for Windows target.

In WSL (or in Linux)

```bash
GOOS=windows GOARCH=amd64 go build .
```

In Windows

```bash
go build .
```

## Settings

### VSCode

`Settings.json`

```json
{
  "remote.SSH.path": "C:\\path\\to\\wsl-ssh-amb.exe",
  "remote.SSH.remoteServerListenOnSocket": true
}
```

I don't know what `remoteServerListenOnSocket` actually does.
But, without this, Remote-SSH fails with various errors.

- failed to install vscode-server to remote
- failed to establish socket between local and remote
- etc.

### SSH config

You need to add dummy target to Windows side config.
Dummy must be correspond to WSL side config.

For example, if you set target like this in WSL,

```plaintext
Host foo
    HostName foo.example.com
    User tbistr
    IdentityFile ~/.ssh/id_rsa
Host bar
    HostName bar.example.com
    User tbistr
    IdentityFile ~/.ssh/id_rsa
    ProxyCommand ssh -W %h:%p foo
```

add dummy like this in Windows.

```plaintext
Host wsl_foo
Host wsl_bar
```

The default prefix is `wsl_`.

## Background

VSCode Remote-SSH Extension uses Windows `ssh.exe` and config file.
You cant use WSL `ssh`, `.ssh/config` and keys.
The problem was reported by [this issue](https://github.com/microsoft/vscode-remote-release/issues/937).

Great workaround is provided by [this gist](https://gist.github.com/diablodale/54756043c395d712053cf0d50a86086a).
However, this approach completely switches to use `wsl ssh`.
**I want to switch `wsl ssh` and `ssh.exe` depend on target.**
(I use vagrant (virtualbox) on Windows and use `vagrant ssh-config >> ~/.ssh/config` to access VMs.)

This little program invokes `wsl ssh` or `ssh.exe` by detect ssh-target has special prefix (or not).
