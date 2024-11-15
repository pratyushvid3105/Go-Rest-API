package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pratyushvid3105/Go-Rest-API/utils"
)

// it's important to keep in mind that as mentioned, this request handler will be executed in the middle of a request. And actually the way gin works, and the way we register this function, is such that other request handlers after this handler here, would still run. And therefore we would essentially try to send multiple responses if we also send the responses in Authenticate function.
// Instead, what we should do is use the special AbortWithStatusJSON method that's provided by the gin package on this context. Because this does what the name implies. It aborts the current request, so to say, or the current response generation, and sends response put inside the method, and no other request handlers thereafter will be executed. And that's what we need in Authenticate function, to make sure that if something goes wrong here, we stop and no other code on the server runs.
func Authenticate(context *gin.Context){
	token := context.Request.Header.Get(("Authorization"))

	if token == "" {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not authorized."})
		return
	}

	// we still need to get the userId from here to createEvent handler in events.go file. And that's thankfully quite straightforward to do because Gin, this web framework we're using here gives us a special method we can call on this context value on which we're operating here, and keep in mind that we got a pointer to that value so we're always working with exactly the same value in memory, no matter in which middleware function or handler function we're using the context. And that special method it exposes is the Set method. 
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized", "error": err.Error()})
		return
	}

	// Set method is simply a method that allows us to add some data to this context value, some data that is then attached to the context and hence, can be used anywhere where the context is available. Now, Set wants two parameters, a key, so an identifier, and then the value that should be identifiable through that key.
	context.Set("userId", userId)
	
	// Now if we make it to the end of this Authenticate function though, if we got a valid token, we also should call a special function on context, and that's the Next function. And this will then ensure that the next request handler in line will execute correctly.
	context.Next()
}