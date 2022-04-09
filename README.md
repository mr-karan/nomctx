# nomctx

Faster way to switch between Nomad clusters and namespaces. Inspired from [kubectx](https://github.com/ahmetb/kubectx/).

## Why was this created

If you're operating multiple Nomad clusters, switching between these clusters is a cumbersome task of exporting variables in shell. Ideally `nomad`, should use a file similar to `~/.kube/config` for authenticating against multiple clusters as described in this [issue](https://github.com/hashicorp/nomad/issues/11043). Since this feature isn't available as of yet, I've created `nomctx` which can emit the [environment-variables](https://www.nomadproject.io/docs/commands#environment-variables) required by `nomad` CLI for authentication.

## Installation

Grab the latest release from [Releases](https://github.com/mr-karan/nomctx/releases).

To run:

```
$ nomctx
```


## Usage

### Interactive Mode

If you have `fzf` installed, the tool will show an interactive prompt for switching `clusters` or `namespace`:


### Non Interactive Mode



## Configuration

Here's a sample config file which shows 2 clusters: `prod` and `dev`:

```hcl
clusters "dev" {
  address   = "http://10.0.0.1:4646"
  http_auth = "user:pass"
  namespace = "default"
  region    = "abc"
  token     = "26a57a4c-1fe4-4220-a60b-576ea637100a"
}

clusters "prod" {
  address   = "http://127.0.0.1:4646"
  namespace = "default"
  token     = "f8cb5774-749a-4548-acc9-054df3b52e83"
}
```

By default `nomctx` searches for the file in `~/.nomctx/config.hcl` but you can override that with `--config=</path/to/config.hcl>` flag.