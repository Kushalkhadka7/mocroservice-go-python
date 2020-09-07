import grpc

import proto.laptopService_pb2 as pb
import proto.laptopService_pb2_grpc as pb_grpc

def run():
    with grpc.insecure_channel("localhost:8080") as channel:
        stub = pb_grpc.LaptopServiceStub(channel)
        hello_req = pb.Hello(Hello="hello from python client.")

        response = stub.SayHello(hello_req)

        print(response)


if __name__ == "__main__":
    run()