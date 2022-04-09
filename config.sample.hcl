clusters "local" {
  address   = "http://127.0.0.1:4646"
  namespace = "default"
}

clusters "bangalore" {
  address   = "http://10.0.0.1:4646"
  http_auth = "user:pass"
  namespace = "pink"
  token     = "c0a7d714-46df-4c6e-954a-269578c3804d"
}

clusters "tokyo" {
  address   = "http://10.0.0.2:4646"
  namespace = "purple"
  token     = "08d63db2-b630-43c3-b614-6b7a6a553187"
}

clusters "paris" {
  address   = "http://10.0.0.3:4646"
  namespace = "blue"
  region    = "paris"
  token     = "f7a50344-67ab-4b6c-bf10-14f31ccba025"
}

clusters "singapore" {
  address   = "http://10.0.0.4:4646"
  namespace = "yellow"
  token     = "efb54871-b56e-4516-a70e-dee5c3bbf102"
}
