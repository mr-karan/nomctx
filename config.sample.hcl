cluster "local" {
  address   = "http://127.0.0.1:4646"
  namespace = "default"
}

cluster "bangalore" {
  address = "https://nomad.hashicorp.rocks"
  auth {
    method   = "gitlab"
    provider = "nomad"
  }
}

cluster "tokyo" {
  address   = "http://10.0.0.1:4646"
  http_auth = "user:pass"
  namespace = "pink"
  token     = "c0a7d714-46df-4c6e-954a-269578c3804d"
}
