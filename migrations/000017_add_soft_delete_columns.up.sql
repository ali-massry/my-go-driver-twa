-- Add soft delete columns to support soft deletes instead of hard deletes

-- Add deleted_at to drivers table
ALTER TABLE drivers ADD COLUMN deleted_at TIMESTAMP NULL;
CREATE INDEX idx_drivers_deleted_at ON drivers(deleted_at);

-- Add deleted_at to company_admins table
ALTER TABLE company_admins ADD COLUMN deleted_at TIMESTAMP NULL;
CREATE INDEX idx_company_admins_deleted_at ON company_admins(deleted_at);

-- Add deleted_at to company_modules table
ALTER TABLE company_modules ADD COLUMN deleted_at TIMESTAMP NULL;
CREATE INDEX idx_company_modules_deleted_at ON company_modules(deleted_at);

-- Note: companies table already has deleted_at column from earlier migration
