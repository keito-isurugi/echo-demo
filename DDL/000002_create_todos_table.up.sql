CREATE TABLE IF NOT EXISTS todos
(
    id         SERIAL PRIMARY KEY NOT NULL,
    title      VARCHAR(255)        NOT NULL,
    content    VARCHAR(255)        NOT NULL,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE todos IS 'Todosテーブル';
COMMENT ON COLUMN todos.id IS 'ID';
COMMENT ON COLUMN todos.title IS 'タイトル';
COMMENT ON COLUMN todos.content IS '内容';
COMMENT ON COLUMN todos.created_at IS '登録日時';
COMMENT ON COLUMN todos.updated_at IS '更新日時';