-- Criar o schema adopt_pethub
CREATE SCHEMA IF NOT EXISTS adopt_pethub;

-- Criar a tabela de usuários
CREATE TABLE IF NOT EXISTS adopt_pethub.usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    senha VARCHAR(255) NOT NULL,
    telefone VARCHAR(255),
    endereco TEXT,
    tipo_usuario VARCHAR(255),
    data_cadastro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Criar a tabela de animais
CREATE TABLE IF NOT EXISTS adopt_pethub.animais (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    especie VARCHAR(255) NOT NULL,
    raca VARCHAR(255),
    idade INTEGER,
    sexo VARCHAR(1),
    vacinado BOOLEAN,
    vermifugado BOOLEAN,
    castrado BOOLEAN,
    descricao TEXT NOT NULL,
    foto_url TEXT,
    status_adocao VARCHAR(255),
    data_resgate TIMESTAMP,
    data_cadastro TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Criar a tabela de adoções
CREATE TABLE IF NOT EXISTS adopt_pethub.adocoes (
    id SERIAL PRIMARY KEY,
    id_animal INTEGER NOT NULL REFERENCES adopt_pethub.animais(id),
    id_usuario INTEGER NOT NULL REFERENCES adopt_pethub.usuarios(id),
    data_adocao TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_adocao VARCHAR(255),
    observacoes TEXT,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Criar a tabela de abrigos
CREATE TABLE IF NOT EXISTS adopt_pethub.abrigos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    endereco VARCHAR(255),
    telefone VARCHAR(20),
    email VARCHAR(255),
    id_tutor INT REFERENCES adopt_pethub.usuarios(id),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Criar a tabela de feedbacks
CREATE TABLE IF NOT EXISTS adopt_pethub.feedbacks (
    id SERIAL PRIMARY KEY,
    usuario_id INT NOT NULL REFERENCES adopt_pethub.usuarios(id),
    animal_id INT NOT NULL REFERENCES adopt_pethub.animais(id),
    mensagem TEXT NOT NULL,
    data_avaliacao TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    avaliacao INTEGER,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);