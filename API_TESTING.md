# API Testing Guide

## Products API

### 1. List All Products
```bash
curl http://localhost:8080/api/products
```

**Expected Response:**
- 10 products with pagination
- Each product includes: id, name, slug, description, product_type, state_origin, saree_type, base_price, discount_percentage, final_price, fabric, weave_type, occasion, stock_quantity, images, metadata

### 2. Filter Products by State
```bash
# Uttar Pradesh Chikankari
curl "http://localhost:8080/api/products?state=UP"

# Kerala Kasavu
curl "http://localhost:8080/api/products?state=KL"

# Tamil Nadu Kanchipuram
curl "http://localhost:8080/api/products?state=TN"
```

### 3. Filter by Saree Type
```bash
# Chikankari Sarees
curl "http://localhost:8080/api/products?saree_type=Chikankari"

# Kasavu Sarees
curl "http://localhost:8080/api/products?saree_type=Kasavu"

# Kanchipuram Sarees
curl "http://localhost:8080/api/products?saree_type=Kanchipuram"
```

### 4. Filter by Product Type
```bash
# Sarees Only
curl "http://localhost:8080/api/products?product_type=SAREE"

# Kurtis Only
curl "http://localhost:8080/api/products?product_type=CHIKANKARI_KURTI"
```

### 5. Filter by Fabric
```bash
# Pure Silk
curl "http://localhost:8080/api/products?fabric=Pure Silk"

# Cotton
curl "http://localhost:8080/api/products?fabric=Cotton"

# Georgette
curl "http://localhost:8080/api/products?fabric=Georgette"
```

### 6. Filter by Occasion
```bash
# Wedding Sarees
curl "http://localhost:8080/api/products?occasion=Wedding"

# Casual Wear
curl "http://localhost:8080/api/products?occasion=Casual"

# Festival Sarees
curl "http://localhost:8080/api/products?occasion=Festival"
```

### 7. Filter by Price Range
```bash
# Products under ₹3000
curl "http://localhost:8080/api/products?max_price=3000"

# Products between ₹5000-10000
curl "http://localhost:8080/api/products?min_price=5000&max_price=10000"

# Premium products above ₹10000
curl "http://localhost:8080/api/products?min_price=10000"
```

### 8. Combine Multiple Filters
```bash
# Silk sarees from Tamil Nadu for weddings
curl "http://localhost:8080/api/products?state=TN&fabric=Pure Silk&occasion=Wedding"

# Cotton products under ₹3000
curl "http://localhost:8080/api/products?fabric=Cotton&max_price=3000"
```

### 9. Sorting
```bash
# Price: Low to High
curl "http://localhost:8080/api/products?sort=price_asc"

# Price: High to Low
curl "http://localhost:8080/api/products?sort=price_desc"

# Newest First
curl "http://localhost:8080/api/products?sort=newest"

# Most Popular
curl "http://localhost:8080/api/products?sort=popular"
```

### 10. Get Single Product by Slug
```bash
# Lucknow Chikankari Saree
curl http://localhost:8080/api/products/lucknow-white-chikankari-cotton-saree

# Kanchipuram Silk Saree
curl http://localhost:8080/api/products/kanchipuram-pure-silk-saree-red-gold

# Kerala Kasavu
curl http://localhost:8080/api/products/kerala-traditional-kasavu-cotton-saree
```

### 11. Pagination
```bash
# Page 1, 10 items per page (default)
curl "http://localhost:8080/api/products?page=1&per_page=10"

# Page 2, 5 items per page
curl "http://localhost:8080/api/products?page=2&per_page=5"
```

---

## Create Product (Admin Only)

**Endpoint:** `POST /api/products`  
**Authorization:** Bearer Token (Admin role required)

### Example Request:
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -d '{
    "name": "Banarasi Silk Saree - Golden Yellow",
    "description": "Luxurious Banarasi silk saree with intricate golden zari work",
    "product_type": "SAREE",
    "state_origin": "UP",
    "saree_type": "Banarasi",
    "base_price": 12999,
    "discount_percentage": 15,
    "fabric": "Pure Silk",
    "weave_type": "Handloom",
    "occasion": "Wedding, Festival",
    "stock_quantity": 20,
    "images": [
      {
        "image_url": "/images/products/banarasi-golden-yellow-1.jpg",
        "alt_text": "Banarasi Silk Saree - Golden Yellow - Front View",
        "display_order": 1,
        "is_primary": true
      },
      {
        "image_url": "/images/products/banarasi-golden-yellow-2.jpg",
        "alt_text": "Banarasi Silk Saree - Golden Yellow - Detail View",
        "display_order": 2,
        "is_primary": false
      }
    ],
    "metadata": {
      "length": "6.5 meters",
      "blouse": "Included",
      "care": "Dry Clean Only",
      "color": "Golden Yellow",
      "pattern": "Zari Work",
      "washable": false
    }
  }'
```

---

## Update Product (Admin Only)

**Endpoint:** `PUT /api/products/:slug`  
**Authorization:** Bearer Token (Admin role required)

### Example Request:
```bash
curl -X PUT http://localhost:8080/api/products/lucknow-white-chikankari-cotton-saree \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN" \
  -d '{
    "base_price": 5499,
    "discount_percentage": 25,
    "stock_quantity": 30
  }'
```

---

## Delete Product (Admin Only)

**Endpoint:** `DELETE /api/products/:slug`  
**Authorization:** Bearer Token (Admin role required)

### Example Request:
```bash
curl -X DELETE http://localhost:8080/api/products/lucknow-white-chikankari-cotton-saree \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

---

## Sample Products in Database

