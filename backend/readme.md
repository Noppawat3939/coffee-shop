# Backend

This is the backend service for the Coffee Shop application, built with Go, Gin framework, and PostgreSQL.

## Features

- REST API built with Gin
- PostgreSQL database with GORM ORM
- Auto migrations for database schema
<!-- - Cascade delete for related price logs on variation delete -->

## Prerequisites

- Go 1.20+ installed
- PostgreSQL database running (e.g., via Docker)
- `docker` and `docker-compose` (optional, for easy DB setup)

# Database Schema

This document describes the database schema for the Ordering & Payment System.

---

## Menu Management

### menus

| Column       | Type | Key | Description         |
| ------------ | ---- | --- | ------------------- |
| id           | int  | PK  | Menu ID             |
| name         | text |     | Menu name           |
| description  | text |     | Menu description    |
| is_available | bool |     | Availability status |

---

### menu_variations

| Column  | Type   | Key           | Description      |
| ------- | ------ | ------------- | ---------------- |
| id      | int    | PK            | Variation ID     |
| menu_id | int    | FK → menus.id | Parent menu      |
| type    | text   |               | hot, ice, frappe |
| price   | number |               | Current price    |

---

### menu_price_logs

| Column            | Type      | Key                     | Description         |
| ----------------- | --------- | ----------------------- | ------------------- |
| id                | int       | PK                      | Price log ID        |
| menu_variation_id | int       | FK → menu_variations.id | Variation reference |
| price             | number    |                         | Historical price    |
| created_at        | timestamp |                         | Created time        |

---

## User Management

### members

| Column       | Type      | Key       | Description                    |
| ------------ | --------- | --------- | ------------------------------ |
| id           | int       | PK        | Member ID                      |
| phone_number | text      |           | Phone number                   |
| provider     | text      |           | Authentication provider (line) |
| profile      | text      | encrypted | Encrypted profile data         |
| created_at   | timestamp |           | Created time                   |
| updated_at   | timestamp |           | Updated time                   |

---

### employees

| Column     | Type      | Key    | Description     |
| ---------- | --------- | ------ | --------------- |
| id         | int       | PK     | Employee ID     |
| username   | text      | UNIQUE | Login username  |
| password   | text      |        | Hashed password |
| name       | text      |        | Employee name   |
| active     | bool      |        | Active status   |
| role       | text      |        | admin, staff    |
| created_at | timestamp |        | Created time    |
| updated_at | timestamp |        | Updated time    |

---

## Orders

### orders

| Column       | Type         | Key               | Description       |
| ------------ | ------------ | ----------------- | ----------------- |
| id           | int          | PK                | Order ID          |
| order_number | text         | UNIQUE            | Order reference   |
| status       | order_status |                   | Order status      |
| customer     | text         |                   | Member or guest   |
| total        | number       |                   | Total amount      |
| employee_id  | int          | FK → employees.id | Staff who handled |
| created_at   | timestamp    |                   | Created time      |
| updated_at   | timestamp    |                   | Updated time      |

---

### order_menu_variation

| Column            | Type   | Key                     | Description     |
| ----------------- | ------ | ----------------------- | --------------- |
| id                | int    | PK                      | Item ID         |
| order_id          | int    | FK → orders.id          | Order reference |
| menu_variation_id | int    | FK → menu_variations.id | Menu variation  |
| amount            | number |                         | Quantity        |
| price             | number |                         | Unit price      |

---

### order_status_logs

| Column     | Type      | Key            | Description     |
| ---------- | --------- | -------------- | --------------- |
| id         | int       | PK             | Log ID          |
| order_id   | int       | FK → orders.id | Order reference |
| status     | text      |                | Status value    |
| created_at | timestamp |                | Logged time     |

---

## Payments

### payment_order_transaction_logs

| Column             | Type                                 | Key            | Description                |
| ------------------ | ------------------------------------ | -------------- | -------------------------- |
| id                 | int                                  | PK             | Transaction ID             |
| order_id           | int                                  | FK → orders.id | Order reference            |
| transaction_number | text                                 |                | Payment transaction number |
| order_number_ref   | text                                 |                | Public order reference     |
| status             | payment_order_transaction_log_status |                | Payment status             |
| payment_code       | text                                 |                | QR / payment payload       |
| qr_signature       | text                                 |                | QR signature               |
| amount             | number                               |                | Payment amount             |
| expired_at         | timestamp                            |                | Expiration time            |
| created_at         | timestamp                            |                | Created time               |

---

## Enums

### order_status

| Value    |
| -------- |
| to_pay   |
| paid     |
| canceled |

---

### payment_order_transaction_log_status

| Value    |
| -------- |
| to_pay   |
| paid     |
| canceled |

---

## Relationships Summary

- One `menu` has many `menu_variations`
- One `menu_variation` has many `menu_price_logs`
- One `order` has many `order_menu_variation`
- One `order` has many `order_status_logs`
- One `order` has many `payment_order_transaction_logs`
- Employees manage orders via `employee_id`

---

## Notes

- Payment logs support retry and expiration handling
- Order status history is tracked separately for auditing
- Price logs preserve historical menu pricing
