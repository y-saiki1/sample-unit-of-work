package querymodel

type InvoiceQuery interface {
	FindByFilter(filter InvoiceQueryFilter) ([]InvoiceQueryModel, error)
}
