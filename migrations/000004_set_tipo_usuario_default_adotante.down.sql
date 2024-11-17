ALTER TABLE
    adopt_pethub.usuarios
ALTER COLUMN
    tipo_usuario DROP DEFAULT;

ALTER TABLE
    adopt_pethub.animais
ALTER COLUMN
    status_adocao DROP DEFAULT;