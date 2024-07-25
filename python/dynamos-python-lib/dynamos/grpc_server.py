import grpc
from .logger import InitLogger
from .base_client import BaseClient
from .rabbit_client import RabbitClient
from concurrent import futures
import time
import threading
from opentelemetry import trace
from google.protobuf.empty_pb2 import Empty
from typing import Callable, Dict

import health_pb2_grpc as healthServer
import health_pb2 as healthTypes
import microserviceCommunication_pb2_grpc as msCommServer
import microserviceCommunication_pb2 as msCommTypes

# Configure logging
logger = InitLogger()

CallbackType = Callable[[grpc.ServicerContext, msCommTypes.MicroserviceCommunication], Empty]


class HealthServicer(healthServer.HealthServicer):
    def Check(self, request, context):
        logger.info("Received health check request")
        return healthTypes.HealthCheckResponse(
            status=healthTypes.HealthCheckResponse.SERVING
        )

    def Watch(self, request, context):
        logger.info("Received health watch request")
        yield healthTypes.HealthCheckResponse(
            status=healthTypes.HealthCheckResponse.SERVING
        )



class MicroserviceServicer(msCommServer.MicroserviceServicer):
    def __init__(self):
        self.callbacks: Dict[int, CallbackType] = {}

    def SendData(self, msComm):
        logger.debug(f"Starting grpc_server.py/SendData: {msComm.RequestMetadata.DestinationQueue}")

        span = trace.get_current_span()
        try:
            # Start a new span
            with trace.get_tracer(__name__).start_as_current_span("grpc_server.py/SendData") as span:
                pass
        except Exception as err:
            logger.warn(f"Error starting span: {err}")
            span.end()

        callback = self.callbacks.get(request.Type)
        if not callback:
            logger.warn("no callback registered for this message type")
            return Empty()

        logger.debug("In SendData function:")
        try:
            callback(context, request)
        except Exception as err:
            logger.error(f"Callback Error: {err}")
            return Empty()

        return Empty()

    def register_callback(self, message_type: int, callback: CallbackType) -> None:
        self.callbacks[message_type] = callback


class GRPCServer:
    def __init__(self, grpc_addr):
        self.grpc_addr = grpc_addr

        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        healthServer.add_HealthServicer_to_server(HealthServicer(), self.server)
    #     # rabbitServer.add_RabbitServicer_to_server(RabbitServicer(), self.server)
    #     # etcdServer.add_EtcdServicer_to_server(EtcdServicer(), self.server)
    #     # .add_MicroserviceServicer_to_server(MicroserviceServicer(), self.server)

        self.server.add_insecure_port(self.grpc_addr)
        self.stop_event = threading.Event()
        self.condition = threading.Condition()


    def start_server(self):
        self.server.start()
        logger.info(f"gRPC server started on {self.grpc_addr}")
        with self.condition:
            while not self.stop_event.is_set():
                self.condition.wait()  # Wait for the signal to stop
        self.server.stop(0)
        logger.info("Server stopped")


    def start(self):
        self.thread = threading.Thread(target=self.start_server)
        self.thread.daemon = True
        self.thread.start()


    def stop(self):
        logger.info("Stopping gRPC server...")
        with self.condition:
            self.stop_event.set()
            self.condition.notify()  # Notify the condition to wake up the server
        self.thread.join()
        logger.info("gRPC server stopped")