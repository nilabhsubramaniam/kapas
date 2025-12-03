# Admin API Documentation

## Overview
This document describes the backend API architecture for the Tantuka Admin Panel (Angular).

## Authentication
All admin endpoints require:
1. Valid JWT token in `Authorization: Bearer <token>` header
2. User role must be `admin`

## API Architecture Pattern

### Customer APIs vs Admin APIs

```
┌─────────────────────────────────────────────────────────────┐
│                    CUSTOMER APIS                             │
│  Scope: Current logged-in user only                         │
├─────────────────────────────────────────────────────────────┤
│  GET  /api/orders          → My orders only                 │
│  GET  /api/cart            → My cart only                   │
│  GET  /api/addresses       → My addresses only              │
│  GET  /api/auth/me         → My profile                     │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│                     ADMIN APIS                               │
│  Scope: Global access to ALL users' data                    │
├─────────────────────────────────────────────────────────────┤
│  GET  /api/admin/orders           → ALL orders              │
│  GET  /api/admin/users            → ALL users               │
│  GET  /api/admin/users/:id/orders → Any user's orders       │
│  GET  /api/admin/analytics        → Global analytics        │
└─────────────────────────────────────────────────────────────┘
```

---

## Admin Endpoints

### 📊 Dashboard & Analytics

#### 1. Get Dashboard Statistics
```http
GET /api/admin/dashboard
Authorization: Bearer <admin_token>
```

**Response:**
```json
{
  "total_users": 1250,
  "total_orders": 3400,
  "total_products": 150,
  "total_revenue": 1250000.50,
  "pending_orders": 45,
  "low_stock_products": 12
}
```

**Angular Service Example:**
```typescript
getDashboardStats(): Observable<DashboardStats> {
  return this.http.get<DashboardStats>('/api/admin/dashboard');
}
```

---

#### 2. Get Sales Analytics
```http
GET /api/admin/analytics/sales?period=month
Authorization: Bearer <admin_token>
```

**Query Parameters:**
- `period`: `day` | `week` | `month` | `year`

**Response:**
```json
{
  "period": "month",
  "data": [
    {
      "date": "2024-01-01",
      "orders": 120,
      "revenue": 45000.00
    },
    {
      "date": "2024-02-01",
      "orders": 135,
      "revenue": 52000.00
    }
  ]
}
```

**Use Case:** Line charts showing sales trends over time

---

#### 3. Get Revenue Analytics
```http
GET /api/admin/analytics/revenue
Authorization: Bearer <admin_token>
```

**Response:**
```json
{
  "by_product_type": [
    {
      "product_type": "SAREE",
      "revenue": 850000.00,
      "orders": 2100
    },
    {
      "product_type": "CHIKANKARI_KURTI",
      "revenue": 320000.00,
      "orders": 980
    }
  ],
  "by_state": [
    {
      "state": "UP",
      "revenue": 450000.00,
      "orders": 1200
    },
    {
      "state": "KL",
      "revenue": 380000.00,
      "orders": 950
    }
  ]
}
```

**Use Case:** Pie charts, bar charts for revenue breakdown

---

### 👥 User Management

#### 4. List All Users
```http
GET /api/admin/users?page=1&per_page=20&role=customer&search=john
Authorization: Bearer <admin_token>
```

