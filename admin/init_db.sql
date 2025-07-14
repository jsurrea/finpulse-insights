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
    time TIMESTAMP NOT NULL
);
