-- 1. Alterar as colunas usuario_id e animal_id para permitir NULL
ALTER TABLE
    adopt_pethub.feedbacks
ALTER COLUMN
    usuario_id DROP NOT NULL;

ALTER TABLE
    adopt_pethub.feedbacks
ALTER COLUMN
    animal_id DROP NOT NULL;

-- 2. Remover as restrições de chave estrangeira antigas
ALTER TABLE
    adopt_pethub.feedbacks DROP CONSTRAINT feedbacks_usuario_id_fkey;

ALTER TABLE
    adopt_pethub.feedbacks DROP CONSTRAINT feedbacks_animal_id_fkey;

-- 3. Adicionar as novas restrições de chave estrangeira com ON DELETE SET NULL
ALTER TABLE
    adopt_pethub.feedbacks
ADD
    CONSTRAINT feedbacks_usuario_id_fkey FOREIGN KEY (usuario_id) REFERENCES adopt_pethub.usuarios(id) ON DELETE
SET
    NULL;

ALTER TABLE
    adopt_pethub.feedbacks
ADD
    CONSTRAINT feedbacks_animal_id_fkey FOREIGN KEY (animal_id) REFERENCES adopt_pethub.animais(id) ON DELETE
SET
    NULL;