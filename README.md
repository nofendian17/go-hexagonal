This is an example directory structure for a Go application that follows the Clean Architecture principles. Here's a brief description of each directory:

- `cmd`: This directory contains the main program that will run the application. It usually contains the `main` package and is responsible for initializing the application and running it.

- `configs`: This directory contains the configuration files for the application. This can include environment-specific configuration files, as well as files that contain secrets or other sensitive information.

- `database`: This directory contains the database-related code, including the database migrations.

- `internal`: This directory contains the internal code of the application, which should not be exposed to the outside world. It is further divided into several subdirectories:

    - `adapters`: This directory contains the adapters that connect the application to the outside world. This can include things like HTTP handlers, database and cache clients, and other external integrations.

    - `core`: This directory contains the core business logic of the application, including the domain models, use cases, and services. This code should be independent of any external dependencies and should be able to be tested in isolation.

    - `middleware`: This directory contains the middleware code that can be used to wrap the HTTP handlers and perform common tasks like logging, authentication, and request validation.

    - `mocks`: This directory contains the mock implementations of interfaces used in tests.

    - `shared`: This directory contains the code that is shared across the application, including utilities, common errors, and other helpers.

Overall, this directory structure is designed to make it easy to separate the concerns of the application, making it more modular and easier to test and maintain.