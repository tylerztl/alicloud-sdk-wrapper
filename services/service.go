package services

// define all kinds of services

type Provider interface {
	CreateInstance()
	RunInstances()
}