{{if .plugin}}
    SELECT 1;
{{else}}
    DROP TABLE {{.prefix}}file_info;
{{end}}
