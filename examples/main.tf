terraform {
  required_version = ">=1.2.0"

  required_providers {
    xkcdpass = {
      source  = "advian-oss/xkcdpass"
      version = "~>1.0"
    }
  }
}


resource "xkcdpass_generate" "mypw" {
    length = 6
    capitalize = true
    separator = "-"
}

output "mypw" {
  value       = xkcdpass_generate.mypw.result
  sensitive   = true
}
