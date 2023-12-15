Creating a simple RESTful API in Go involves setting up a web server, defining routes, and handling HTTP requests and responses. Several third-party libraries simplify this process. One popular choice is the "gin" framework. Here's an example of how you can create a simple RESTful API using Gin:

### 1. Install the Gin Framework:

```bash
go get -u github.com/gin-gonic/gin
```

### 2. Create a Simple RESTful API:

```go
// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the endpoint "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the RESTful API!",
		})
	})

	// Define a route for the endpoint "/api/resource"
	router.GET("/api/resource", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "This is a resource from the API.",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
```

### 3. Run the API:

```bash
go run main.go
```

This simple example creates a web server using the Gin framework and defines two routes:

1. The root route ("/") responds with a welcome message.
2. The "/api/resource" route responds with a sample resource message.

### Using Middleware with Gin:

Gin allows you to use middleware to perform actions before or after handling a request. Here's an example using Gin middleware for logging:

```go
// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("[%s] %s", c.Request.Method, c.Request.RequestURI)
		c.Next()
	}
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Use the Logger middleware
	router.Use(Logger())

	// Define routes as before

	// Run the server on port 8080
	router.Run(":8080")
}
```

In this example, the `Logger` middleware logs each incoming request before passing it to the next handler.

### Handling Parameters:

You can handle parameters from the URL or request body easily with Gin. Here's an example using URL parameters:

```go
// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route with a URL parameter
	router.GET("/api/user/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
```

### Third-Party Libraries:

1. **gin (github.com/gin-gonic/gin):**
   - Gin is a popular web framework for Go. It provides a fast and minimalist framework for building APIs with features like routing, middleware support, and request/response handling.

2. **Logrus (github.com/sirupsen/logrus):**
   - Logrus is a structured logger for Go. It's commonly used with Gin to add logging capabilities. In the example, the standard `log` package is used for simplicity, but Logrus is a great choice for more advanced logging requirements.

3. **viper (github.com/spf13/viper):**
   - Viper is a configuration management library for Go. It allows you to read configuration values from various sources, such as environment variables or configuration files. This can be handy when configuring your RESTful API.

Remember to use `go get` to install any additional third-party libraries you decide to include in your project.

This is a basic example, and as your API grows, you may want to explore more advanced topics such as database integration, middleware for authentication and authorization, request validation, and testing. The Go ecosystem offers many libraries to help with these aspects as well.