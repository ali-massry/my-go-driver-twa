CREATE TABLE IF NOT EXISTS vehicles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    company_id BIGINT UNSIGNED NOT NULL,
    store_id BIGINT UNSIGNED,
    plate_number VARCHAR(50) NOT NULL,
    type ENUM('bike', 'car', 'van', 'truck') NOT NULL,
    capacity DECIMAL(10, 2),
    status ENUM('active', 'maintenance', 'out_of_service') DEFAULT 'active',
    fuel_type VARCHAR(50),
    last_oil_change DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
    FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE SET NULL,
    UNIQUE KEY unique_plate_company (plate_number, company_id),
    INDEX idx_vehicles_company (company_id),
    INDEX idx_vehicles_store (store_id),
    INDEX idx_vehicles_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
