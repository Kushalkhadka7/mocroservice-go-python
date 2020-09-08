import grpc
import threading

import proto.laptopService_pb2 as pb
import proto.laptopService_pb2_grpc as pb_grpc

class LaptopClient:

    def __init__(self):
        channel = grpc.insecure_channel("localhost:8080")
        self.conn = pb_grpc.LaptopServiceStub(channel)
        threading.Thread(target=self.fetch_laptops, daemon=True).start()

    def create_laptop(self):
        response = self.conn.CreateLaptop()

        print(response)
     
    def fetch_laptops(self):
        for response in self.conn.FetchAllLaptops(pb.Void()):
            print(response)

            
def run():
    stub = LaptopClient()
    stub.fetch_laptops()


if __name__ == "__main__":
    run()