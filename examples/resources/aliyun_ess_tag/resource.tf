resource "aliyun_ess_tag" "demo" {
  ess_id = "asg-id"
  tags = {
    "a" = "bb"
    "c" = "dd"
  }
}