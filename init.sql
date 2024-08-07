CREATE TABLE Librarian (
    cpf varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255),
    password varchar(255),
    phoneNumber varchar(255)
);

CREATE TABLE Admin (
    cpf varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255),
    password varchar(255),
    phoneNumber varchar(255)
);

CREATE TABLE Client (
    cpf varchar(255) PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255),
    password varchar(255),
    phoneNumber varchar(255)
);

CREATE TABLE Book (
    bookId varchar(255) PRIMARY KEY,
    title varchar(255) NOT NULL,
    author varchar(255),
    category varchar(255),
    availability boolean DEFAULT FALSE
);

CREATE TABLE Loan (
    loanId varchar(255) PRIMARY KEY,
    cpf varchar(255),
    bookId varchar(255),
    dueDate DATE NOT NULL,
    returnDate DATE,
    FOREIGN KEY (cpf) REFERENCES Client(cpf),
    FOREIGN KEY (bookId) REFERENCES Book(bookId)
);

CREATE TABLE Fine (
    fineId varchar(255) PRIMARY KEY,
    cpf varchar(255),
    loanId varchar(255),
    amount float,
    paid boolean DEFAULT FALSE,
    FOREIGN KEY (cpf) REFERENCES Client(cpf),
    FOREIGN KEY (loanId) REFERENCES Loan(loanId)
);


INSERT INTO Librarian (cpf, name, email, password, phoneNumber) VALUES
('12345678901', 'João Silva', 'joao.silva@example.com', 'senha123', '11987654321'),
('23456789012', 'Maria Oliveira', 'maria.oliveira@example.com', 'senha456', '11976543210');

INSERT INTO Admin (cpf, name, email, password, phoneNumber) VALUES
('34567890123', 'Carlos Souza', 'carlos.souza@example.com', 'senha789', '11965432109'),
('45678901234', 'Ana Costa', 'ana.costa@example.com', 'senha101', '11954321098');

INSERT INTO Client (cpf, name, email, password, phoneNumber) VALUES
('56789012345', 'Fernanda Lima', 'fernanda.lima@example.com', 'senha202', '11943210987'),
('67890123456', 'Pedro Santos', 'pedro.santos@example.com', 'senha303', '11932109876');

INSERT INTO Book (bookId, title, author, category, availability) VALUES
('B001', 'O Senhor dos Anéis', 'J.R.R. Tolkien', 'Fantasia', TRUE),
('B002', '1984', 'George Orwell', 'Distopia', TRUE),
('B003', 'O Pequeno Príncipe', 'Antoine de Saint-Exupéry', 'Infantil', FALSE);

INSERT INTO Loan (loanId, cpf, bookId, dueDate, returnDate) VALUES
('L001', '56789012345', 'B001', '2024-09-01', NULL),
('L002', '67890123456', 'B002', '2024-09-15', '2024-09-10');

INSERT INTO Fine (fineId, cpf, loanId, amount, paid) VALUES
('F001', '67890123456', 'L002', 15.00, TRUE),
('F002', '56789012345', 'L001', 0.00, FALSE);
