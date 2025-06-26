

CREATE TABLE suggestions (
    id SERIAL PRIMARY KEY,
    collaborator_name VARCHAR(100) NOT NULL,
    sector VARCHAR(50),
    description TEXT NOT NULL,
    status VARCHAR(20) DEFAULT 'open' CHECK (status IN ('open', 'under review', 'implemented')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO suggestions (collaborator_name, sector, description, status)
VALUES
('Alice Johnson', 'Logistics', 'Implement a barcode system.', 'open'),
('Bruno Silva', 'IT', 'Create an internal knowledge base.', 'under review'),
('Carla Mendes', 'Customer Service', 'Add a chat option to the website.', 'implemented');



