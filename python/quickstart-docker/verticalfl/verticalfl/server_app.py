from flwr.common import Context
from flwr.server import ServerApp, ServerAppComponents, ServerConfig

from verticalfl.strategy import Strategy
from verticalfl.task import process_dataset


def server_fn(context: Context) -> ServerAppComponents:
    """Construct components that set the ServerApp behaviour."""

    # Get dataset
    processed_df, _ = process_dataset()

    # Define the strategy
    strategy = Strategy(processed_df["Survived"].values)

    # Construct ServerConfig
    num_rounds = context.run_config["num-server-rounds"]
    config = ServerConfig(num_rounds=num_rounds)
    print("Creating the server")

    return ServerAppComponents(strategy=strategy, config=config)


# Start Flower server
app = ServerApp(server_fn=server_fn)
