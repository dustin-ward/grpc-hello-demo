from __future__ import print_function

import logging

import grpc
import chat_pb2
import chat_pb2_grpc


def run():
    with grpc.insecure_channel('localhost:9000') as channel:
        stub = chat_pb2_grpc.ChatServiceStub(channel)
        response = stub.SayHello(chat_pb2.Message(body='Hello from Python client'))
    print("Message received: " + response.body)


if __name__ == '__main__':
    logging.basicConfig()
    run()