-- Create database if it doesn't exist
SELECT 'CREATE DATABASE hermes'  
WHERE NOT EXISTS (
    SELECT FROM pg_database WHERE datname = 'hermes'
)\gexec; 