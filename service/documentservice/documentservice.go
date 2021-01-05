package service

var DocumentService DocumentServiceInterface = new(documentService)
type DocumentServiceInterface interface {
	CreateDocument()
}