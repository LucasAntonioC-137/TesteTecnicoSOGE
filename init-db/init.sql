

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

('Lucas Ferreira', 'Financeiro', 'Criar relatórios mensais automáticos.', 'open'),
('Thiago Costa', 'Suporte Técnico', 'Criar manual digital para atendimento ao cliente.', 'under review'),
('Camila Ribeiro', 'Financeiro', 'Automatizar o envio de boletos aos clientes.', 'implemented'),
('Sérgio Martins', 'TI', 'Implementar backup automático no servidor.', 'implemented'),
('Juliana Rocha', 'RH', 'Implementar programa de bem-estar para colaboradores.', 'open'),
('Marcos Vinícius', 'Operações', 'Otimizar a rota de entrega dos produtos.', 'under review'),
('João Pedro', 'Comercial', 'Adicionar assinatura digital nos contratos.', 'implemented'),
('Fernanda Lima', 'Marketing', 'Sugerir parceria com influenciadores locais.', 'open'),
('Patrícia Almeida', 'Logística', 'Aumentar o espaço do estoque com prateleiras modulares.', 'under review'),
('Diego Santos', 'TI', 'Melhorar a autenticação no sistema interno.', 'open'),
('Amanda Lopes', 'RH', 'Adotar sistema de ponto eletrônico via app.', 'implemented');




