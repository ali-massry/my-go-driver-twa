-- Rollback: Remove soft delete columns

-- Remove from company_modules table
DROP INDEX IF EXISTS idx_company_modules_deleted_at ON company_modules;
ALTER TABLE company_modules DROP COLUMN deleted_at;

-- Remove from company_admins table
DROP INDEX IF EXISTS idx_company_admins_deleted_at ON company_admins;
ALTER TABLE company_admins DROP COLUMN deleted_at;

-- Remove from drivers table
DROP INDEX IF EXISTS idx_drivers_deleted_at ON drivers;
ALTER TABLE drivers DROP COLUMN deleted_at;
