# Event Sourcing To-Do Example

This repository contains a demonstration of an event sourcing system applied to a to-do list management application, built using Go. It showcases how to structure and implement an application that leverages event sourcing principles to maintain the state of to-do lists and their tasks through events.

This example is based on the [eventsourcing package](https://github.com/thefabric-io/eventsourcing) to show how to use the latest to implement an event-sourced system.

## Features

- **Environment Configuration:** Utilizes `.env` files for environment configuration.
- **Event Sourcing Integration:** Includes initialization and use of an event sourcing system for managing to-dos and tasks.
- **CQRS Pattern:** Implements the Command Query Responsibility Segregation (CQRS) pattern to separate the read and write models of the application. The read model is not implemented in this example.

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Access to a PostgreSQL database for event storage
- [godotenv](https://github.com/joho/godotenv) library for loading environment variables

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/thefabric-io/eventsourcing.example.git
   cd eventsourcing.example
   ```

2. **Set up your `.env` file:**

   Create a `.env` file at the root of the project or copy the `.env.example` and customize it with your environment variables.

3. **Run the application:**

   Execute the main program:

   ```sh
   go run main.go
   ```

## Usage

This application is a demonstration and does not include a CLI or web interface. It is designed to be run directly to observe the behavior of an event-sourced system within a Go application. To modify the behavior or add tasks, adjust the `main.go` file and re-run the application.

The application is designed for plug-and-play. By entering the PostgreSQL connection string into your `.env` file located at the project's root, you can get the application up and running by simply executing the main program.

## How It Works

The application initializes a CQRS-based architecture with a focus on event sourcing. The main file demonstrates the creation of a to-do item and the addition of a task to that to-do item through commands. Each command generates events (in our case one event per command) that are stored and can be replayed to rebuild the state of the application at any point.

You might consider adding a command that allows for the creation of a new task and its addition to a to-do list simultaneously. In the application layer, this would necessitate the handler to invoke `todo.AddTask` subsequent to `todo.Create`, all within a single transaction boundary.
### Key Concepts Demonstrated

- **Event Sourcing Initialization:** How to set up and initialize an event sourcing system.
- **Command Handling:** Handling commands for creating to-dos and adding tasks.
- **Event Storage and Replay:** Storing events and using them to rebuild application state.

## Contributing

Contributions are welcome! Please feel free to submit pull requests, suggest features, or report bugs.

## License

This project is open-source and available under the MIT License.