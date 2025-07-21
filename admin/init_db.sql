CREATE TABLE IF NOT EXISTS stock_recommendations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticker STRING NOT NULL,
    company STRING NOT NULL,
    brokerage STRING NOT NULL,
    action STRING NOT NULL,
    target_from DECIMAL(12,2),
    target_to DECIMAL(12,2),
    rating_from STRING,
    rating_to STRING,
    time TIMESTAMP NOT NULL,

    -- Nuevos campos para la recomendación calculada
    recommendation STRING NOT NULL DEFAULT 'HOLD',   -- BUY/SELL/HOLD
    confidence FLOAT8 NOT NULL DEFAULT 0.5,          -- 0.0 - 1.0
    score INT NOT NULL DEFAULT 0,                    -- Score numérico
    reason STRING NOT NULL DEFAULT ''                -- Explicación textual
);

CREATE INDEX IF NOT EXISTS recommendations_ticker_time_idx ON stock_recommendations (ticker, time DESC);
CREATE INDEX IF NOT EXISTS recommendations_recommendation_idx ON stock_recommendations (recommendation);
CREATE INDEX IF NOT EXISTS recommendations_brokerage_idx ON stock_recommendations (brokerage);
