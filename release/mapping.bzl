def gen_mappings(os, arch, ver):
  return {
    "v2fly_core/release/config": "",
    "v2fly_core/main/" + os + "/" + arch + "/" + ver: "",
    "v2fly_core/infra/control/main/" + os + "/" + arch + "/" + ver : "",
  }
