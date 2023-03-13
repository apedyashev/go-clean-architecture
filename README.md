# Clean Architecture 
Go implementation of the Clean Architecture pattern
![The Clean Architecture](docs/clean-architect.webp "The Clean Architecture")

[More info](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31)

## Rules
* The dependencies can only point inward (inner layer doesn't know anything about outer layers)
* Separation of details and abstracts. details  are the data, framework, database, and APIs. Using the details in the core layer means it violates The Dependency Rule

### Dependency Inversion Principle
Inner layers communicate to outer layer via interfaces. These interfaces are the defined in the inner layer, but outer layer has to implement them

## Layers
### Entities
Entities are domain models that have enterprise business rules and can be a set of data structures and functions.

### Use Cases
Use cases handle application business logic and have the Input Port (interface that use case must implement) and the Output Port (outputs data from server to client).  The ports are defined as interfaces in this level. Outer levels implement them.
The Output port is not necessary, because in the most frameworks controller allows to ouptut data (like `context.JSON()`)

Often other ports are needed, for example we might need a Repository, that allows to communicate to a DB or some external service. *On the level of usecase, we only define the interface of the repository*. The implementation will be done by the **Adapters Layer** (gateway)

### Interface adapters layer
* **Controllers** - depend on the usecase layer and execute processing through their interfaces
* **Presenters** - a set of specific implementations of the Output Port (can output JSON, HTML, XML, CSV etc)
* **Gateways** - a set of implementations for storing data in databases (like ORM, Repository)
* loggers, emails, push notifications etc

### Frameworks and Drivers
* databases
* frameworks
* routers

### Registry
The registry handles resolving dependencies between ports (usecases) and adapters (controllers, repositories), using constructor injection.

## Directories
|directory|layer|
|---------|-----|
|domain|entities|
|infrastructure|Frameworks and drivers|
|adapter|Interface adapters|
|usecase|Use cases|

Source: https://manakuro.medium.com/clean-architecture-with-go-bce409427d31

## Try it out
`go run .`

In Postman `GET http://localhost:3005/posts`