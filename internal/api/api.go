package api

type Api interface {
	// Assumes we are using the default mux, etc.
	// If we wanted customizations we'd have to add parameters.
	RegisterEndpoints()
}
