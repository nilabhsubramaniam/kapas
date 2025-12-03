package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nilabhsubramaniam/kapas/internal/config"
	"github.com/nilabhsubramaniam/kapas/internal/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	config.InitDatabase()

	log.Println("🌱 Starting database seeding...")

	// Seed products
	seedProducts()

	log.Println("✅ Database seeding completed!")
}

func seedProducts() {
	log.Println("📦 Seeding products...")

	// Get region IDs for products
	var lucknow, kerala, kanchipuram, mysore, bengal models.Region
	config.DB.Where("slug = ?", "lucknow").First(&lucknow)
	config.DB.Where("slug = ?", "kerala").First(&kerala)
	config.DB.Where("slug = ?", "kanchipuram").First(&kanchipuram)
	config.DB.Where("slug = ?", "mysore").First(&mysore)
	config.DB.Where("slug = ?", "bengal").First(&bengal)

	products := []models.Product{
		// Lucknow Chikankari Sarees
		{
			Name:               "Lucknow White Chikankari Cotton Saree",
			Slug:               "lucknow-white-chikankari-cotton-saree",
			Description:        "Elegant white cotton saree with intricate Chikankari embroidery from Lucknow. Perfect for summer occasions.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &lucknow.ID,
			SareeType:          "Chikankari",
			BasePrice:          4999,
			DiscountPercentage: 20,
			FinalPrice:         3999,
			Fabric:             "Cotton",
			WeaveType:          "Hand Embroidered",
			Occasion:           "Casual, Festive",
			StockQuantity:      25,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "6.5 meters",
				"blouse":    "Included",
				"care":      "Dry Clean Only",
				"color":     "White",
				"pattern":   "Chikankari",
				"washable":  false,
			},
		},
		{
			Name:               "Lucknow Pastel Pink Chikankari Georgette Saree",
			Slug:               "lucknow-pastel-pink-chikankari-georgette-saree",
			Description:        "Delicate pastel pink georgette saree with fine Chikankari work. Ideal for parties and celebrations.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &lucknow.ID,
			SareeType:          "Chikankari",
			BasePrice:          6999,
			DiscountPercentage: 15,
			FinalPrice:         5949,
			Fabric:             "Georgette",
			WeaveType:          "Hand Embroidered",
			Occasion:           "Party, Wedding",
			StockQuantity:      15,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "6.5 meters",
				"blouse":    "Included",
				"care":      "Dry Clean Only",
				"color":     "Pastel Pink",
				"pattern":   "Chikankari",
				"washable":  false,
			},
		},

		// Kerala Kasavu Sarees
		{
			Name:               "Kerala Traditional Kasavu Cotton Saree",
			Slug:               "kerala-traditional-kasavu-cotton-saree",
			Description:        "Classic Kerala Kasavu saree with golden zari border. Traditional off-white cotton saree perfect for Onam and festivals.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &kerala.ID,
			SareeType:          "Kasavu",
			BasePrice:          3499,
			DiscountPercentage: 10,
			FinalPrice:         3149,
			Fabric:             "Cotton",
			WeaveType:          "Handloom",
			Occasion:           "Festival, Traditional",
			StockQuantity:      30,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "5.5 meters",
				"blouse":    "Not Included",
				"care":      "Hand Wash",
				"color":     "Off White with Gold Border",
				"pattern":   "Kasavu",
				"washable":  true,
			},
		},
		{
			Name:               "Kerala Kasavu Silk Saree with Tissue Border",
			Slug:               "kerala-kasavu-silk-saree-tissue-border",
			Description:        "Luxurious Kerala Kasavu silk saree with tissue border. Perfect for special occasions and weddings.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &kerala.ID,
			SareeType:          "Kasavu",
			BasePrice:          8999,
			DiscountPercentage: 12,
			FinalPrice:         7919,
			Fabric:             "Silk",
			WeaveType:          "Handloom",
			Occasion:           "Wedding, Formal",
			StockQuantity:      10,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "6 meters",
				"blouse":    "Included",
				"care":      "Dry Clean Only",
				"color":     "Off White with Gold & Silver",
				"pattern":   "Kasavu with Tissue",
				"washable":  false,
			},
		},

		// Tamil Nadu Kanchipuram Sarees
		{
			Name:               "Kanchipuram Pure Silk Saree - Red & Gold",
			Slug:               "kanchipuram-pure-silk-saree-red-gold",
			Description:        "Authentic Kanchipuram pure silk saree in vibrant red with traditional gold zari work. A timeless classic.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &kanchipuram.ID,
			SareeType:          "Kanchipuram",
			BasePrice:          15999,
			DiscountPercentage: 18,
			FinalPrice:         13119,
			Fabric:             "Pure Silk",
			WeaveType:          "Handloom",
			Occasion:           "Wedding, Festival",
			StockQuantity:      8,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "6.5 meters",
				"blouse":    "Included",
				"care":      "Dry Clean Only",
				"color":     "Red & Gold",
				"pattern":   "Traditional Zari",
				"washable":  false,
				"temple":    "Temple Border",
			},
		},
		{
			Name:               "Kanchipuram Silk Saree - Royal Blue",
			Slug:               "kanchipuram-silk-saree-royal-blue",
			Description:        "Stunning royal blue Kanchipuram silk saree with intricate pallu design and contrasting border.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &kanchipuram.ID,
			SareeType:          "Kanchipuram",
			BasePrice:          18999,
			DiscountPercentage: 20,
			FinalPrice:         15199,
			Fabric:             "Pure Silk",
			WeaveType:          "Handloom",
			Occasion:           "Wedding, Grand Occasion",
			StockQuantity:      5,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "6.5 meters",
				"blouse":    "Included",
				"care":      "Dry Clean Only",
				"color":     "Royal Blue",
				"pattern":   "Zari Weave",
				"washable":  false,
				"temple":    "Temple Border",
			},
		},

		// Karnataka Mysore Silk Sarees
		{
			Name:               "Mysore Silk Saree - Green & Maroon",
			Slug:               "mysore-silk-saree-green-maroon",
			Description:        "Elegant Mysore silk saree in green with maroon border. Soft silk with beautiful drape.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &mysore.ID,
			SareeType:          "Mysore Silk",
			BasePrice:          7999,
			DiscountPercentage: 15,
			FinalPrice:         6799,
			Fabric:             "Pure Silk",
			WeaveType:          "Handloom",
			Occasion:           "Festival, Party",
			StockQuantity:      12,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "6 meters",
				"blouse":    "Included",
				"care":      "Dry Clean Only",
				"color":     "Green & Maroon",
				"pattern":   "Plain with Zari Border",
				"washable":  false,
			},
		},

		// West Bengal Tant Sarees
		{
			Name:               "Bengal Tant Cotton Saree - Red with White Border",
			Slug:               "bengal-tant-cotton-saree-red-white",
			Description:        "Traditional Bengal Tant saree in red with classic white border. Lightweight and comfortable for daily wear.",
			ProductType:        models.ProductTypeSaree,
			RegionID:           &bengal.ID,
			SareeType:          "Tant",
			BasePrice:          2499,
			DiscountPercentage: 10,
			FinalPrice:         2249,
			Fabric:             "Cotton",
			WeaveType:          "Handloom",
			Occasion:           "Daily Wear, Casual",
			StockQuantity:      40,
			IsActive:           true,
			Metadata: models.JSONB{
				"length":    "5.5 meters",
				"blouse":    "Not Included",
				"care":      "Hand Wash",
				"color":     "Red & White",
				"pattern":   "Plain",
				"washable":  true,
			},
		},

		// Chikankari Kurtis
		{
			Name:               "Lucknow Chikankari White Cotton Kurti",
			Slug:               "lucknow-chikankari-white-cotton-kurti",
			Description:        "Beautiful white cotton kurti with fine Chikankari embroidery. Perfect for summer.",
			ProductType:        models.ProductTypeChikankariKurti,
			RegionID:           &lucknow.ID,
			SareeType:          "Chikankari",
			BasePrice:          1999,
			DiscountPercentage: 20,
			FinalPrice:         1599,
			Fabric:             "Cotton",
			WeaveType:          "Hand Embroidered",
			Occasion:           "Casual, Office",
			StockQuantity:      50,
			IsActive:           true,
			Metadata: models.JSONB{
				"size":      "S, M, L, XL, XXL",
				"length":    "42 inches",
				"care":      "Hand Wash",
				"color":     "White",
				"pattern":   "Chikankari",
				"washable":  true,
				"neckline":  "Round Neck",
				"sleeves":   "3/4 Sleeves",
			},
		},
		{
			Name:               "Lucknow Chikankari Black Cotton Kurti",
			Slug:               "lucknow-chikankari-black-cotton-kurti",
			Description:        "Elegant black cotton kurti with white Chikankari work. Modern and stylish.",
			ProductType:        models.ProductTypeChikankariKurti,
			RegionID:           &lucknow.ID,
			SareeType:          "Chikankari",
			BasePrice:          2499,
			DiscountPercentage: 15,
			FinalPrice:         2124,
			Fabric:             "Cotton",
			WeaveType:          "Hand Embroidered",
			Occasion:           "Party, Casual",
			StockQuantity:      35,
			IsActive:           true,
			Metadata: models.JSONB{
				"size":      "S, M, L, XL, XXL",
				"length":    "44 inches",
				"care":      "Hand Wash",
				"color":     "Black",
				"pattern":   "Chikankari",
				"washable":  true,
				"neckline":  "V Neck",
				"sleeves":   "Full Sleeves",
			},
		},
	}

	for i, product := range products {
		// Check if product already exists
		var existing models.Product
		if err := config.DB.Where("slug = ?", product.Slug).First(&existing).Error; err == nil {
			log.Printf("Product '%s' already exists, skipping...", product.Name)
			continue
		}

		// Create product
		if err := config.DB.Create(&product).Error; err != nil {
			log.Printf("Failed to create product '%s': %v", product.Name, err)
			continue
		}

		// Create sample images for each product
		images := []models.ProductImage{
			{
				ProductID:    product.ID,
				ImageURL:     "/images/products/product-" + product.Slug + "-1.jpg",
				AltText:      product.Name + " - Front View",
				DisplayOrder: 1,
				IsPrimary:    true,
			},
			{
				ProductID:    product.ID,
				ImageURL:     "/images/products/product-" + product.Slug + "-2.jpg",
				AltText:      product.Name + " - Detail View",
				DisplayOrder: 2,
				IsPrimary:    false,
			},
			{
				ProductID:    product.ID,
				ImageURL:     "/images/products/product-" + product.Slug + "-3.jpg",
				AltText:      product.Name + " - Side View",
				DisplayOrder: 3,
				IsPrimary:    false,
			},
		}

		for _, img := range images {
			if err := config.DB.Create(&img).Error; err != nil {
				log.Printf("Failed to create image for product '%s': %v", product.Name, err)
			}
		}

		log.Printf("✅ Created product %d: %s", i+1, product.Name)
	}

	log.Printf("✅ Seeded %d products", len(products))
}
