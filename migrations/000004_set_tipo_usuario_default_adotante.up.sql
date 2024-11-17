ALTER TABLE
    adopt_pethub.usuarios
ALTER COLUMN
    tipo_usuario
SET
    DEFAULT 'ADOTANTE';

ALTER TABLE
    adopt_pethub.animais
ALTER COLUMN
    status_adocao
SET
    DEFAULT 'DISPONÍVEL';

    ALTER TABLE
    adopt_pethub.animais
ALTER COLUMN
    sexo TYPE character varying(10);