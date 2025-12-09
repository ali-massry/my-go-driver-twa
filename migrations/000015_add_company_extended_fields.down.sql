-- Remove indexes
DROP INDEX IF EXISTS idx_companies_deleted_at ON companies;
DROP INDEX IF EXISTS idx_companies_plan ON companies;
DROP INDEX IF EXISTS idx_companies_country ON companies;

-- Remove foreign key
-- ALTER TABLE companies DROP FOREIGN KEY IF EXISTS fk_companies_depot;

-- Remove all added columns
ALTER TABLE companies DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE companies DROP COLUMN IF EXISTS api_rate_limit;
ALTER TABLE companies DROP COLUMN IF EXISTS seats_limit;
ALTER TABLE companies DROP COLUMN IF EXISTS billing_cycle;
ALTER TABLE companies DROP COLUMN IF EXISTS plan;
ALTER TABLE companies DROP COLUMN IF EXISTS max_allowed_drivers;
ALTER TABLE companies DROP COLUMN IF EXISTS broadcast_enabled;
ALTER TABLE companies DROP COLUMN IF EXISTS notification_settings;
ALTER TABLE companies DROP COLUMN IF EXISTS allowed_products;
ALTER TABLE companies DROP COLUMN IF EXISTS enable_product_catalog;
ALTER TABLE companies DROP COLUMN IF EXISTS enable_vehicle_stock;
ALTER TABLE companies DROP COLUMN IF EXISTS has_multiple_stores;
ALTER TABLE companies DROP COLUMN IF EXISTS depot_id;
ALTER TABLE companies DROP COLUMN IF EXISTS gps_accuracy;
ALTER TABLE companies DROP COLUMN IF EXISTS routing_mode;
ALTER TABLE companies DROP COLUMN IF EXISTS max_extra_delivery_qty;
ALTER TABLE companies DROP COLUMN IF EXISTS vehicle_assignment_mode;
ALTER TABLE companies DROP COLUMN IF EXISTS driver_shift_rules;
ALTER TABLE companies DROP COLUMN IF EXISTS pod_required;
ALTER TABLE companies DROP COLUMN IF EXISTS cash_handling_rules;
ALTER TABLE companies DROP COLUMN IF EXISTS delivery_pricing_rules;
ALTER TABLE companies DROP COLUMN IF EXISTS auto_assign_rules;
ALTER TABLE companies DROP COLUMN IF EXISTS currency;
ALTER TABLE companies DROP COLUMN IF EXISTS date_format;
ALTER TABLE companies DROP COLUMN IF EXISTS locale;
ALTER TABLE companies DROP COLUMN IF EXISTS custom_css;
ALTER TABLE companies DROP COLUMN IF EXISTS theme;
ALTER TABLE companies DROP COLUMN IF EXISTS country;
ALTER TABLE companies DROP COLUMN IF EXISTS whatsapp;
ALTER TABLE companies DROP COLUMN IF EXISTS legal_name;
