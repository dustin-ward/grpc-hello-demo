#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>

#ifdef BAZEL_BUILD
#include "./chat.grpc.pb.h"
#else
#include "./chat.grpc.pb.h"
#endif

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using chat::ChatService;
using chat::Message;

class ChatClient {
 public:
  ChatClient(std::shared_ptr<Channel> channel)
      : stub_(ChatService::NewStub(channel)) {}

  // Assembles the client's payload, sends it and presents the response back
  // from the server.
  std::string SayHello(const std::string& body) {
    // Data we are sending to the server.
    Message message;
    message.set_body(body);

    // Container for the data we expect from the server.
    Message reply;

    // Context for the client. It could be used to convey extra information to
    // the server and/or tweak certain RPC behaviors.
    ClientContext context;

    // The actual RPC.
    Status status = stub_->SayHello(&context, message, &reply);

    // Act upon its status.
    if (status.ok()) {
      return reply.body();
    } else {
      std::cout << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "RPC failed";
    }
  }

 private:
  std::unique_ptr<ChatService::Stub> stub_;
};

int main(int argc, char** argv) {
  std::string target_str = "localhost:9000";

  ChatClient chat(
      grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials()));
  std::string reply = chat.SayHello("Hello from C++ client");
  std::cout << "Message received: " << reply << std::endl;

  return 0;
}