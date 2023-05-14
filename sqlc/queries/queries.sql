-- name: GetTransaction :one
SELECT *
FROM desafio_dev.public.transactions
WHERE market = $1;

-- name: ListMarketTransaction :many
SELECT *
from desafio_dev.public.transactions
WHERE market = $1;

-- name: MarketBalance :one
SELECT CAST(sum(
        CASE
            WHEN type IN ('1', '4', '5', '6', '7', '8') THEN value
            WHEN type IN ('2', '3', '9') THEN -value
            END
    ) AS FLOAT) AS total_balance
FROM desafio_dev.public.transactions
WHERE market = $1;

-- name: ListMarkets :many
SELECT market
FROM desafio_dev.public.transactions
GROUP BY market;

-- name: CreateTransaction :one
INSERT INTO desafio_dev.public.transactions (type, date, value, cpf, card, time, owner, market)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

