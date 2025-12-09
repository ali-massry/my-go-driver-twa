-- Add extended company fields

-- Basic Company Info (additional)
ALTER TABLE companies ADD COLUMN legal_name VARCHAR(255) AFTER name;
ALTER TABLE companies ADD COLUMN whatsapp VARCHAR(50) AFTER phone;
ALTER TABLE companies ADD COLUMN country VARCHAR(100) AFTER address;

-- Branding & UI
ALTER TABLE companies ADD COLUMN theme ENUM('light', 'dark', 'custom') DEFAULT 'light' AFTER font_family;
ALTER TABLE companies ADD COLUMN custom_css TEXT AFTER theme;

-- Localization
ALTER TABLE companies ADD COLUMN locale VARCHAR(10) DEFAULT 'en' AFTER timezone;
ALTER TABLE companies ADD COLUMN date_format VARCHAR(50) DEFAULT 'dd/mm/yyyy' AFTER locale;
ALTER TABLE companies ADD COLUMN currency VARCHAR(10) DEFAULT 'USD' AFTER date_format;

-- Business Rules (JSON fields)
ALTER TABLE companies ADD COLUMN auto_assign_rules JSON AFTER currency;
ALTER TABLE companies ADD COLUMN delivery_pricing_rules JSON AFTER auto_assign_rules;
ALTER TABLE companies ADD COLUMN cash_handling_rules JSON AFTER delivery_pricing_rules;
ALTER TABLE companies ADD COLUMN pod_required BOOLEAN DEFAULT FALSE AFTER cash_handling_rules;
ALTER TABLE companies ADD COLUMN driver_shift_rules JSON AFTER pod_required;
ALTER TABLE companies ADD COLUMN vehicle_assignment_mode ENUM('auto', 'manual') DEFAULT 'manual' AFTER driver_shift_rules;
ALTER TABLE companies ADD COLUMN max_extra_delivery_qty INT DEFAULT 0 AFTER vehicle_assignment_mode;

-- Logistics Settings
ALTER TABLE companies ADD COLUMN routing_mode ENUM('simple', 'optimized', 'AI') DEFAULT 'simple' AFTER max_extra_delivery_qty;
ALTER TABLE companies ADD COLUMN gps_accuracy ENUM('low', 'medium', 'high') DEFAULT 'medium' AFTER routing_mode;
ALTER TABLE companies ADD COLUMN depot_id BIGINT UNSIGNED AFTER gps_accuracy;
ALTER TABLE companies ADD COLUMN has_multiple_stores BOOLEAN DEFAULT FALSE AFTER depot_id;

-- Inventory Settings
ALTER TABLE companies ADD COLUMN enable_vehicle_stock BOOLEAN DEFAULT FALSE AFTER has_multiple_stores;
ALTER TABLE companies ADD COLUMN enable_product_catalog BOOLEAN DEFAULT FALSE AFTER enable_vehicle_stock;
ALTER TABLE companies ADD COLUMN allowed_products JSON AFTER enable_product_catalog;

-- Notifications
ALTER TABLE companies ADD COLUMN notification_settings JSON AFTER allowed_products;
ALTER TABLE companies ADD COLUMN broadcast_enabled BOOLEAN DEFAULT TRUE AFTER notification_settings;

-- Driver Limit
ALTER TABLE companies ADD COLUMN max_allowed_drivers INT DEFAULT 10 AFTER broadcast_enabled;

-- Billing & Subscription
ALTER TABLE companies ADD COLUMN plan ENUM('free', 'basic', 'pro', 'enterprise') DEFAULT 'free' AFTER max_allowed_drivers;
ALTER TABLE companies ADD COLUMN billing_cycle ENUM('monthly', 'yearly') DEFAULT 'monthly' AFTER plan;
ALTER TABLE companies ADD COLUMN seats_limit INT DEFAULT 10 AFTER billing_cycle;
ALTER TABLE companies ADD COLUMN api_rate_limit INT DEFAULT 1000 AFTER seats_limit;

-- Audit (soft delete)
ALTER TABLE companies ADD COLUMN deleted_at TIMESTAMP NULL AFTER updated_at;

-- Add foreign key for depot_id if stores table exists
-- ALTER TABLE companies ADD CONSTRAINT fk_companies_depot FOREIGN KEY (depot_id) REFERENCES stores(id) ON DELETE SET NULL;

-- Add indexes for commonly queried fields
CREATE INDEX idx_companies_country ON companies(country);
CREATE INDEX idx_companies_plan ON companies(plan);
CREATE INDEX idx_companies_deleted_at ON companies(deleted_at);
