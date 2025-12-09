-- Populate modules_master table with all available modules (using INSERT IGNORE to skip duplicates)

-- Core Modules
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('store_management', 'Store Management', 'Core', 'Manage multiple stores/branches and their operations', TRUE),
('driver_management', 'Driver Management', 'Core', 'Manage drivers, assignments, and schedules', TRUE),
('vehicle_management', 'Vehicle Management', 'Core', 'Manage company vehicles and maintenance', TRUE),
('order_management', 'Order Management', 'Core', 'Create, track, and manage orders', TRUE),
('client_management', 'Client Management', 'Core', 'Manage client information and relationships', TRUE),
('product_inventory', 'Product Inventory', 'Core', 'Track products and inventory levels', FALSE),
('proof_of_delivery', 'Proof of Delivery (POD)', 'Core', 'Capture delivery confirmations', TRUE),
('cash_handling', 'Cash Handling', 'Core', 'Manage cash collections and reconciliation', FALSE),
('task_scheduling', 'Task Scheduling', 'Core', 'Schedule tasks and deliveries', TRUE),
('shift_management', 'Shift Management', 'Core', 'Manage driver shifts and work hours', TRUE),
('gps_tracking', 'GPS Tracking', 'Core', 'Real-time driver location tracking', TRUE),
('chat_communication', 'Chat / Communication', 'Core', 'Internal messaging system', FALSE),
('notifications', 'Notifications & Announcements', 'Core', 'Push notifications and broadcasts', TRUE),
('reports_analytics', 'Reports & Analytics', 'Core', 'Business intelligence and reporting', TRUE);

-- Advanced Logistics Modules
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('route_optimization', 'Route Optimization', 'Advanced Logistics', 'Optimize delivery routes for efficiency', FALSE),
('zone_territory', 'Zone & Territory Management', 'Advanced Logistics', 'Define and manage delivery zones', FALSE),
('auto_assignment', 'Auto-Assignment Engine', 'Advanced Logistics', 'Automatically assign orders to drivers', FALSE),
('multi_depot', 'Multi-Depot Support', 'Advanced Logistics', 'Manage multiple distribution centers', FALSE),
('vehicle_load_planning', 'Vehicle Load Planning', 'Advanced Logistics', 'Optimize vehicle loading', FALSE),
('driver_efficiency', 'Driver Efficiency Scoring', 'Advanced Logistics', 'Track and score driver performance', FALSE),
('delivery_sla', 'Delivery SLA Tracking', 'Advanced Logistics', 'Monitor delivery time commitments', FALSE),
('incident_reporting', 'Incident Reporting', 'Advanced Logistics', 'Report and track delivery incidents', FALSE);

-- Delivery & POD Modules
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('signature_pod', 'Signature POD', 'Delivery & POD', 'Capture customer signatures', TRUE),
('photo_pod', 'Photo POD', 'Delivery & POD', 'Photo-based delivery confirmation', TRUE),
('otp_qr_delivery', 'OTP/QR Delivery', 'Delivery & POD', 'OTP or QR code verification', FALSE),
('document_scan_pod', 'Document Scan POD', 'Delivery & POD', 'Scan delivery documents', FALSE),
('task_checklist', 'Task Checklist', 'Delivery & POD', 'Step-by-step task verification', FALSE),
('return_to_depot', 'Return-to-Depot Workflow', 'Delivery & POD', 'Handle returns and undelivered items', FALSE);

-- Inventory Modules
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('vehicle_stock', 'Vehicle Stock Module', 'Inventory', 'Track inventory in vehicles', FALSE),
('warehouse_stock', 'Warehouse Stock Module', 'Inventory', 'Warehouse inventory management', FALSE),
('stock_movement', 'Stock Movement Logs', 'Inventory', 'Track all stock movements', FALSE),
('realtime_driver_inventory', 'Real-Time Driver Inventory', 'Inventory', 'Live driver stock updates', FALSE),
('eod_stock_return', 'End-of-Day Stock Return', 'Inventory', 'Daily stock reconciliation', FALSE),
('auto_reorder', 'Auto-Reorder Alerts', 'Inventory', 'Automated reorder notifications', FALSE);

-- Payments & Finance
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('cash_collection', 'Cash Collection Tracking', 'Payments & Finance', 'Track cash collected by drivers', FALSE),
('online_payment', 'Online Payment Tracking', 'Payments & Finance', 'Track digital payments', FALSE),
('cash_discrepancy', 'Cash Discrepancy Alerts', 'Payments & Finance', 'Alert on cash mismatches', FALSE),
('driver_reconciliation', 'Driver Cash Reconciliation', 'Payments & Finance', 'End-of-day driver settlements', FALSE),
('client_billing', 'Client Billing Module', 'Payments & Finance', 'Generate client invoices', FALSE),
('invoice_generation', 'Invoice Generation', 'Payments & Finance', 'Automated invoice creation', FALSE);

-- Communication & Support
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('driver_admin_chat', 'Driver and Admin Chat', 'Communication', 'Direct messaging between drivers and admin', FALSE),
('broadcast_messaging', 'Broadcast Messaging', 'Communication', 'Send announcements to all drivers', TRUE),
('sos_emergency', 'SOS / Emergency Alerts', 'Communication', 'Emergency notification system', FALSE),
('issue_reporting', 'Issue Reporting Module', 'Communication', 'Report and track issues', FALSE);

-- Integration Modules
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('api_access', 'API Access', 'Integration', 'RESTful API for external integrations', FALSE),
('webhooks', 'Webhooks', 'Integration', 'Event-driven notifications', FALSE),
('third_party_delivery', 'Third-party Delivery Integrations', 'Integration', 'Connect with delivery services', FALSE),
('accounting_integration', 'Accounting Integrations', 'Integration', 'Xero, QuickBooks integration', FALSE),
('erp_integration', 'ERP Integrations', 'Integration', 'Connect with ERP systems', FALSE),
('crm_integration', 'CRM Integrations', 'Integration', 'Connect with CRM systems', FALSE);

-- Additional SaaS Modules
INSERT IGNORE INTO modules_master (module_key, name, category, description, default_enabled) VALUES
('custom_branding', 'Custom Branding Package', 'SaaS Features', 'Complete brand customization', FALSE),
('whitelabel_app', 'White-Label Mobile App', 'SaaS Features', 'Branded mobile applications', FALSE),
('roles_permissions', 'Roles & Permissions Customization', 'SaaS Features', 'Advanced access control', FALSE),
('audit_logs', 'Audit Logs Module', 'SaaS Features', 'Complete activity logging', FALSE),
('ai_driver_assignment', 'AI Driver Assignment', 'SaaS Features', 'AI-powered driver matching', FALSE),
('ai_route_planning', 'AI Route Planning', 'SaaS Features', 'AI-optimized routing', FALSE);
