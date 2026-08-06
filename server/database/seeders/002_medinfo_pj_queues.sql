DELETE FROM medinfo_pj_queues;

INSERT INTO medinfo_pj_queues (
    user_id,
    position,
    is_current
)
SELECT
    users.id,
    CASE users.email
        WHEN 'caca.medinfo@bem.unair.ac.id' THEN 1
        WHEN 'ocha.medinfo@bem.unair.ac.id' THEN 2
        WHEN 'febi.medinfo@bem.unair.ac.id' THEN 3
    END,
    CASE
        WHEN users.email = 'caca.medinfo@bem.unair.ac.id'
            THEN TRUE
        ELSE FALSE
    END
FROM users
WHERE users.email IN (
    'caca.medinfo@bem.unair.ac.id',
    'ocha.medinfo@bem.unair.ac.id',
    'febi.medinfo@bem.unair.ac.id'
)
ORDER BY
    CASE users.email
        WHEN 'caca.medinfo@bem.unair.ac.id' THEN 1
        WHEN 'ocha.medinfo@bem.unair.ac.id' THEN 2
        WHEN 'febi.medinfo@bem.unair.ac.id' THEN 3
    END
ON DUPLICATE KEY UPDATE
    position = VALUES(position),
    is_current = VALUES(is_current),
    updated_at = NOW();

SELECT
    queue.id,
    queue.position,
    queue.is_current,
    users.id AS user_id,
    users.name,
    users.email,
    users.role,
    users.ministry
FROM medinfo_pj_queues AS queue
INNER JOIN users
    ON users.id = queue.user_id
ORDER BY queue.position ASC;