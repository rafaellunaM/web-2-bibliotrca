# Listar Empréstimos com Informações do Cliente e Livro:
    SELECT
        loan.loanId,
        loan.cpf,
        client.name AS client_name,
        client.name,
        book.title AS book_title,
        loan.dueDate,
        loan.returnDate
    FROM loan
    JOIN client ON loan.cpf = client.cpf
    JOIN book ON loan.bookId = book.bookId;

# Listar Multas com Informações do Cliente e Empréstimo:
    SELECT
        fine.fineId,
        client.name AS client_name,
        loan.loanId,
        fine.amount,
        fine.paid
    FROM fine
    JOIN client ON fine.cpf = client.cpf
    JOIN loan ON fine.loanId = loan.loanId;

# Listar Livros Disponíveis:
    SELECT title FROM book WHERE availability = TRUE;

# Buscar Empréstimos Pendentes:
    SELECT * FROM loan WHERE returnDate IS NULL;

# Listar Multas Não Pagas:
    SELECT * FROM fine WHERE paid = TRUE;

# Contar o Número de Bibliotecários:
    SELECT COUNT(*) FROM librarian;

# Contar o Número de Livros Disponíveis:
    SELECT COUNT(*) FROM book WHERE availability = TRUE;

# Contar o Número de Empréstimos Pendentes:
    SELECT COUNT(*) FROM loan WHERE returnDate IS NULL;



# QUERY EXAMPLE USING GORM

    Librarian == Admin == Client == Book
	err = db.AutoMigrate(&tables.Book{})
	checErr(err)

	var book tables.Book
	result := db.First(&book)
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
	} else {
		fmt.Printf("Book: %+v\n", book)
	}

