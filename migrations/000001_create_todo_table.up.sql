CREATE TABLE todos (
  id UUID NOT NULL PRIMARY KEY,
  name varchar(100) NOT NULL,
  completed BOOLEAN NOT NULL
);

INSERT INTO todos (id, name, completed)
VALUES 
	('8bf94cd1-98db-44dd-b82a-ce7af028b677', 'teste1', FALSE),
	('3d2d1916-e11c-4b87-a514-97b082632b0b', 'teste2', TRUE);