**Query Parameters:**
- `page`: Page number (default: 1)
- `per_page`: Items per page (default: 20)
- `role`: Filter by `customer` | `admin` | `vendor`
- `search`: Search by name or email

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "email": "customer@example.com",
      "name": "John Doe",
      "phone": "+919876543210",
      "role": "customer",
      "is_active": true,
      "email_verified": true,
      "created_at": "2024-01-15T10:30:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 20,
    "total": 1250,
    "total_pages": 63
  }
}
```

**Angular Component:**
```typescript
loadUsers(page: number, filters: UserFilters) {
  this.adminService.getUsers(page, filters).subscribe(
    response => {
      this.users = response.data;
      this.totalUsers = response.pagination.total;
    }
  );
}
```

---

#### 5. Get User Details
```http
GET /api/admin/users/:id
Authorization: Bearer <admin_token>
```

**Response:**
```json
{
  "id": 1,
  "email": "customer@example.com",
  "name": "John Doe",
  "phone": "+919876543210",
  "role": "customer",
  "is_active": true,
  "email_verified": true,
  "addresses": [
    {
      "id": 1,
      "address_type": "HOME",
      "address_line1": "123 Main St",
      "pincode": "110001",
      "is_default": true
    }
  ],
  "created_at": "2024-01-15T10:30:00Z",
  "last_login": "2024-12-01T14:20:00Z"
}
```

**Use Case:** User profile modal/page in admin panel

---

#### 6. Get User's Orders
```http
GET /api/admin/users/:id/orders?page=1&per_page=10
Authorization: Bearer <admin_token>
```

**Response:**
```json
{
  "data": [
    {
      "id": 1001,
      "order_number": "ORD-2024-1001",
      "status": "DELIVERED",
      "total_amount": 4500.00,
      "created_at": "2024-11-20T10:00:00Z",
      "order_items": [
        {
          "product_id": 5,
          "product_name": "Lucknow Chikankari Saree",
          "quantity": 1,
          "price": 3999.00
        }
      ]
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 10,
    "total": 15,
    "total_pages": 2
  }
}
```

**Use Case:** View customer's order history in admin panel

---

#### 7. Update User Status
```http
PUT /api/admin/users/:id/status
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "is_active": false,
  "email_verified": true,
  "role": "vendor"
}
```

**Response:**
```json
{
  "message": "User updated successfully"
}
```

**Use Case:** 
- Deactivate problematic users
- Verify email manually
- Promote customer to vendor

---

### 📦 Order Management

#### 8. List All Orders
```http
GET /api/admin/orders?page=1&per_page=20&status=PENDING&user_id=123
Authorization: Bearer <admin_token>
```

**Query Parameters:**
- `page`: Page number
- `per_page`: Items per page
- `status`: Filter by `PENDING` | `CONFIRMED` | `SHIPPED` | `DELIVERED` | `CANCELLED`
- `user_id`: Filter by specific user

**Response:**
```json
{
  "data": [
    {
      "id": 1001,
      "order_number": "ORD-2024-1001",
      "user": {
        "id": 123,
        "name": "John Doe",
        "email": "john@example.com"
      },
      "status": "PENDING",
      "total_amount": 4500.00,
      "payment": {
        "status": "PAID",
        "method": "UPI"
      },
      "order_items": [
        {
          "product": {
            "id": 5,
            "name": "Lucknow Chikankari Saree"
          },
          "quantity": 1,
          "price": 3999.00
        }
      ],
      "created_at": "2024-12-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 20,
    "total": 3400,
    "total_pages": 170
  }
}
```

**Angular Usage:**
```typescript
// In your orders.component.ts
loadOrders(filters: OrderFilters) {
  this.adminService.getAllOrders(filters).subscribe(
    response => {
      this.orders = response.data;
      this.displayedOrders = this.orders; // For AG-Grid or Material Table
    }
  );
}
```

---

#### 9. Update Order Status
```http
PUT /api/admin/orders/:id/status
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "status": "SHIPPED",
  "tracking_number": "DHL123456789",
  "notes": "Package shipped via DHL, expected delivery in 3 days"
}
```

**Response:**
```json
{
  "message": "Order updated successfully"
}
```

**Use Case:** Update order workflow: PENDING → CONFIRMED → SHIPPED → DELIVERED

---

### 📊 Inventory Management

#### 10. Get Inventory
```http
GET /api/admin/inventory?page=1&per_page=50&low_stock=true
Authorization: Bearer <admin_token>
```

**Query Parameters:**
- `page`: Page number
- `per_page`: Items per page
- `low_stock`: If `true`, show only items with stock < 10

**Response:**
```json
{
  "data": [
    {
      "id": 5,
      "name": "Lucknow White Chikankari Cotton Saree",
      "slug": "lucknow-white-chikankari-cotton-saree",
      "product_type": "SAREE",
      "state_origin": "UP",
      "stock_quantity": 8,
      "base_price": 4999.00,
      "final_price": 3999.00,
      "is_active": true
    }
  ],
  "pagination": {
    "page": 1,
    "per_page": 50,
    "total": 150,
    "total_pages": 3
  }
}
```

**Use Case:** Stock management dashboard with alerts for low stock

---

#### 11. Update Inventory
```http
PUT /api/admin/inventory/:id
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "stock_quantity": 50,
  "is_active": true
}
```

**Response:**
```json
{
  "message": "Inventory updated successfully"
}
```

---

## Angular Admin Panel Architecture

### Recommended Structure

```
tantuka-admin/
├── src/
│   ├── app/
│   │   ├── core/
│   │   │   ├── guards/
│   │   │   │   └── admin-auth.guard.ts    // Check admin role
│   │   │   ├── interceptors/
│   │   │   │   └── auth.interceptor.ts    // Add JWT token
│   │   │   └── services/
│   │   │       ├── auth.service.ts
│   │   │       └── admin-api.service.ts
│   │   │
│   │   ├── features/
│   │   │   ├── dashboard/
│   │   │   │   ├── dashboard.component.ts
│   │   │   │   └── widgets/
│   │   │   │       ├── stats-card.component.ts
│   │   │   │       └── sales-chart.component.ts
│   │   │   │
│   │   │   ├── users/
│   │   │   │   ├── user-list/
│   │   │   │   ├── user-detail/
│   │   │   │   └── user-orders/
│   │   │   │
│   │   │   ├── orders/
│   │   │   │   ├── order-list/
│   │   │   │   ├── order-detail/
│   │   │   │   └── order-status-update/
│   │   │   │
│   │   │   └── inventory/
│   │   │       ├── inventory-list/
│   │   │       └── stock-update/
│   │   │
│   │   └── shared/
│   │       ├── components/
│   │       │   ├── data-table/
│   │       │   └── pagination/
│   │       └── models/
│   │           ├── user.model.ts
│   │           ├── order.model.ts
│   │           └── product.model.ts
```

### Key Angular Services

#### admin-api.service.ts
```typescript
import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class AdminApiService {
  private baseUrl = 'http://localhost:8080/api/admin';

  constructor(private http: HttpClient) {}

  // Dashboard
  getDashboard(): Observable<DashboardStats> {
    return this.http.get<DashboardStats>(`${this.baseUrl}/dashboard`);
  }

  getSalesAnalytics(period: string): Observable<SalesData> {
    return this.http.get<SalesData>(`${this.baseUrl}/analytics/sales`, {
      params: { period }
    });
  }

  // Users
  getUsers(page: number, filters: UserFilters): Observable<PaginatedResponse<User>> {
    let params = new HttpParams()
      .set('page', page.toString())
      .set('per_page', '20');
    
    if (filters.role) params = params.set('role', filters.role);
    if (filters.search) params = params.set('search', filters.search);

    return this.http.get<PaginatedResponse<User>>(`${this.baseUrl}/users`, { params });
  }

  getUserDetails(userId: number): Observable<User> {
    return this.http.get<User>(`${this.baseUrl}/users/${userId}`);
  }

  getUserOrders(userId: number, page: number): Observable<PaginatedResponse<Order>> {
    return this.http.get<PaginatedResponse<Order>>(
      `${this.baseUrl}/users/${userId}/orders`,
      { params: { page: page.toString() } }
    );
  }

  updateUserStatus(userId: number, data: UserStatusUpdate): Observable<MessageResponse> {
    return this.http.put<MessageResponse>(`${this.baseUrl}/users/${userId}/status`, data);
  }

  // Orders
  getAllOrders(filters: OrderFilters): Observable<PaginatedResponse<Order>> {
    let params = new HttpParams()
      .set('page', filters.page.toString())
      .set('per_page', '20');
    
    if (filters.status) params = params.set('status', filters.status);
    if (filters.userId) params = params.set('user_id', filters.userId.toString());

    return this.http.get<PaginatedResponse<Order>>(`${this.baseUrl}/orders`, { params });
  }

  updateOrderStatus(orderId: number, data: OrderStatusUpdate): Observable<MessageResponse> {
    return this.http.put<MessageResponse>(`${this.baseUrl}/orders/${orderId}/status`, data);
  }

  // Inventory
  getInventory(page: number, lowStockOnly: boolean): Observable<PaginatedResponse<Product>> {
    let params = new HttpParams()
      .set('page', page.toString())
      .set('per_page', '50');
    
    if (lowStockOnly) params = params.set('low_stock', 'true');

    return this.http.get<PaginatedResponse<Product>>(`${this.baseUrl}/inventory`, { params });
  }

  updateInventory(productId: number, data: InventoryUpdate): Observable<MessageResponse> {
    return this.http.put<MessageResponse>(`${this.baseUrl}/inventory/${productId}`, data);
  }
}
```

#### auth.interceptor.ts
```typescript
import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler } from '@angular/common/http';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler) {
    const token = localStorage.getItem('admin_token');
    
    if (token) {
      const cloned = req.clone({
        headers: req.headers.set('Authorization', `Bearer ${token}`)
      });
      return next.handle(cloned);
    }
    
    return next.handle(req);
  }
}
```

#### admin-auth.guard.ts
```typescript
import { Injectable } from '@angular/core';
import { Router, CanActivate } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Injectable({ providedIn: 'root' })
export class AdminAuthGuard implements CanActivate {
  constructor(
    private auth: AuthService,
    private router: Router
  ) {}

