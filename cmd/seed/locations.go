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

	log.Println("🌱 Starting location data seeding...")

	// Seed in order: Countries → States → Districts → Regions
	seedCountries()
	seedStates()
	seedDistricts()
	seedRegions()

	log.Println("✅ Location data seeding completed!")
}

func seedCountries() {
	log.Println("🌍 Seeding countries...")

	countries := []models.Country{
		{
			Name:      "India",
			Code:      "IN",
			PhoneCode: "+91",
			Currency:  "INR",
			IsActive:  true,
		},
	}

	for _, country := range countries {
		var existing models.Country
		if err := config.DB.Where("code = ?", country.Code).First(&existing).Error; err == nil {
			log.Printf("Country '%s' already exists, skipping...", country.Name)
			continue
		}

		if err := config.DB.Create(&country).Error; err != nil {
			log.Printf("Failed to create country '%s': %v", country.Name, err)
			continue
		}

		log.Printf("✅ Created country: %s", country.Name)
	}
}

func seedStates() {
	log.Println("🗺️  Seeding Indian states...")

	// Get India country ID
	var india models.Country
	if err := config.DB.Where("code = ?", "IN").First(&india).Error; err != nil {
		log.Fatal("India not found in countries table")
	}

	states := []models.State{
		// Major states for saree/textile production
		{CountryID: india.ID, Name: "Uttar Pradesh", Code: "UP", IsActive: true},
		{CountryID: india.ID, Name: "Kerala", Code: "KL", IsActive: true},
		{CountryID: india.ID, Name: "Tamil Nadu", Code: "TN", IsActive: true},
		{CountryID: india.ID, Name: "Karnataka", Code: "KA", IsActive: true},
		{CountryID: india.ID, Name: "West Bengal", Code: "WB", IsActive: true},
		{CountryID: india.ID, Name: "Bihar", Code: "BR", IsActive: true},
		{CountryID: india.ID, Name: "Maharashtra", Code: "MH", IsActive: true},
		{CountryID: india.ID, Name: "Gujarat", Code: "GJ", IsActive: true},
		{CountryID: india.ID, Name: "Rajasthan", Code: "RJ", IsActive: true},
		{CountryID: india.ID, Name: "Andhra Pradesh", Code: "AP", IsActive: true},
		{CountryID: india.ID, Name: "Telangana", Code: "TG", IsActive: true},
		{CountryID: india.ID, Name: "Odisha", Code: "OR", IsActive: true},
		{CountryID: india.ID, Name: "Madhya Pradesh", Code: "MP", IsActive: true},
		{CountryID: india.ID, Name: "Assam", Code: "AS", IsActive: true},
		{CountryID: india.ID, Name: "Punjab", Code: "PB", IsActive: true},
		{CountryID: india.ID, Name: "Haryana", Code: "HR", IsActive: true},
		{CountryID: india.ID, Name: "Delhi", Code: "DL", IsActive: true},
		{CountryID: india.ID, Name: "Jharkhand", Code: "JH", IsActive: true},
		{CountryID: india.ID, Name: "Uttarakhand", Code: "UK", IsActive: true},
		{CountryID: india.ID, Name: "Himachal Pradesh", Code: "HP", IsActive: true},
	}

	for _, state := range states {
		var existing models.State
		if err := config.DB.Where("code = ? AND country_id = ?", state.Code, india.ID).First(&existing).Error; err == nil {
			log.Printf("State '%s' already exists, skipping...", state.Name)
			continue
		}

		if err := config.DB.Create(&state).Error; err != nil {
			log.Printf("Failed to create state '%s': %v", state.Name, err)
			continue
		}

		log.Printf("✅ Created state: %s", state.Name)
	}
}

