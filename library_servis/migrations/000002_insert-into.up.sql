INSERT INTO users (id, username, email, password, created_at, updated_at, deleted_at) VALUES
('1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed', 'Ali', 'ali@example.com', 'password1', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('2c153ade-8212-4e01-ba31-3b2a3f57309f', 'Valijon', 'valijon@example.com', 'password2', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('31b5424b-01eb-4925-a515-87e2a1e4c2f7', 'Shirin', 'shirin@example.com', 'password3', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('4d5e7a4f-759e-4bcf-8a79-b66e4508a62c', 'Otabek', 'otabek@example.com', 'password4', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('5e8d6a73-188f-41b8-9b92-ff34c8a1e1b2', 'Madina', 'madina@example.com', 'password5', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('6f2a8f7c-d43d-4394-92b2-98060fddbb8b', 'Zilola', 'zilola@example.com', 'password6', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('72a2c2d8-1b69-4c2e-916d-0d57b59733a6', 'Sanobar', 'sanobar@example.com', 'password7', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('87b4fd47-c2f1-49e0-aa6c-8b9c20a54fa8', 'Javohir', 'javohir@example.com', 'password8', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('9b6e4eef-4f19-4b06-8c9b-8a6f0d2f6723', 'Shaxzoda', 'shaxzoda@example.com', 'password9', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('a1f20276-79f3-4b66-8fe6-9f0b6754e7de', 'Bobur', 'bobur@example.com', 'password10', '2024-06-19 19:48:50.680764', '2024-06-19 19:48:50.680764', 0),
('4ebcb8f6-52ea-404c-bb98-c6d47c288d65', 'testuser', 'testuser@example.com', 'password123', '2024-06-20 12:10:50.941412', '2024-06-20 12:10:50.941412', 0);



-- Authors jadvaliga ma'lumotlarni kiritish
INSERT INTO authors (id, name, biography)
VALUES
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Abdulla Qodiriy', 'Abdulla Qodiriy o’zbek yozuvchisi'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Cholpon', 'Cholpon (Abdurauf Fitrat) o’zbek shoir va yozuvchisi'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Hamid Olimjon', 'Hamid Olimjon o’zbek adabiyotining taniqli vakili'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'O’tkir Hoshimov', 'O’tkir Hoshimov o’zbek yozuvchisi'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Asqad Muxtor', 'Asqad Muxtor o’zbek yozuvchisi va dramaturgi'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 'Erkin Vohidov', 'Erkin Vohidov o’zbek shoir va yozuvchisi'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 'Muhammad Yusuf', 'Muhammad Yusuf o’zbek shoiri'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 'Gafur G’ulom', 'Gafur G’ulom o’zbek shoir va yozuvchisi'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 'Said Ahmad', 'Said Ahmad o’zbek yozuvchisi va dramaturgi'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a1a', 'Abdulla Oripov', 'Abdulla Oripov o’zbek shoiri');


-- Janrlarni Genres jadvaliga kiriting
INSERT INTO genres (id, name)
VALUES
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c11', 'Fantastika'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c12', 'Detektiv'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c13', 'Romantika'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c14', 'Ilmiy-fantastika'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c15', 'Tarixiy'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c16', 'Dramaturgiya'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c17', 'Sarguzasht'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c18', 'Triller'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c19', 'She’riyat'),
('b1eebc99-9c0b-4ef8-bb6d-6bb9bd380c1a', 'Badiiy adabiyot');


INSERT INTO Books (id, title, author_id, genre_id, summary)
VALUES
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c51', 'O’tgan kunlar', 
    (SELECT id FROM Authors WHERE name = 'Abdulla Qodiriy'), 
    (SELECT id FROM Genres WHERE name = 'Tarixiy'), 
    'Bu roman XIX asr oxiri va XX asr boshidagi Turkistonning ijtimoiy va siyosiy hayotini tasvirlaydi.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c52', 'Kecha va kunduz', 
    (SELECT id FROM Authors WHERE name = 'Cholpon'), 
    (SELECT id FROM Genres WHERE name = 'Romantika'), 
    'Cholponning bu romani o’zbek adabiyotida romantik yo’nalishdagi muhim asar.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c53', 'Parvoz', 
    (SELECT id FROM Authors WHERE name = 'Hamid Olimjon'), 
    (SELECT id FROM Genres WHERE name = 'Dramaturgiya'), 
    'Hamid Olimjonning bu asari o’zbek adabiyotida dramaturgik yuksalishlardan biri sifatida tan olingan.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c54', 'Dunyoning ishlari', 
    (SELECT id FROM Authors WHERE name = 'O’tkir Hoshimov'), 
    (SELECT id FROM Genres WHERE name = 'Badiiy adabiyot'), 
    'O’tkir Hoshimovning bu romani insoniy munosabatlar va jamiyat muammolari haqida hikoya qiladi.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c55', 'Bahor keldi', 
    (SELECT id FROM Authors WHERE name = 'Asqad Muxtor'), 
    (SELECT id FROM Genres WHERE name = 'Sarguzasht'), 
    'Asqad Muxtorning bu asari inson va tabiat o’rtasidagi munosabatlar haqida.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c56', 'Uzbekistan', 
    (SELECT id FROM Authors WHERE name = 'Erkin Vohidov'), 
    (SELECT id FROM Genres WHERE name = 'She’riyat'), 
    'Erkin Vohidovning bu she’riy asari Vatan sevgisi va uning madaniyati haqida.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c57', 'Muhabbatnoma', 
    (SELECT id FROM Authors WHERE name = 'Muhammad Yusuf'), 
    (SELECT id FROM Genres WHERE name = 'She’riyat'), 
    'Muhammad Yusufning bu to’plami sevgi va his-tuyg’ularni ifodalaydi.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c58', 'Shum bola', 
    (SELECT id FROM Authors WHERE name = 'Gafur G’ulom'), 
    (SELECT id FROM Genres WHERE name = 'Sarguzasht'), 
    'Gafur G’ulomning bu asari yosh bolalar va ularning sarguzashtlari haqida.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c59', 'Kelinlar qo’zg’oloni', 
    (SELECT id FROM Authors WHERE name = 'Said Ahmad'), 
    (SELECT id FROM Genres WHERE name = 'Dramaturgiya'), 
    'Said Ahmadning bu p`yesasi kelinlarning oiladagi o’rni haqida.'),
    
('c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c65', 'Ozbekiston', 
    (SELECT id FROM Authors WHERE name = 'Abdulla Oripov'), 
    (SELECT id FROM Genres WHERE name = 'She’riyat'), 
    'Abdulla Oripovning bu she’ri Vatan haqida buyuk sevgi va faxrni ifodalaydi.');



-- Borrowers jadvaliga ma'lumotlarni kiritish
INSERT INTO Borrowers (id, user_id, book_id, borrow_date, return_date)
VALUES
('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380d01', 
    '1b9d6bcd-bbfd-4b2d-9b5d-ab8dfbbd4bed', 
    'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c65', 
    '2024-07-01 10:00:00', 
    '2024-07-15 10:00:00'),
    
('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380d02', 
    '2c153ade-8212-4e01-ba31-3b2a3f57309f', 
    'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c59', 
    '2024-07-02 11:00:00', 
    '2024-07-16 11:00:00'),
    
('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380d03', 
    '31b5424b-01eb-4925-a515-87e2a1e4c2f7', 
    'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c58', 
    '2024-07-03 12:00:00', 
    '2024-07-17 12:00:00'),
    
('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380d04', 
    '4d5e7a4f-759e-4bcf-8a79-b66e4508a62c', 
    'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c57', 
    '2024-07-04 13:00:00', 
    '2024-07-18 13:00:00'),
    
('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380d05', 
    '5e8d6a73-188f-41b8-9b92-ff34c8a1e1b2', 
    'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c56', 
    '2024-07-05 14:00:00', 
    '2024-07-19 14:00:00'),
    
('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380d06', 
    '6f2a8f7c-d43d-4394-92b2-98060fddbb8b', 
    'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c55', 
    '2024-07-06 15:00:00', 
    '2024-07-20 15:00:00'),
    
('d1eebc99-9c0b-4ef8-bb6d-6bb9bd380d07', 
    '87b4fd47-c2f1-49e0-aa6c-8b9c20a54fa8', 
    'c1eebc99-9c0b-4ef8-bb6d-6bb9bd380c54', 
    '2024-07-07 16:00:00', 
    '2024-07-21 16:00:00'
);