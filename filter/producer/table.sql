SELECT 
    TABLE_NAME,
    TABLE_COMMENT
FROM 
    TABLES
WHERE 
    TABLE_SCHEMA = ?;