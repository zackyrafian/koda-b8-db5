CREATE TABLE "contact" ( 
  id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL, 
  email VARCHAR(80) UNIQUE NOT NULL, 
  phone VARCHAR(15) NOT NULL, 
  address VARCHAR(255) NOT NULL
);


INSERT INTO "contact" (name, email, phone, address) VALUES 
('Dua', 'dua@mail.com', '08999993758', 'Depok'),
('Tiga', 'tiga@mail.com', '08999993758', 'Depok'),
('Empat', 'empat@mail.com', '08999993758', 'Depok'),
('Lima', 'lima@mail.com', '08999993758', 'Depok'),
('Enam', 'enam@mail.com', '08999993758', 'Depok'),
('Tujuh', 'tujuh@mail.com', '08999993758', 'Depok'),
('Depalan', 'depalan@mail.com', '08999993758', 'Depok');

