# Basics-Go-middleware

Explanation:

    Middleware Function (loggingMiddleware):
        Takes a http.Handler as input (the next handler in the chain).
        Returns a new http.Handler that wraps the original one.
        Logs the request before and after calling the next handler.

    Applying Middleware:
        http.Handle("/", loggingMiddleware(http.HandlerFunc(rootHandler)))
            Wraps the rootHandler with loggingMiddleware.
            Registers the wrapped handler for the root path (/).

    Execution Flow:
        When a request arrives for /:
            loggingMiddleware logs the request.
            It calls next.ServeHTTP(w, r), passing control to rootHandler.
            rootHandler writes "Hello, world!".
            Control returns to loggingMiddleware, which logs the processed request.

Key Points:

    Middleware functions can modify requests, responses, or both.
    They can be chained together for various purposes (e.g., authentication, logging, access control).
    Go's http.HandlerFunc type facilitates middleware implementation.
    Middleware promotes code reusability and separation of concerns.

Common Use Cases:

    Authentication and authorization
    Logging and request tracking
    Error handling
    Rate limiting
    Content compression
    CORS handling
    Session management