package fluentsql

import (
	"fmt"
	"strings"
)

// Insert clause represents an SQL INSERT statement with a table name and columns.
// Cases:
// INSERT INTO Customers (CustomerName, ContactName, Address, City, PostalCode, Country)
// VALUES ('Cardinal', 'Tom B. Erichsen', 'Skagen 21', 'Stavanger', '4006', 'Norway');
//
// INSERT INTO dependents (first_name, last_name, relationship, employee_id)
// VALUES ('Cameron', 'Bell', 'Child', 192), ('Michelle', 'Bell', 'Child', 192);
//
// INSERT INTO Customers (CustomerName, ContactName, Address, City, PostalCode, Country)
// SELECT SupplierName, ContactName, Address, City, PostalCode, Country FROM Suppliers;
//
// INSERT INTO Customers (CustomerName, City, Country)
// SELECT SupplierName, City, Country FROM Suppliers WHERE Country='Germany';
type Insert struct {
	Table   string   // Table specifies the name of the table into which the data will be inserted.
	Columns []string // Columns defines the list of column names for the INSERT statement.
}

// String returns the SQL INSERT statement as a string.
// It joins the Columns slice with commas and formats it into the SQL syntax.
// Returns: A string representation of the SQL INSERT statement.
func (i *Insert) String() string {
	columnsStr := strings.Join(i.Columns, ", ") // Joins the column names with commas.

	return fmt.Sprintf("INSERT INTO %s (%s)", i.Table, columnsStr) // Formats the final SQL string.
}
