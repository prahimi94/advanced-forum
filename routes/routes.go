package routes

import (
	"forum/middlewares"
	forumManagementControllers "forum/modules/forumManagement/controllers"
	userManagementControllers "forum/modules/userManagement/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	// Initialize a new router
	router := mux.NewRouter()

	// Serve static files
	router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("assets/")))
	router.PathPrefix("/js/").Handler(http.FileServer(http.Dir("assets/")))
	router.PathPrefix("/img/").Handler(http.FileServer(http.Dir("assets/")))
	router.PathPrefix("/uploads/").Handler(http.FileServer(http.Dir("assets/")))

	// Public routes (directly registered)
	router.HandleFunc("/", forumManagementControllers.MainPageHandler).Methods("GET")
	router.HandleFunc("/auth/", userManagementControllers.AuthHandler).Methods("GET")
	router.HandleFunc("/logout/", userManagementControllers.Logout).Methods("GET")
	router.HandleFunc("/register", userManagementControllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", userManagementControllers.LoginHandler).Methods("POST")
	router.HandleFunc("/post/{id}", forumManagementControllers.ReadPost).Methods("GET")
	router.HandleFunc("/posts/{categoryName}", forumManagementControllers.ReadPostsByCategory).Methods("GET")
	router.HandleFunc("/filterPosts", forumManagementControllers.FilterPosts).Methods("GET")

	// Protected routes (using middleware)
	protected := router.PathPrefix("/").Subrouter() // No need to add "/protected"
	protected.Use(middlewares.AuthMiddleware)       // Apply AuthMiddleware to protected routes
	protected.HandleFunc("/newPost/", forumManagementControllers.CreatePost).Methods("GET", "POST")
	protected.HandleFunc("/submitPost", forumManagementControllers.SubmitPost).Methods("POST")
	protected.HandleFunc("/editPost", forumManagementControllers.EditPost).Methods("GET", "POST")
	protected.HandleFunc("/updatePost", forumManagementControllers.UpdatePost).Methods("POST")
	protected.HandleFunc("/deletePost", forumManagementControllers.DeletePost).Methods("POST")
	protected.HandleFunc("/myCreatedPosts/", forumManagementControllers.ReadMyCreatedPosts).Methods("GET")
	protected.HandleFunc("/myLikedPosts/", forumManagementControllers.ReadMyLikedPosts).Methods("GET")
	protected.HandleFunc("/likePost", forumManagementControllers.LikePost).Methods("POST")
	protected.HandleFunc("/likeComment", forumManagementControllers.LikeComment).Methods("POST")
	protected.HandleFunc("/submitComment", forumManagementControllers.SubmitComment).Methods("POST")
	protected.HandleFunc("/updateComment", forumManagementControllers.UpdateComment).Methods("POST")
	protected.HandleFunc("/deleteComment", forumManagementControllers.DeleteComment).Methods("POST")

	return router
}