func seedDistricts() {
	log.Println("🏙️  Seeding major districts...")

	// Get state IDs
	var up, kl, tn, ka, wb, br, mh models.State
	config.DB.Where("code = ?", "UP").First(&up)
	config.DB.Where("code = ?", "KL").First(&kl)
	config.DB.Where("code = ?", "TN").First(&tn)
	config.DB.Where("code = ?", "KA").First(&ka)
	config.DB.Where("code = ?", "WB").First(&wb)
	config.DB.Where("code = ?", "BR").First(&br)
	config.DB.Where("code = ?", "MH").First(&mh)

	districts := []models.District{
		// Uttar Pradesh - Chikankari hubs
		{StateID: up.ID, Name: "Lucknow", Code: "LKO", IsActive: true},
		{StateID: up.ID, Name: "Varanasi", Code: "VNS", IsActive: true},
		{StateID: up.ID, Name: "Agra", Code: "AGR", IsActive: true},
		{StateID: up.ID, Name: "Kanpur", Code: "KNP", IsActive: true},

		// Kerala - Kasavu saree centers
		{StateID: kl.ID, Name: "Kochi", Code: "KOC", IsActive: true},
		{StateID: kl.ID, Name: "Thiruvananthapuram", Code: "TVM", IsActive: true},
		{StateID: kl.ID, Name: "Kozhikode", Code: "KZD", IsActive: true},
		{StateID: kl.ID, Name: "Thrissur", Code: "TSR", IsActive: true},

		// Tamil Nadu - Kanchipuram silk
		{StateID: tn.ID, Name: "Kanchipuram", Code: "KAN", IsActive: true},
		{StateID: tn.ID, Name: "Chennai", Code: "CHN", IsActive: true},
		{StateID: tn.ID, Name: "Madurai", Code: "MDU", IsActive: true},
		{StateID: tn.ID, Name: "Coimbatore", Code: "COI", IsActive: true},

		// Karnataka - Mysore silk
		{StateID: ka.ID, Name: "Mysore", Code: "MYS", IsActive: true},
		{StateID: ka.ID, Name: "Bangalore", Code: "BLR", IsActive: true},
		{StateID: ka.ID, Name: "Belgaum", Code: "BEL", IsActive: true},

		// West Bengal - Tant sarees
		{StateID: wb.ID, Name: "Kolkata", Code: "KOL", IsActive: true},
		{StateID: wb.ID, Name: "Murshidabad", Code: "MUR", IsActive: true},
		{StateID: wb.ID, Name: "Hooghly", Code: "HOO", IsActive: true},

		// Bihar - Madhubani art
		{StateID: br.ID, Name: "Madhubani", Code: "MAD", IsActive: true},
		{StateID: br.ID, Name: "Patna", Code: "PAT", IsActive: true},

		// Maharashtra - Paithani
		{StateID: mh.ID, Name: "Mumbai", Code: "MUM", IsActive: true},
		{StateID: mh.ID, Name: "Pune", Code: "PUN", IsActive: true},
		{StateID: mh.ID, Name: "Aurangabad", Code: "AUR", IsActive: true},
	}

	for _, district := range districts {
		var existing models.District
		if err := config.DB.Where("name = ? AND state_id = ?", district.Name, district.StateID).First(&existing).Error; err == nil {
			log.Printf("District '%s' already exists, skipping...", district.Name)
			continue
		}

		if err := config.DB.Create(&district).Error; err != nil {
			log.Printf("Failed to create district '%s': %v", district.Name, err)
			continue
		}

		log.Printf("✅ Created district: %s", district.Name)
	}
}

