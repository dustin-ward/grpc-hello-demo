licenses(["notice"])

cc_binary(
    name = "client",
    srcs = ["client.cc"],
    defines = ["BAZEL_BUILD"],
    deps = [
        "//:grpc++",
        "//grpc-demo:chat_cc_grpc",
    ],
)