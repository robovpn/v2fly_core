package(default_visibility = ["//visibility:public"])

load("//infra/bazel:zip.bzl", "pkg_zip")
load("//release:mapping.bzl", "gen_mappings")

filegroup(
    name = "config_json",
    srcs = [
        "config/config.json",
        "config/vpoint_socks_vmess.json",
        "config/vpoint_vmess_freedom.json",
    ],
)

filegroup(
    name = "systemd",
    srcs = [
        "config/systemd/system/v2fly.service",
        "config/systemd/system/v2fly@.service",
    ],
)

filegroup(
    name = "geodata",
    srcs = [
        "config/geoip.dat",
        "config/geosite.dat",
    ],
)

pkg_zip(
    name = "v2fly_darwin_amd64_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_darwin_amd64",
        "//main:v2fly_darwin_amd64",
    ],
    out = "v2fly-macos-64.zip",
    mappings = gen_mappings("darwin", "amd64", "0"),
)

pkg_zip(
    name = "v2fly_windows_amd64_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_windows_amd64",
        "//main:v2fly_windows_amd64",
        "//main:v2fly_windows_amd64_nowindow",
    ],
    out = "v2fly-windows-64.zip",
    mappings = gen_mappings("windows", "amd64", "0"),
)

pkg_zip(
    name = "v2fly_windows_x86_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_windows_386",
        "//main:v2fly_windows_386",
        "//main:v2fly_windows_386_nowindow",
    ],
    out = "v2fly-windows-32.zip",
    mappings = gen_mappings("windows", "386", "0"),
)

pkg_zip(
    name = "v2fly_windows_armv7_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_windows_arm_7",
        "//main:v2fly_windows_arm_7",
        "//main:v2fly_windows_arm_7_nowindow",
    ],
    out = "v2fly-windows-arm32-v7a.zip",
    mappings = gen_mappings("windows", "arm", "7"),
)

pkg_zip(
    name = "v2fly_freebsd_amd64_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_freebsd_amd64",
        "//main:v2fly_freebsd_amd64",
    ],
    out = "v2fly-freebsd-64.zip",
    mappings = gen_mappings("freebsd", "amd64", "0"),
)

pkg_zip(
    name = "v2fly_freebsd_x86_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_freebsd_386",
        "//main:v2fly_freebsd_386",
    ],
    out = "v2fly-freebsd-32.zip",
    mappings = gen_mappings("freebsd", "386", "0"),
)

pkg_zip(
    name = "v2fly_openbsd_amd64_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_openbsd_amd64",
        "//main:v2fly_openbsd_amd64",
    ],
    out = "v2fly-openbsd-64.zip",
    mappings = gen_mappings("openbsd", "amd64", "0"),
)

pkg_zip(
    name = "v2fly_openbsd_x86_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_openbsd_386",
        "//main:v2fly_openbsd_386",
    ],
    out = "v2fly-openbsd-32.zip",
    mappings = gen_mappings("openbsd", "386", "0"),
)

pkg_zip(
    name = "v2fly_dragonfly_amd64_package",
    srcs = [
        ":config_json",
        ":geodata",
        "//infra/control/main:v2ctl_dragonfly_amd64",
        "//main:v2fly_dragonfly_amd64",
    ],
    out = "v2fly-dragonfly-64.zip",
    mappings = gen_mappings("dragonfly", "amd64", "0"),
)

pkg_zip(
    name = "v2fly_linux_amd64_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_amd64",
        "//main:v2fly_linux_amd64",
    ],
    out = "v2fly-linux-64.zip",
    mappings = gen_mappings("linux", "amd64", "0"),
)

pkg_zip(
    name = "v2fly_linux_x86_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_386",
        "//main:v2fly_linux_386",
    ],
    out = "v2fly-linux-32.zip",
    mappings = gen_mappings("linux", "386", "0"),
)

pkg_zip(
    name = "v2fly_linux_arm64_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_arm64",
        "//main:v2fly_linux_arm64",
    ],
    out = "v2fly-linux-arm64-v8a.zip",
    mappings = gen_mappings("linux", "arm64", "0"),
)

pkg_zip(
    name = "v2fly_linux_armv7_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_arm_7",
        "//main:v2fly_linux_arm_7",
    ],
    out = "v2fly-linux-arm32-v7a.zip",
    mappings = gen_mappings("linux", "arm", "7"),
)

pkg_zip(
    name = "v2fly_linux_armv6_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_arm_6",
        "//main:v2fly_linux_arm_6",
    ],
    out = "v2fly-linux-arm32-v6.zip",
    mappings = gen_mappings("linux", "arm", "6"),
)

pkg_zip(
    name = "v2fly_linux_armv5_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_arm_5",
        "//main:v2fly_linux_arm_5",
    ],
    out = "v2fly-linux-arm32-v5.zip",
    mappings = gen_mappings("linux", "arm", "5"),
)

pkg_zip(
    name = "v2fly_linux_mips32_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_mips",
        "//infra/control/main:v2ctl_linux_mips_softfloat",
        "//main:v2fly_linux_mips",
        "//main:v2fly_linux_mips_softfloat",
    ],
    out = "v2fly-linux-mips32.zip",
    mappings = gen_mappings("linux", "mips", "0"),
)

pkg_zip(
    name = "v2fly_linux_mips32le_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_mipsle",
        "//infra/control/main:v2ctl_linux_mipsle_softfloat",
        "//main:v2fly_linux_mipsle",
        "//main:v2fly_linux_mipsle_softfloat",
    ],
    out = "v2fly-linux-mips32le.zip",
    mappings = gen_mappings("linux", "mipsle", "0"),
)

pkg_zip(
    name = "v2fly_linux_mips64_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_mips64",
        "//main:v2fly_linux_mips64",
    ],
    out = "v2fly-linux-mips64.zip",
    mappings = gen_mappings("linux", "mips64", "0"),
)

pkg_zip(
    name = "v2fly_linux_mips64le_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_mips64le",
        "//main:v2fly_linux_mips64le",
    ],
    out = "v2fly-linux-mips64le.zip",
    mappings = gen_mappings("linux", "mips64le", "0"),
)

pkg_zip(
    name = "v2fly_linux_riscv64_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_riscv64",
        "//main:v2fly_linux_riscv64",
    ],
    out = "v2fly-linux-riscv64.zip",
    mappings = gen_mappings("linux", "riscv64", "0"),
)

pkg_zip(
    name = "v2fly_linux_s390x_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_s390x",
        "//main:v2fly_linux_s390x",
    ],
    out = "v2fly-linux-s390x.zip",
    mappings = gen_mappings("linux", "s390x", "0"),
)

pkg_zip(
    name = "v2fly_linux_ppc64_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_ppc64",
        "//main:v2fly_linux_ppc64",
    ],
    out = "v2fly-linux-ppc64.zip",
    mappings = gen_mappings("linux", "ppc64", "0"),
)

pkg_zip(
    name = "v2fly_linux_ppc64le_package",
    srcs = [
        ":config_json",
        ":geodata",
        ":systemd",
        "//infra/control/main:v2ctl_linux_ppc64le",
        "//main:v2fly_linux_ppc64le",
    ],
    out = "v2fly-linux-ppc64le.zip",
    mappings = gen_mappings("linux", "ppc64le", "0"),
)