  canActivate(): boolean {
    const user = this.auth.getCurrentUser();
    
    if (user && user.role === 'admin') {
      return true;
    }
    
    this.router.navigate(['/login']);
    return false;
  }
}
```

---

## Summary

### ✅ What Your Backend Provides for Angular Admin

1. **Dashboard APIs**
   - Real-time statistics
   - Sales trends
   - Revenue breakdown

2. **User Management APIs**
   - List all users with search/filter
   - View any user's profile
   - View any user's orders
   - Update user status/role

3. **Order Management APIs**
   - List all orders from all customers
   - Filter by status, user, date
   - Update order status with tracking
   - Activity logging

4. **Inventory Management APIs**
   - View all products with stock
   - Low stock alerts
   - Update stock quantities

### 🔐 Security Layer
- All admin endpoints protected by:
  - JWT authentication (valid token required)
  - Admin role check (role must be "admin")
  - Middleware: `AuthMiddleware() + AdminOnly()`

### 📊 Data Flow

```
Angular Admin Panel
        ↓
  HTTP Request with JWT Token
        ↓
  Go Backend (AuthMiddleware + AdminOnly)
        ↓
  PostgreSQL (Global queries - ALL users' data)
        ↓
  JSON Response
        ↓
  Angular displays in tables/charts
```

This architecture gives your Angular admin panel **complete visibility and control** over all customer data, orders, and inventory! 🚀
