BEGIN;

-- -------------------------------------------------------------------
-- Reset sample data
-- -------------------------------------------------------------------
DELETE FROM myschema.users;

-- 1) Ensure required users exist (create if missing)
INSERT INTO myschema.users (username, password, role, address, country)
VALUES ('admin', 'admin', 'admin', 'Admin street 1', 'Sweden')
ON CONFLICT (username) DO UPDATE
SET password = EXCLUDED.password,
    role     = EXCLUDED.role;

INSERT INTO myschema.users (username, password, role, address, country)
VALUES ('buyer', 'buyer', 'buyer', 'Buyer lane 1', 'Sweden')
ON CONFLICT (username) DO UPDATE
SET password = EXCLUDED.password,
    role     = EXCLUDED.role;

INSERT INTO myschema.users (username, password, role, address, country)
VALUES ('seller', 'seller', 'seller', 'Seller road 1', 'Sweden')
ON CONFLICT (username) DO UPDATE
SET password = EXCLUDED.password,
    role     = EXCLUDED.role;

-- 2) Clear existing sample data (keeps users)
DELETE FROM myschema.basketitem;
DELETE FROM myschema.orderitem;
DELETE FROM myschema.order;
DELETE FROM myschema.basket;
DELETE FROM myschema.products;

-- -------------------------------------------------------------------
-- Products
-- -------------------------------------------------------------------
-- 3) Insert products for the seller user
INSERT INTO myschema.products
(product_name, manufacturer, seller_user_id, description, screen_size, picture_url, price, stock)
VALUES
('UltraView 24"', 'Acme',
 (SELECT user_id FROM myschema.users WHERE username='seller'),
 '24-inch IPS monitor', 24.0, 'https://example.com/monitor1.jpg', 149.99, 200),

('UltraView 27"', 'Acme',
 (SELECT user_id FROM myschema.users WHERE username='seller'),
 '27-inch QHD monitor', 27.0, 'https://example.com/monitor2.jpg', 249.99, 150),

('ProScreen 32"', 'ViewBest',
 (SELECT user_id FROM myschema.users WHERE username='seller'),
 '32-inch 4K monitor', 32.0, 'https://example.com/monitor3.jpg', 399.00, 100),

('Office 22"', 'ViewBest',
 (SELECT user_id FROM myschema.users WHERE username='seller'),
 '22-inch budget monitor', 22.0, 'https://example.com/monitor4.jpg', 99.00, 300);

-- -------------------------------------------------------------------
-- Baskets (optional sample)
-- -------------------------------------------------------------------
-- Create a basket for buyer
INSERT INTO myschema.basket (basket_user_id)
VALUES ((SELECT user_id FROM myschema.users WHERE username='buyer'));

