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
