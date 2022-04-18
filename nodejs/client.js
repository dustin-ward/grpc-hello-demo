var PROTO_PATH = __dirname + '/../chat.proto';

var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var chat_proto = grpc.loadPackageDefinition(packageDefinition).chat;

function main() {
  var target = 'zoscan2b.pok.stglabs.ibm.com:9000';

  var client = new chat_proto.ChatService(target, grpc.credentials.createInsecure());

  client.sayHello({body: "Hello from Nodejs client"}, function(err, response) {
    console.log('Response from server:', response.body);
  });
}

main();