-- Add items to basket
INSERT INTO myschema.basketitem (basket_id, product_id, quantity)
VALUES
(
  (SELECT basket_id FROM myschema.basket b
   JOIN myschema.users u ON u.user_id=b.basket_user_id
   WHERE u.username='buyer'
   ORDER BY b.basket_id DESC
   LIMIT 1),
  (SELECT product_id FROM myschema.products WHERE product_name='UltraView 24"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='UltraView 24"')
),
(
  (SELECT basket_id FROM myschema.basket b
   JOIN myschema.users u ON u.user_id=b.basket_user_id
   WHERE u.username='buyer'
   ORDER BY b.basket_id DESC
   LIMIT 1),
  (SELECT product_id FROM myschema.products WHERE product_name='Office 22"'),
  2,
  (SELECT price FROM myschema.products WHERE product_name='Office 22"')
);

-- -------------------------------------------------------------------
-- Orders
-- Requirement: one user has multiple orders (example: buyer has 8 orders)
-- -------------------------------------------------------------------

-- Create 8 separate orders for buyer
-- Each order has its own totalprice and its own orderitems.
-- NOTE: We use a CTE with RETURNING so each order gets the correct order_id.
WITH new_orders AS (
  INSERT INTO myschema.order (order_user_id, totalprice)
  SELECT
    (SELECT user_id FROM myschema.users WHERE username='buyer'),
    v.totalprice
  FROM (VALUES
    -- Order 1: 1x UltraView 24" + 2x Office 22"
    (
      (SELECT price FROM myschema.products WHERE product_name='UltraView 24"') * 1
    + (SELECT price FROM myschema.products WHERE product_name='Office 22"')   * 2
    ),

    -- Order 2: 1x UltraView 27"
    (
      (SELECT price FROM myschema.products WHERE product_name='UltraView 27"') * 1
    ),

    -- Order 3: 1x ProScreen 32"
    (
      (SELECT price FROM myschema.products WHERE product_name='ProScreen 32"') * 1
    ),

    -- Order 4: 2x Office 22"
    (
      (SELECT price FROM myschema.products WHERE product_name='Office 22"') * 2
    ),

    -- Order 5: 2x UltraView 24"
    (
      (SELECT price FROM myschema.products WHERE product_name='UltraView 24"') * 2
    ),

    -- Order 6: 1x UltraView 24" + 1x UltraView 27"
    (
      (SELECT price FROM myschema.products WHERE product_name='UltraView 24"') * 1
    + (SELECT price FROM myschema.products WHERE product_name='UltraView 27"') * 1
    ),

    -- Order 7: 1x UltraView 27" + 1x Office 22"
    (
      (SELECT price FROM myschema.products WHERE product_name='UltraView 27"') * 1
    + (SELECT price FROM myschema.products WHERE product_name='Office 22"')   * 1
    ),

    -- Order 8: 1x ProScreen 32" + 2x Office 22"
    (
      (SELECT price FROM myschema.products WHERE product_name='ProScreen 32"') * 1
    + (SELECT price FROM myschema.products WHERE product_name='Office 22"') * 2
    )
  ) AS v(totalprice)
  RETURNING order_id
),
-- Attach a row number so we can map items to the right inserted order
orders_numbered AS (
  SELECT order_id, ROW_NUMBER() OVER (ORDER BY order_id) AS rn
  FROM new_orders
)
INSERT INTO myschema.orderitem (order_id, product_id, quantity, price)
-- Order 1 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=1),
  (SELECT product_id FROM myschema.products WHERE product_name='UltraView 24"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='UltraView 24"')
UNION ALL
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=1),
  (SELECT product_id FROM myschema.products WHERE product_name='Office 22"'),
  2,
  (SELECT price FROM myschema.products WHERE product_name='Office 22"')

UNION ALL
-- Order 2 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=2),
  (SELECT product_id FROM myschema.products WHERE product_name='UltraView 27"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='UltraView 27"')

UNION ALL
-- Order 3 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=3),
  (SELECT product_id FROM myschema.products WHERE product_name='ProScreen 32"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='ProScreen 32"')

UNION ALL
-- Order 4 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=4),
  (SELECT product_id FROM myschema.products WHERE product_name='Office 22"'),
  2,
  (SELECT price FROM myschema.products WHERE product_name='Office 22"')

UNION ALL
-- Order 5 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=5),
  (SELECT product_id FROM myschema.products WHERE product_name='UltraView 24"'),
  2,
  (SELECT price FROM myschema.products WHERE product_name='UltraView 24"')

UNION ALL
-- Order 6 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=6),
  (SELECT product_id FROM myschema.products WHERE product_name='UltraView 24"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='UltraView 24"')
UNION ALL
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=6),
  (SELECT product_id FROM myschema.products WHERE product_name='UltraView 27"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='UltraView 27"')

UNION ALL
-- Order 7 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=7),
  (SELECT product_id FROM myschema.products WHERE product_name='UltraView 27"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='UltraView 27"')
UNION ALL
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=7),
  (SELECT product_id FROM myschema.products WHERE product_name='Office 22"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='Office 22"')

UNION ALL
-- Order 8 items
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=8),
  (SELECT product_id FROM myschema.products WHERE product_name='ProScreen 32"'),
  1,
  (SELECT price FROM myschema.products WHERE product_name='ProScreen 32"')
UNION ALL
SELECT
  (SELECT order_id FROM orders_numbered WHERE rn=8),
  (SELECT product_id FROM myschema.products WHERE product_name='Office 22"'),
  2,
  (SELECT price FROM myschema.products WHERE product_name='Office 22"');

COMMIT;
