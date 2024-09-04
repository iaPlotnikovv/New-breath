CREATE TABLE Banners
(
    BannerId serial PRIMARY KEY,
    FeatureId INTEGER,
    TagIds INTEGER ARRAY,
    Content JSONB, 
    IsActive BOOL,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    UNIQUE (FeatureId,TagIds)
);