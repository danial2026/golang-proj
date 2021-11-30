package app

import (
	controllers "github.com/danial2026/file-sharing-go/controllers/files"
)

func mapUrls() {
	prefix := "/api/v1/files"

	router.GET(prefix+"/download", controllers.FilesController.Log)
	router.POST(prefix+"/uploadUserPic", controllers.FilesController.StoreUserPic)
	router.POST(prefix+"/uploadUserPicBase64", controllers.FilesController.StoreUserPicBase64)
	router.GET(prefix+"/downloadFile", controllers.FilesController.DownloadFileHandler)
}
