BEGIN;

-- companies
INSERT INTO 
    companies (company_id, name, representative_name, phone_number, postal_code, address, created_at, updated_at) 
VALUES 
    ('company1', 'Company A', 'John Doe', '1234567890', '123-4567', 'Address A', NOW(), NOW()),
    ('company2', 'Company B', 'Jane Smith', '2345678901', '234-5678', 'Address B', NOW(), NOW()),
    ('company3', 'Company C', 'Alice Johnson', '3456789012', '345-6789', 'Address C', NOW(), NOW());

-- users
INSERT INTO 
    users (user_id, company_id, name, email, password, created_at, updated_at) 
VALUES 
    ('user1', 'company1', 'User A1', 'usera1@email.com', 'passwordA1', NOW(), NOW()),
    ('user2', 'company2', 'User B1', 'userb1@email.com', 'passwordB1', NOW(), NOW()),
    ('user3', 'company3', 'User C1', 'userc1@email.com', 'passwordC1', NOW(), NOW());

-- clients
INSERT INTO 
    clients (company_id, client_id)
VALUES 
    ('company1', 'company2'),
    ('company1', 'company3'),
    ('company2', 'company1');

-- company_banks
INSERT INTO 
    company_banks (company_bank_id, company_id, bank_code, branch_code, bank_name, branch_name, account_type, account_number, account_holder, created_at, updated_at) 
VALUES 
    ('companybank1', 'company1', '001', '001A', 'Bank A', 'Branch A1', 'Checking', '12345678', 'Holder A1', NOW(), NOW()),
    ('companybank2', 'company2', '002', '002B', 'Bank B', 'Branch B1', 'Savings', '23456789', 'Holder B1', NOW(), NOW()),
    ('companybank3', 'company3', '003', '003C', 'Bank C', 'Branch C1', 'Checking', '34567890', 'Holder C1', NOW(), NOW());

-- invoices
INSERT INTO 
    invoices (invoice_id, company_id, client_id, payment_amount, fee, fee_rate, tax, tax_rate, total_amount, due_at, created_at, updated_at) 
VALUES 
    ('invoice1', 'company1', 'company2', 1000, 0.04, 40, 0.10, 4, 1044, NOW() + INTERVAL 7 DAY, NOW(), NOW()),
    ('invoice2', 'company2', 'company3', 1500, 0.04, 60, 0.10, 6, 1566, NOW() + INTERVAL 10 DAY, NOW(), NOW()),
    ('invoice3', 'company3', 'company1', 2000, 0.04, 80, 0.10, 8, 2088, NOW() + INTERVAL 14 DAY, NOW(), NOW());

-- invoice_statuses
INSERT INTO 
    invoice_statuses (status_name) 
VALUES 
    ('Unprocessed'),
    ('Processing'),
    ('Paid'),
    ('Error');

-- invoice_status_logs
INSERT INTO 
    invoice_status_logs (invoice_id, status_id, user_id, created_at, updated_at) 
VALUES 
    ('invoice1', 1, 'user1', NOW(), NOW()),
    ('invoice2', 2, 'user2', NOW(), NOW()),
    ('invoice3', 3, 'user3', NOW(), NOW());


COMMIT;