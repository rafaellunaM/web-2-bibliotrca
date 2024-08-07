# Listar Empréstimos com Informações do Cliente e Livro:
    SELECT
        loan.loanId,
        client.name AS client_name,
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
