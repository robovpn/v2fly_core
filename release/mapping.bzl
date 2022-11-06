def gen_mappings(os, arch, ver):
  return {
    "github.com/robovpn/v2fly_core/release/config": "",
    "github.com/robovpn/v2fly_core/main/" + os + "/" + arch + "/" + ver: "",
    "github.com/robovpn/v2fly_core/infra/control/main/" + os + "/" + arch + "/" + ver : "",
  }
