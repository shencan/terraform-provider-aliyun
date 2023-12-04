terraform {
  required_providers {
    aliyun = {
      source = "shencan/aliyun"
      version = "0.0.2"
    }
  }
}

provider "aliyun" {
  ak = "ak"
  sk = "sk"
  region = "cn-hongkong"
}

resource "aliyun_ess_tag" "demo" {
  ess_id = "asg-id"
  tags = {
    "a" = "bb"
    "c" = "dd"
  }
}