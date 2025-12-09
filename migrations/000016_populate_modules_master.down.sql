-- Remove all inserted modules
DELETE FROM modules_master WHERE module_key IN (
    -- Core Modules
    'store_management', 'driver_management', 'vehicle_management', 'order_management',
    'client_management', 'product_inventory', 'proof_of_delivery', 'cash_handling',
    'task_scheduling', 'shift_management', 'gps_tracking', 'chat_communication',
    'notifications', 'reports_analytics',

    -- Advanced Logistics
    'route_optimization', 'zone_territory', 'auto_assignment', 'multi_depot',
    'vehicle_load_planning', 'driver_efficiency', 'delivery_sla', 'incident_reporting',

    -- Delivery & POD
    'signature_pod', 'photo_pod', 'otp_qr_delivery', 'document_scan_pod',
    'task_checklist', 'return_to_depot',

    -- Inventory
    'vehicle_stock', 'warehouse_stock', 'stock_movement', 'realtime_driver_inventory',
    'eod_stock_return', 'auto_reorder',

    -- Payments & Finance
    'cash_collection', 'online_payment', 'cash_discrepancy', 'driver_reconciliation',
    'client_billing', 'invoice_generation',

    -- Communication
    'driver_admin_chat', 'broadcast_messaging', 'sos_emergency', 'issue_reporting',

    -- Integration
    'api_access', 'webhooks', 'third_party_delivery', 'accounting_integration',
    'erp_integration', 'crm_integration',

    -- SaaS Features
    'custom_branding', 'whitelabel_app', 'roles_permissions', 'audit_logs',
    'ai_driver_assignment', 'ai_route_planning'
);
