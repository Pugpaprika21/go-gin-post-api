package routes

import (
	"github.com/Pugpaprika21/go-gin/controller"
	"github.com/Pugpaprika21/go-gin/middleware"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		userRouter := apiRouter.Group("/user")
		{
			userRepository := repository.NewUserRepository()
			userController := controller.User{UserRepositoryInterface: userRepository}
			userRouter.GET("/users", userController.GetUserAll)
			userRouter.POST("/register", userController.Register)
			userRouter.POST("/login", userController.Login)
		}

		postRouter := apiRouter.Group("/post", middleware.AuthJWTProtected())
		{
			postRepository := repository.NewPostRepository()
			postController := controller.Post{PostRepositoryInterface: postRepository}
			postRouter.GET("/posts/*postId", postController.GetPostAll)
			postRouter.GET("/assets/uploads/:filename", postController.ShowPostAttachments)
			postRouter.POST("/create", postController.CreatePost)
			postRouter.GET("/get-post/user/:userId", postController.GetPostByUser)
			postRouter.DELETE("/delete/post/:postId", postController.DeletePost)
			postRouter.PUT("/update/post/:postId", postController.UpdatePost)
		}

		commentRouter := apiRouter.Group("/comment", middleware.AuthJWTProtected())
		{
			commentRepository := repository.NewCommentRepository()
			commentController := controller.Comment{CommentRepositoryInterface: commentRepository}
			commentRouter.POST("/create", commentController.CreateComment)
			commentRouter.GET("/comments", commentController.GetCommentAll)
			commentRouter.GET("/comment/:commentId/user/:userId/action/:action", commentController.GetComment)
			commentRouter.PUT("/update/:commentId", commentController.UpdateComment)
			commentRouter.DELETE("/delete/:commentId", commentController.DeleteComment)
		}
	}
}
