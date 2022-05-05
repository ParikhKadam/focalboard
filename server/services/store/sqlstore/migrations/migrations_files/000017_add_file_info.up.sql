{{if .plugin}}
    SELECT 1;
{{else}}
    CREATE TABLE IF NOT EXISTS {{.prefix}}file_info {
        create_at BIGINT NOT NULL,
        delete_at BIGINT,
        name TEXT NOT NULL,
        extention VARCHAR(50) NOT NULL,
        size BIGINT NOT NULL
        archived BOOEAN
    } {{if .mysql}}DEFAULT CHARACTER SET utf8mb4{{end}};
{{end}}
