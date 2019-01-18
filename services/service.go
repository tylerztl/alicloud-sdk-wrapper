package services

// define all kinds of services

type Provider interface {
	CreateInstance(request Request) Response
	RunInstances(request Request) Response
}