func seedRegions() {
	log.Println("🎨 Seeding craft regions...")

	// Get state IDs
	var up, kl, tn, ka, wb, br, mh, gj, rj models.State
	config.DB.Where("code = ?", "UP").First(&up)
	config.DB.Where("code = ?", "KL").First(&kl)
	config.DB.Where("code = ?", "TN").First(&tn)
	config.DB.Where("code = ?", "KA").First(&ka)
	config.DB.Where("code = ?", "WB").First(&wb)
	config.DB.Where("code = ?", "BR").First(&br)
	config.DB.Where("code = ?", "MH").First(&mh)
	config.DB.Where("code = ?", "GJ").First(&gj)
	config.DB.Where("code = ?", "RJ").First(&rj)

	regions := []models.Region{
		{
			Name:         "Lucknow",
			Slug:         "lucknow",
			Type:         "City",
			StateID:      &up.ID,
			Description:  "Famous for its exquisite Chikankari embroidery, a traditional art form passed down through generations. Lucknow's Chikankari work is known for its delicate and intricate patterns on fine cotton and georgette fabrics.",
			FamousFor:    "Chikankari Sarees, Chikankari Kurtis, Chikankari Dresses, Hand Embroidered Textiles",
			DisplayOrder: 1,
			IsActive:     true,
		},
		{
			Name:         "Kerala",
			Slug:         "kerala",
			Type:         "State",
			StateID:      &kl.ID,
			Description:  "Renowned for its traditional Kasavu sarees with golden zari borders. These handloom sarees are an integral part of Kerala's culture and are worn during festivals, especially Onam.",
			FamousFor:    "Kasavu Sarees, Kerala Cotton Sarees, Set Mundu, Tissue Sarees",
			DisplayOrder: 2,
			IsActive:     true,
		},
		{
			Name:         "Kanchipuram",
			Slug:         "kanchipuram",
			Type:         "City",
			StateID:      &tn.ID,
			Description:  "The silk capital of India, Kanchipuram is world-famous for its pure silk sarees with rich zari work. Each Kanchipuram saree is a masterpiece, known for its durability, vibrant colors, and traditional temple designs.",
			FamousFor:    "Kanchipuram Silk Sarees, Pure Silk Sarees, Temple Border Sarees, Zari Work",
			DisplayOrder: 3,
			IsActive:     true,
		},
		{
			Name:         "Mysore",
			Slug:         "mysore",
			Type:         "City",
			StateID:      &ka.ID,
			Description:  "Known for its soft and lightweight Mysore silk sarees with beautiful zari borders. These sarees are perfect for daily wear and special occasions, offering comfort and elegance.",
			FamousFor:    "Mysore Silk Sarees, Crepe Silk, Traditional Silk",
			DisplayOrder: 4,
			IsActive:     true,
		},
		{
			Name:         "Bengal",
			Slug:         "bengal",
			Type:         "Region",
			StateID:      &wb.ID,
			Description:  "Home to the traditional Tant sarees, made from fine cotton yarn. These handloom sarees are lightweight, comfortable, and perfect for the humid climate. Bengal is also famous for Jamdani and Baluchari sarees.",
			FamousFor:    "Tant Sarees, Jamdani Sarees, Baluchari Sarees, Handloom Cotton",
			DisplayOrder: 5,
			IsActive:     true,
		},
		{
			Name:         "Madhubani",
			Slug:         "madhubani",
			Type:         "City",
			StateID:      &br.ID,
			Description:  "Famous for Madhubani art, a traditional painting style now featured on sarees and textiles. These hand-painted sarees showcase vibrant colors and mythological themes.",
			FamousFor:    "Madhubani Art Sarees, Hand-painted Textiles, Traditional Bihar Crafts",
			DisplayOrder: 6,
			IsActive:     true,
		},
		{
			Name:         "Banarasi",
			Slug:         "banarasi",
			Type:         "Region",
			StateID:      &up.ID,
			Description:  "Varanasi (Banaras) is renowned for its luxurious Banarasi silk sarees with intricate brocade work. These sarees are a symbol of Indian tradition and are often chosen for weddings.",
			FamousFor:    "Banarasi Silk Sarees, Brocade Work, Wedding Sarees, Zari Work",
			DisplayOrder: 7,
			IsActive:     true,
		},
		{
			Name:         "Paithani",
			Slug:         "paithani",
			Type:         "Region",
			StateID:      &mh.ID,
			Description:  "Maharashtra's pride, Paithani sarees are known for their rich peacock and lotus motifs woven in gold and silver threads. These sarees are treasured heirlooms.",
			FamousFor:    "Paithani Sarees, Silk Sarees, Traditional Maharashtra Weaves",
			DisplayOrder: 8,
			IsActive:     true,
		},
		{
			Name:         "Bandhani",
			Slug:         "bandhani",
			Type:         "Region",
			StateID:      &gj.ID,
			Description:  "Gujarat and Rajasthan are famous for Bandhani (tie-dye) sarees. This ancient technique creates beautiful patterns with vibrant colors, perfect for festive occasions.",
			FamousFor:    "Bandhani Sarees, Tie-Dye Textiles, Traditional Gujarat Crafts",
			DisplayOrder: 9,
			IsActive:     true,
		},
		{
			Name:         "Chanderi",
			Slug:         "chanderi",
			Type:         "Region",
			StateID:      &up.ID,
			Description:  "Chanderi sarees are known for their sheer texture, lightweight, and glossy transparency. These sarees are perfect for summer and feature traditional motifs.",
			FamousFor:    "Chanderi Sarees, Silk Cotton Sarees, Lightweight Sarees",
			DisplayOrder: 10,
			IsActive:     true,
		},
	}

	for _, region := range regions {
		var existing models.Region
		if err := config.DB.Where("slug = ?", region.Slug).First(&existing).Error; err == nil {
			log.Printf("Region '%s' already exists, skipping...", region.Name)
			continue
		}

		if err := config.DB.Create(&region).Error; err != nil {
			log.Printf("Failed to create region '%s': %v", region.Name, err)
			continue
		}

		log.Printf("✅ Created region: %s", region.Name)
	}

	log.Printf("✅ Seeded %d regions", len(regions))
}