| ID | Name | Type | State | Fabric | Price | Stock |
|----|------|------|-------|--------|-------|-------|
| 1 | Lucknow White Chikankari Cotton Saree | SAREE | UP | Cotton | ₹3,999 | 25 |
| 2 | Lucknow Pastel Pink Chikankari Georgette Saree | SAREE | UP | Georgette | ₹5,949 | 15 |
| 3 | Kerala Traditional Kasavu Cotton Saree | SAREE | KL | Cotton | ₹3,149 | 30 |
| 4 | Kerala Kasavu Silk Saree with Tissue Border | SAREE | KL | Silk | ₹7,919 | 10 |
| 5 | Kanchipuram Pure Silk Saree - Red & Gold | SAREE | TN | Pure Silk | ₹13,119 | 8 |
| 6 | Kanchipuram Silk Saree - Royal Blue | SAREE | TN | Pure Silk | ₹15,199 | 5 |
| 7 | Mysore Silk Saree - Green & Maroon | SAREE | KA | Pure Silk | ₹6,799 | 12 |
| 8 | Bengal Tant Cotton Saree - Red with White Border | SAREE | WB | Cotton | ₹2,249 | 40 |
| 9 | Lucknow Chikankari White Cotton Kurti | KURTI | UP | Cotton | ₹1,599 | 50 |
| 10 | Lucknow Chikankari Black Cotton Kurti | KURTI | UP | Cotton | ₹2,124 | 35 |

---

## Product Variants & Attributes

### Saree-Specific Attributes (in metadata):
- `length`: Saree length (e.g., "6.5 meters", "5.5 meters")
- `blouse`: Whether blouse is included ("Included", "Not Included")
- `care`: Care instructions ("Dry Clean Only", "Hand Wash")
- `color`: Primary color(s)
- `pattern`: Pattern type (e.g., "Chikankari", "Zari Work", "Temple Border")
- `washable`: Boolean - whether machine/hand washable
- `temple`: Special feature for Kanchipuram sarees ("Temple Border")

### Kurti-Specific Attributes (in metadata):
- `size`: Available sizes (e.g., "S, M, L, XL, XXL")
- `length`: Kurti length in inches (e.g., "42 inches")
- `neckline`: Neckline style ("Round Neck", "V Neck", "Boat Neck")
- `sleeves`: Sleeve type ("3/4 Sleeves", "Full Sleeves", "Sleeveless")
- `color`: Primary color
- `pattern`: Embroidery pattern
- `care`: Care instructions
- `washable`: Boolean

---

## Frontend Integration

### Example: Fetch and Display Products in Next.js

```typescript
// app/products/page.tsx
async function getProducts(filters?: {
  state?: string;
  saree_type?: string;
  product_type?: string;
  fabric?: string;
  occasion?: string;
  min_price?: number;
  max_price?: number;
  sort?: string;
  page?: number;
  per_page?: number;
}) {
  const params = new URLSearchParams();
  
  if (filters) {
    Object.entries(filters).forEach(([key, value]) => {
      if (value !== undefined) {
        params.append(key, value.toString());
      }
    });
  }
  
  const response = await fetch(
    `http://localhost:8080/api/products?${params.toString()}`,
    { next: { revalidate: 60 } } // Cache for 60 seconds
  );
  
  if (!response.ok) {
    throw new Error('Failed to fetch products');
  }
  
  return response.json();
}

export default async function ProductsPage() {
  const data = await getProducts({
    product_type: 'SAREE',
    sort: 'price_asc'
  });
  
  return (
    <div>
      <h1>Products</h1>
      <p>Total: {data.total} products</p>
      
      <div className="grid grid-cols-3 gap-4">
        {data.products.map((product) => (
          <div key={product.id} className="border p-4">
            <img 
              src={product.images[0]?.image_url} 
              alt={product.images[0]?.alt_text}
            />
            <h2>{product.name}</h2>
            <p className="line-through">₹{product.base_price}</p>
            <p className="text-green-600">₹{product.final_price}</p>
            <p className="text-red-500">{product.discount_percentage}% OFF</p>
            <p>{product.fabric} • {product.state_origin}</p>
            <p>Stock: {product.stock_quantity}</p>
          </div>
        ))}
      </div>
      
      {/* Pagination */}
      <div>
        Page {data.page} of {Math.ceil(data.total / data.per_page)}
      </div>
    </div>
  );
}
```

---

## Next Steps

1. **Frontend Integration:**
   - Connect Next.js product listing page to API
   - Create product detail page using slug
   - Implement filters sidebar
   - Add sorting dropdown
   - Add pagination controls

2. **Cart System:**
   - Add to Cart API
   - Update Cart Quantity
   - Remove from Cart
   - Get Cart Total

3. **Checkout Flow:**
   - Address Management
   - Order Creation
   - Payment Integration (Razorpay)

4. **Admin Panel:**
   - Product Management Dashboard
   - Inventory Management
   - Order Processing
   - Analytics

---

## Testing Checklist

- [ ] List all products
- [ ] Filter by state (UP, KL, TN, KA, WB)
- [ ] Filter by saree type (Chikankari, Kasavu, Kanchipuram, etc.)
- [ ] Filter by product type (SAREE, CHIKANKARI_KURTI)
- [ ] Filter by fabric (Cotton, Silk, Georgette)
- [ ] Filter by occasion (Wedding, Festival, Casual)
- [ ] Filter by price range
- [ ] Combine multiple filters
- [ ] Sort by price (asc/desc)
- [ ] Sort by newest
- [ ] Get single product by slug
- [ ] Pagination
- [ ] Create product (admin)
- [ ] Update product (admin)
- [ ] Delete product (admin)
