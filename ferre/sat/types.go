package types

struct Customer {
    Name string 	`json:"legal_name"`
    RFC string		`json:"tax_id"`
    Email string	`json:"email"`
    Address Address	`json:"address"`
}

struct Address {
    Street string	`json:"street"`
    Exterior string	`json:"exterior"`
    Colonia string `json:"neighborhood"`
    ZIP string		`json:"zip"`
    Interior string	`json:"interior"`
}


struct Product {
    Description string	`json:"description"`
    SATCode string	`json:"product_key"`
    Precio float64	`json:"price"`
    Unidad string	`json:"unit_key"`
    SKU string		`json:"sku"`
}

struct Invoice {
    Cliente Customer 	`json:"customer"`
    Items []Item	`json:"items"`
    FormaPago string	`json:"payment_form"`
    Folio int32		`json:"folio_number"`
    Serie string	`json:"series"` // 1-25 caracteres
    QRCode string	`json:"external_id"`
}

